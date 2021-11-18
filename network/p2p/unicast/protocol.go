package unicast

import (
	"fmt"
	"strings"

	libp2pnet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/model/flow"
)

// Flow Libp2p protocols
const (
	// FlowLibP2PProtocolCommonPrefix is the common prefix for all Libp2p protocol IDs used for Flow
	// ALL Flow libp2p protocols must start with this prefix.
	FlowLibP2PProtocolCommonPrefix = "/flow"

	// FlowLibP2POneToOneProtocolIDPrefix is a unique Libp2p protocol ID prefix for Flow (https://docs.libp2p.io/concepts/protocols/)
	// All nodes communicate with each other using this protocol id suffixed with the id of the root block
	FlowLibP2POneToOneProtocolIDPrefix = FlowLibP2PProtocolCommonPrefix + "/push/"

	// FlowLibP2PPingProtocolPrefix is the Flow Ping protocol prefix
	FlowLibP2PPingProtocolPrefix = FlowLibP2PProtocolCommonPrefix + "/ping/"

	// FlowLibP2PProtocolGzipCompressedOneToOne represents the protocol id for compressed streams under gzip compressor.
	FlowLibP2PProtocolGzipCompressedOneToOne = FlowLibP2POneToOneProtocolIDPrefix + "/gzip/"
)

// IsFlowProtocolStream returns true if the libp2p stream is for a Flow protocol
func IsFlowProtocolStream(s libp2pnet.Stream) bool {
	p := string(s.Protocol())
	return strings.HasPrefix(p, FlowLibP2PProtocolCommonPrefix)
}

func FlowProtocolID(rootBlockID flow.Identifier) protocol.ID {
	return protocol.ID(FlowLibP2POneToOneProtocolIDPrefix + rootBlockID.String())
}

func PingProtocolId(rootBlockID flow.Identifier) protocol.ID {
	return protocol.ID(FlowLibP2PPingProtocolPrefix + rootBlockID.String())
}

func FlowGzipProtocolId(rootBlockID flow.Identifier) protocol.ID {
	return protocol.ID(FlowLibP2PProtocolGzipCompressedOneToOne + rootBlockID.String())
}

type ProtocolName string
type ProtocolFactory func(zerolog.Logger, flow.Identifier, libp2pnet.StreamHandler) Protocol

func ToProtocolFactory(name ProtocolName) (ProtocolFactory, error) {
	switch name {
	case GzipCompressionUnicast:
		return func(logger zerolog.Logger, rootBlockID flow.Identifier, handler libp2pnet.StreamHandler) Protocol {
			return NewGzipCompressedUnicast(logger, rootBlockID, handler)
		}, nil
	default:
		return nil, fmt.Errorf("unknown unicast protocol name: %s", name)
	}
}

type Protocol interface {
	NewStream(s libp2pnet.Stream) (libp2pnet.Stream, error)
	Handler() libp2pnet.StreamHandler
	ProtocolId() protocol.ID
}
