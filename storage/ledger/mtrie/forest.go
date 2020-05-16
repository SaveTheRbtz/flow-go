package mtrie

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"sort"

	lru "github.com/hashicorp/golang-lru"

	"github.com/dapperlabs/flow-go/storage/ledger/mtrie/proof"

	"github.com/dapperlabs/flow-go/storage/ledger/mtrie/node"

	"github.com/dapperlabs/flow-go/storage/ledger/mtrie/trie"

	"github.com/dapperlabs/flow-go/storage/ledger/mtrie/common"
)

// MForest is an in memory forest (collection of tries)
type MForest struct {
	tries         *lru.Cache
	dir           string
	cacheSize     int
	maxHeight     int // height of the tree
	keyByteSize   int // acceptable number of bytes for key
	onTreeEvicted func(tree *trie.MTrie) error
}

// NewMForest returns a new instance of memory forest
func NewMForest(maxHeight int, trieStorageDir string, trieCacheSize int, onTreeEvicted func(tree *trie.MTrie) error) (*MForest, error) {

	var cache *lru.Cache
	var err error

	if onTreeEvicted != nil {
		cache, err = lru.NewWithEvict(trieCacheSize, func(key interface{}, value interface{}) {
			trie, ok := value.(*trie.MTrie)
			if !ok {
				panic(fmt.Sprintf("cache contains item of type %T", value))
			}
			//TODO Log error
			_ = onTreeEvicted(trie)
		})
	} else {
		cache, err = lru.New(trieCacheSize)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot create forest cache: %w", err)
	}

	forest := &MForest{tries: cache,
		maxHeight:     maxHeight,
		dir:           trieStorageDir,
		cacheSize:     trieCacheSize,
		keyByteSize:   (maxHeight - 1) / 8,
		onTreeEvicted: onTreeEvicted,
	}

	// add empty roothash
	emptyTrie := trie.NewMTrie(maxHeight)
	emptyTrie.Number = uint64(0)
	emptyTrie.SetRootHash(common.GetDefaultHashForHeight(maxHeight - 1))

	err = forest.AddTrie(emptyTrie)
	if err != nil {
		return nil, err
	}
	return forest, nil
}

// GetTrie returns trie at specific rootHash
// warning, use this function for read-only operation
func (f *MForest) GetTrie(rootHash []byte) (*trie.MTrie, error) {
	encRootHash := hex.EncodeToString(rootHash)
	// if in the cache

	if ent, ok := f.tries.Get(encRootHash); ok {
		return ent.(*trie.MTrie), nil
	}

	// otherwise try to load from disk
	trie, err := f.LoadTrie(filepath.Join(f.dir, encRootHash))
	if err != nil {
		return nil, fmt.Errorf("trie with the given rootHash [%v] not found: %w", hex.EncodeToString(rootHash), err)
	}

	return trie, nil
}

// AddTrie adds a trie to the forest
func (f *MForest) AddTrie(trie *trie.MTrie) error {
	// TODO check if not exist
	encoded := trie.StringRootHash()
	f.tries.Add(encoded, trie)
	return nil
}

// RemoveTrie removes a trie to the forest
func (f *MForest) RemoveTrie(rootHash []byte) {
	// TODO remove from the file as well
	encRootHash := hex.EncodeToString(rootHash)
	f.tries.Remove(encRootHash)
}

// GetEmptyRootHash returns the rootHash of empty forest
func (f *MForest) GetEmptyRootHash() []byte {
	newRoot := node.NewNode(f.maxHeight - 1)
	rootHash := newRoot.ComputeNodeHash(false)
	return rootHash
}

// Read reads values for an slice of keys and returns values and error (if any)
func (f *MForest) Read(keys [][]byte, rootHash []byte) ([][]byte, error) {

	// no key no change
	if len(keys) == 0 {
		return [][]byte{}, nil
	}

	// lookup the trie by rootHash
	trie, err := f.GetTrie(rootHash)
	if err != nil {
		return nil, err
	}

	// sort keys and deduplicate keys
	sortedKeys := make([][]byte, 0)
	keyOrgIndex := make(map[string][]int)
	for i, key := range keys {
		// check key sizes
		if len(key) != f.keyByteSize {
			return nil, fmt.Errorf("key size doesn't match the trie height: %x", key)
		}
		// only collect dupplicated keys once
		if _, ok := keyOrgIndex[hex.EncodeToString(key)]; !ok {
			sortedKeys = append(sortedKeys, key)
			keyOrgIndex[hex.EncodeToString(key)] = []int{i}
		} else {
			// handles duplicated keys
			keyOrgIndex[hex.EncodeToString(key)] = append(keyOrgIndex[hex.EncodeToString(key)], i)
		}
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		return bytes.Compare(sortedKeys[i], sortedKeys[j]) < 0
	})

	values, err := trie.UnsafeRead(sortedKeys)

	if err != nil {
		return nil, err
	}

	// reconstruct the values in the same key order that called the method
	orderedValues := make([][]byte, len(keys))
	for i, k := range sortedKeys {
		for _, j := range keyOrgIndex[hex.EncodeToString(k)] {
			orderedValues[j] = values[i]
		}
	}
	return orderedValues, nil
}

