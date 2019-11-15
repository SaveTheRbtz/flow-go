package badger_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dapperlabs/flow-go/model/flow"

	"github.com/dapperlabs/flow-go/sdk/emulator/storage"
	"github.com/dapperlabs/flow-go/sdk/emulator/types"

	"github.com/dapperlabs/flow-go/sdk/emulator/storage/badger"
	"github.com/dapperlabs/flow-go/utils/unittest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlocks(t *testing.T) {
	store, dir := setupStore(t)
	defer func() {
		require.Nil(t, store.Close())
		require.Nil(t, os.RemoveAll(dir))
	}()

	block1 := types.Block{
		Number: 1,
	}
	block2 := types.Block{
		Number: 2,
	}

	t.Run("should return error for not found", func(t *testing.T) {
		t.Run("GetBlockByHash", func(t *testing.T) {
			_, err := store.GetBlockByHash(unittest.HashFixture(32))
			if assert.Error(t, err) {
				assert.IsType(t, storage.ErrNotFound{}, err)
			}
		})

		t.Run("GetBlockByNumber", func(t *testing.T) {
			_, err := store.GetBlockByNumber(block1.Number)
			if assert.Error(t, err) {
				assert.IsType(t, storage.ErrNotFound{}, err)
			}
		})

		t.Run("GetLatestBlock", func(t *testing.T) {
			_, err := store.GetLatestBlock()
			if assert.Error(t, err) {
				assert.IsType(t, storage.ErrNotFound{}, err)
			}
		})
	})

	t.Run("should be able to insert block", func(t *testing.T) {
		err := store.InsertBlock(block1)
		assert.NoError(t, err)
	})

	// insert block 1
	err := store.InsertBlock(block1)
	assert.NoError(t, err)

	t.Run("should be able to get inserted block", func(t *testing.T) {
		t.Run("GetBlockByNumber", func(t *testing.T) {
			block, err := store.GetBlockByNumber(block1.Number)
			assert.NoError(t, err)
			assert.Equal(t, block1, block)
		})

		t.Run("GetBlockByHash", func(t *testing.T) {
			block, err := store.GetBlockByHash(block1.Hash())
			assert.NoError(t, err)
			assert.Equal(t, block1, block)
		})

		t.Run("GetLatestBlock", func(t *testing.T) {
			block, err := store.GetLatestBlock()
			assert.NoError(t, err)
			assert.Equal(t, block1, block)
		})
	})

	// insert block 2
	err = store.InsertBlock(block2)
	assert.NoError(t, err)

	t.Run("Latest block should update", func(t *testing.T) {
		block, err := store.GetLatestBlock()
		assert.NoError(t, err)
		assert.Equal(t, block2, block)
	})
}

func TestTransactions(t *testing.T) {
	store, dir := setupStore(t)
	defer func() {
		require.Nil(t, store.Close())
		require.Nil(t, os.RemoveAll(dir))
	}()

	tx := unittest.TransactionFixture()

	t.Run("should return error for not found", func(t *testing.T) {
		_, err := store.GetTransaction(tx.Hash())
		if assert.Error(t, err) {
			assert.IsType(t, storage.ErrNotFound{}, err)
		}
	})

	t.Run("should be able to insert tx", func(t *testing.T) {
		err := store.InsertTransaction(tx)
		assert.NoError(t, err)

		t.Run("should be able to get inserted tx", func(t *testing.T) {
			storedTx, err := store.GetTransaction(tx.Hash())
			require.Nil(t, err)
			assert.Equal(t, tx, storedTx)
		})
	})
}

