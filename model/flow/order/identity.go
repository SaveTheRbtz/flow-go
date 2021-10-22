// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package order

import (
	"bytes"
	"encoding/binary"

	"github.com/onflow/flow-go/model/flow"
)

// Canonical represents the canonical ordering for identity lists.
var Canonical = ByNodeIDAsc

func ByNodeIDAsc(identity1 *flow.Identity, identity2 *flow.Identity) bool {
	return bytes.Compare(identity1.NodeID[:], identity2.NodeID[:]) < 0
}

func ByIdentifierAsc(idA flow.Identifier, idB flow.Identifier) bool {
	num1 := idA[:]
	num2 := idB[:]
	first := binary.BigEndian.Uint64(num1)
	second := binary.BigEndian.Uint64(num2)

	if first == second {
		num1 = num1[8:]
		num2 = num2[8:]
		first = binary.BigEndian.Uint64(num1)
		second = binary.BigEndian.Uint64(num2)

		if first == second {
			num1 = num1[8:]
			num2 = num2[8:]
			first = binary.BigEndian.Uint64(num1)
			second = binary.BigEndian.Uint64(num2)

			if first == second {
				num1 = num1[8:]
				num2 = num2[8:]
				first = binary.BigEndian.Uint64(num1)
				second = binary.BigEndian.Uint64(num2)
				return first < second
			}
			return first < second
		}
		return first < second
	}
	return first < second
}

func ByReferenceOrder(nodeIDs []flow.Identifier) func(*flow.Identity, *flow.Identity) bool {
	indices := make(map[flow.Identifier]uint)
	for index, nodeID := range nodeIDs {
		_, ok := indices[nodeID]
		if ok {
			panic("should never order by reference order with duplicate node IDs")
		}
		indices[nodeID] = uint(index)
	}
	return func(identity1 *flow.Identity, identity2 *flow.Identity) bool {
		return indices[identity1.NodeID] < indices[identity2.NodeID]
	}
}