// Update updates the Values for the registers and returns rootHash and error (if any)
func (f *MForest) Update(keys [][]byte, values [][]byte, rootHash []byte) ([]byte, error) {

	// no key no change
	if len(keys) == 0 {
		return rootHash, nil
	}

	// sort keys and deduplicate keys (we only consider the last occurrence, and ignore the rest)
	sortedKeys := make([][]byte, 0)
	valueMap := make(map[string][]byte)
	for i, key := range keys {
		// check key sizes
		if len(key) != f.keyByteSize {
			return nil, fmt.Errorf("key size doesn't match the trie height: %x", key)
		}
		// check if doesn't exist
		if _, ok := valueMap[hex.EncodeToString(key)]; !ok {
			//do something here
			sortedKeys = append(sortedKeys, key)
			valueMap[hex.EncodeToString(key)] = values[i]
		} else {
			valueMap[hex.EncodeToString(key)] = values[i]
		}
	}

	// TODO we might be able to remove this
	sort.Slice(sortedKeys, func(i, j int) bool {
		return bytes.Compare(sortedKeys[i], sortedKeys[j]) < 0
	})

	sortedValues := make([][]byte, 0, len(sortedKeys))
	for _, key := range sortedKeys {
		sortedValues = append(sortedValues, valueMap[hex.EncodeToString(key)])
	}

	parentTrie, err := f.GetTrie(rootHash)
	if err != nil {
		return nil, err
	}

	newTrie := trie.NewMTrie(f.maxHeight)
	newTrie.ParentRootHash = parentTrie.Root.GetNodeHash()
	newTrie.Number = parentTrie.Number + 1

	err = newTrie.UnsafeUpdate(parentTrie, sortedKeys, sortedValues)
	if err != nil {
		return nil, err
	}

	newTrie.Root.PopulateNodeHashValues()
	newRootHash := newTrie.Root.GetNodeHash()
	newTrie.SetRootHash(newRootHash)
	err = f.AddTrie(newTrie)
	if err != nil {
		return nil, err
	}
	// go func() {
	// 	_ = newTrie.Store(filepath.Join(f.dir, hex.EncodeToString(newRootHash)))
	// }()
	return newRootHash, nil
}

// Proofs returns a batch proof for the given keys
func (f *MForest) Proofs(keys [][]byte, rootHash []byte) (*proof.BatchProof, error) {

	// no key, empty batchproof
	if len(keys) == 0 {
		return proof.NewBatchProof(), nil
	}

	// look up for non exisitng keys
	notFoundKeys := make([][]byte, 0)
	notFoundValues := make([][]byte, 0)
	retValues, err := f.Read(keys, rootHash)
	if err != nil {
		return nil, err
	}

	sortedKeys := make([][]byte, 0)
	keyOrgIndex := make(map[string][]int)
	for i, key := range keys {
		// check key sizes
		if len(key) != f.keyByteSize {
			return nil, fmt.Errorf("key size doesn't match the trie height: %x", key)
		}
		// only collect duplicated keys once
		if _, ok := keyOrgIndex[hex.EncodeToString(key)]; !ok {
			sortedKeys = append(sortedKeys, key)
			keyOrgIndex[hex.EncodeToString(key)] = []int{i}

			// add it only once
			if len(retValues[i]) == 0 {
				notFoundKeys = append(notFoundKeys, key)
				notFoundValues = append(notFoundValues, []byte{})
			}
		} else {
			// handles duplicated keys
			keyOrgIndex[hex.EncodeToString(key)] = append(keyOrgIndex[hex.EncodeToString(key)], i)
		}

	}

	stateTrie, err := f.GetTrie(rootHash)
	if err != nil {
		return nil, err
	}

	// if we have to insert empty values
	if len(notFoundKeys) > 0 {
		newTrie := trie.NewMTrie(f.maxHeight)
		newRoot := newTrie.Root

		sort.Slice(notFoundKeys, func(i, j int) bool {
			return bytes.Compare(notFoundKeys[i], notFoundKeys[j]) < 0
		})

		err = newTrie.UnsafeUpdate(stateTrie, notFoundKeys, notFoundValues)
		if err != nil {
			return nil, err
		}

		// rootHash shouldn't change
		if !bytes.Equal(newRoot.GetNodeHash(), rootHash) {
			return nil, errors.New("Root hash has changed during the operation")
		}
		stateTrie = newTrie
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		return bytes.Compare(sortedKeys[i], sortedKeys[j]) < 0
	})

	bp := proof.NewBatchProofWithEmptyProofs(len(sortedKeys))

	for _, p := range bp.Proofs {
		p.Flags = make([]byte, f.keyByteSize)
		p.Inclusion = false
	}

	err = stateTrie.UnsafeProofs(sortedKeys, bp.Proofs)
	if err != nil {
		return nil, err
	}

	// reconstruct the proofs in the same key order that called the method
	retbp := proof.NewBatchProofWithEmptyProofs(len(keys))
	for i, k := range sortedKeys {
		for _, j := range keyOrgIndex[hex.EncodeToString(k)] {
			retbp.Proofs[j] = bp.Proofs[i]
		}
	}

	return retbp, nil
}

// StoreTrie stores a trie on disk
func (f *MForest) StoreTrie(rootHash []byte, path string) error {
	trie, err := f.GetTrie(rootHash)
	if err != nil {
		return err
	}
	return trie.Store(path)
}

// LoadTrie loads a trie from the disk
func (f *MForest) LoadTrie(path string) (*trie.MTrie, error) {
	stateTrie := trie.NewMTrie(f.maxHeight)
	err := stateTrie.Load(path)
	if err != nil {
		return nil, err
	}
	stateTrie.Root.PopulateNodeHashValues()
	if !bytes.Equal(stateTrie.RootHash(), stateTrie.Root.GetNodeHash()) {
		return nil, errors.New("error loading a trie, rootHash doesn't match")
	}
	err = f.AddTrie(stateTrie)
	if err != nil {
		return nil, err
	}

	return stateTrie, nil
}

func (f *MForest) Size() int {
	return f.tries.Len()
}