func TestLedger(t *testing.T) {
	t.Run("get/set", func(t *testing.T) {
		store, dir := setupStore(t)
		defer func() {
			require.Nil(t, store.Close())
			require.Nil(t, os.RemoveAll(dir))
		}()

		var blockNumber uint64 = 1
		ledger := unittest.LedgerFixture()

		t.Run("should get able to set ledger", func(t *testing.T) {
			err := store.SetLedger(blockNumber, ledger)
			assert.NoError(t, err)
		})

		t.Run("should be to get set ledger", func(t *testing.T) {
			view, err := store.GetLedgerView(blockNumber)
			assert.NoError(t, err)
			assert.Equal(t, ledger.NewView(), &view)
		})
	})

	t.Run("versioning", func(t *testing.T) {
		store, dir := setupStore(t)
		defer func() {
			require.Nil(t, store.Close())
			require.Nil(t, os.RemoveAll(dir))
		}()

		// Create a list of ledgers, where the ledger at index i has
		// keys (i+2)-1->(i+2)+1 set.
		totalBlocks := 10
		var ledgers []flow.Ledger
		for i := 2; i < totalBlocks+2; i++ {
			ledger := make(flow.Ledger)
			for j := i - 1; j <= i+1; j++ {
				ledger[fmt.Sprintf("%d", j)] = []byte{byte(j)}
			}
			ledgers = append(ledgers, ledger)
		}
		require.Equal(t, totalBlocks, len(ledgers))

		// Insert all the ledgers, starting with block 1.
		// This will result in a ledger state that looks like this:
		// Block 1: {1: 1, 2: 2, 3: 3}
		// Block 2: {2: 2, 3: 3, 4: 4}
		// ...
		for i, ledger := range ledgers {
			err := store.SetLedger(uint64(i+1), ledger)
			require.NoError(t, err)
		}

		// We didn't insert anything at block 0, so this should be empty.
		t.Run("should return empty view for block 0", func(t *testing.T) {
			view, err := store.GetLedgerView(0)
			require.NoError(t, err)
			expected := make(flow.Ledger).NewView()
			assert.Equal(t, *expected, view)
		})

		// View at block 1 should have keys 1, 2, 3
		t.Run("should version the first written block", func(t *testing.T) {
			view, err := store.GetLedgerView(1)
			require.NoError(t, err)
			for i := 1; i <= 3; i++ {
				val, ok := view.Get(fmt.Sprintf("%d", i))
				assert.True(t, ok)
				assert.Equal(t, []byte{byte(i)}, val)
			}
		})

		// View at block N should have values 1->N+2
		t.Run("should version the first written block", func(t *testing.T) {
			for block := 2; block < totalBlocks; block++ {
				view, err := store.GetLedgerView(uint64(block))
				require.NoError(t, err)
				for i := 1; i <= block+2; i++ {
					val, ok := view.Get(fmt.Sprintf("%d", i))
					assert.True(t, ok)
					assert.Equal(t, []byte{byte(i)}, val)
				}
			}
		})
	})
}

