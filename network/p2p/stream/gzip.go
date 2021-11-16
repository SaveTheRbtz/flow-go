package p2pstream

import (
	libp2pnet "github.com/libp2p/go-libp2p-core/network"

	"github.com/onflow/flow-go/network/compressor"
	"github.com/onflow/flow-go/network/p2p/compressed"
)

// GzipStream is a stream compression creates and returns a gzip-compressed stream out of input stream.
type GzipStream struct {
	handler libp2pnet.StreamHandler
}

func (g GzipStream) NewStream(s libp2pnet.Stream) (libp2pnet.Stream, error) {
	return compressed.NewCompressedStream(s, compressor.GzipStreamCompressor{})
}

func (g GzipStream) Handler() libp2pnet.StreamHandler {
	return g.handler
}