func TestEvents(t *testing.T) {
	store, dir := setupStore(t)
	defer func() {
		require.Nil(t, store.Close())
		require.Nil(t, os.RemoveAll(dir))
	}()

	t.Run("should be able to insert events", func(t *testing.T) {
		events := []flow.Event{unittest.EventFixture()}
		var blockNumber uint64 = 1

		err := store.InsertEvents(blockNumber, events...)
		assert.NoError(t, err)

		t.Run("should be able to get inserted events", func(t *testing.T) {
			gotEvents, err := store.GetEvents("", blockNumber, blockNumber)
			assert.NoError(t, err)
			assert.Equal(t, events, gotEvents)
		})
	})

	t.Run("should be able to insert many events", func(t *testing.T) {
		// block 1 will have 1 event type=1
		// block 2 will have 2 events, types=1,2
		// and so on...
		eventsByBlock := make(map[uint64][]flow.Event)
		for i := 1; i <= 10; i++ {
			var events []flow.Event
			for j := 1; j <= i; j++ {
				event := unittest.EventFixture()
				event.Type = fmt.Sprintf("%d", j)
				events = append(events, event)
			}
			eventsByBlock[uint64(i)] = events
			err := store.InsertEvents(uint64(i), events...)
			assert.NoError(t, err)
		}

		t.Run("should be able to query by block", func(t *testing.T) {
			t.Run("block 1", func(t *testing.T) {
				gotEvents, err := store.GetEvents("", 1, 1)
				assert.NoError(t, err)
				assert.Equal(t, eventsByBlock[1], gotEvents)
			})

			t.Run("block 2", func(t *testing.T) {
				gotEvents, err := store.GetEvents("", 2, 2)
				assert.NoError(t, err)
				assert.Equal(t, eventsByBlock[2], gotEvents)
			})
		})

		t.Run("should be able to query by block interval", func(t *testing.T) {
			t.Run("block 1->2", func(t *testing.T) {
				gotEvents, err := store.GetEvents("", 1, 2)
				assert.NoError(t, err)
				assert.Equal(t, append(eventsByBlock[1], eventsByBlock[2]...), gotEvents)
			})

			t.Run("block 5->10", func(t *testing.T) {
				gotEvents, err := store.GetEvents("", 5, 10)
				assert.NoError(t, err)

				var expectedEvents []flow.Event
				for i := 5; i <= 10; i++ {
					expectedEvents = append(expectedEvents, eventsByBlock[uint64(i)]...)
				}
				assert.Equal(t, expectedEvents, gotEvents)
			})
		})

		t.Run("should be able to query by event type", func(t *testing.T) {
			t.Run("type=1, block=1", func(t *testing.T) {
				// should be one event type=1 in block 1
				gotEvents, err := store.GetEvents("1", 1, 1)
				assert.NoError(t, err)
				assert.Len(t, gotEvents, 1)
				assert.Equal(t, "1", gotEvents[0].Type)
			})

			t.Run("type=1, block=1->10", func(t *testing.T) {
				// should be 10 events type=1 in Blocks 1->10
				gotEvents, err := store.GetEvents("1", 1, 10)
				assert.NoError(t, err)
				assert.Len(t, gotEvents, 10)
				for _, event := range gotEvents {
					assert.Equal(t, "1", event.Type)
				}
			})

			t.Run("type=2, block=1", func(t *testing.T) {
				// should be 0 type=2 events here
				gotEvents, err := store.GetEvents("2", 1, 1)
				assert.NoError(t, err)
				assert.Len(t, gotEvents, 0)
			})
		})
	})
}

func TestPersistence(t *testing.T) {
	store, dir := setupStore(t)
	defer func() {
		require.Nil(t, store.Close())
		require.Nil(t, os.RemoveAll(dir))
	}()

	block := types.Block{Number: 1}
	tx := unittest.TransactionFixture()
	events := []flow.Event{unittest.EventFixture()}
	ledger := unittest.LedgerFixture()

	// insert some stuff to to the store
	err := store.InsertBlock(block)
	assert.NoError(t, err)
	err = store.InsertTransaction(tx)
	assert.NoError(t, err)
	err = store.InsertEvents(block.Number, events...)
	assert.NoError(t, err)
	err = store.SetLedger(block.Number, ledger)

	// close the store
	err = store.Close()
	assert.NoError(t, err)

	// create a new store with the same database directory
	store, err = badger.New(dir)
	require.Nil(t, err)

	// should be able to retrieve what we stored
	gotBlock, err := store.GetLatestBlock()
	assert.NoError(t, err)
	assert.Equal(t, block, gotBlock)

	gotTx, err := store.GetTransaction(tx.Hash())
	assert.NoError(t, err)
	assert.Equal(t, tx, gotTx)

	gotEvents, err := store.GetEvents("", block.Number, block.Number)
	assert.NoError(t, err)
	assert.Equal(t, events, gotEvents)

	gotLedger, err := store.GetLedgerView(block.Number)
	assert.NoError(t, err)
	assert.Equal(t, ledger.NewView(), &gotLedger)
}

// setupStore creates a temporary directory for the Badger and creates a
// badger.Store instance. The caller is responsible for closing the store
// and deleting the temporary directory.
func setupStore(t *testing.T) (badger.Store, string) {
	dir, err := ioutil.TempDir("", "badger-test")
	require.Nil(t, err)

	store, err := badger.New(dir)
	require.Nil(t, err)

	return store, dir
}
