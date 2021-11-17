package unicast

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/libp2p/go-libp2p-core/host"
	libp2pnet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/network/p2p"
)

// maximum number of milliseconds to wait between attempts for a 1-1 direct connection
const maxConnectAttemptSleepDuration = 5

type ProtocolBuilder struct {
	logger         zerolog.Logger
	host           host.Host
	unicasts       []Protocol
	defaultHandler libp2pnet.StreamHandler
	rootBlockId    flow.Identifier
}

func NewProtocolBuilder(logger zerolog.Logger, host host.Host, rootBlockId flow.Identifier) *ProtocolBuilder {

	builder := &ProtocolBuilder{
		logger:      logger,
		host:        host,
		rootBlockId: rootBlockId,
	}

	return builder
}

func (builder *ProtocolBuilder) WithDefaultHandler(defaultHandler libp2pnet.StreamHandler) {
	defaultProtocolID := p2p.FlowProtocolID(builder.rootBlockId)
	builder.defaultHandler = defaultHandler

	builder.unicasts = []Protocol{
		&PlainStream{
			protocolId: defaultProtocolID,
			handler:    defaultHandler,
		},
	}

	builder.host.SetStreamHandler(defaultProtocolID, defaultHandler)
}

func (builder *ProtocolBuilder) Register(unicast ProtocolName) error {
	factory, err := ToProtocolFactory(unicast)
	if err != nil {
		return fmt.Errorf("could not translate protocol name into factory: %w", err)
	}

	u := factory(builder.logger, builder.rootBlockId, builder.defaultHandler)

	builder.unicasts = append(builder.unicasts, u)
	builder.host.SetStreamHandler(u.ProtocolId(), u.Handler())

	return nil
}

// CreateStream makes at most `maxAttempts` to create a stream with the peer.
// This was put in as a fix for #2416. PubSub and 1-1 communication compete with each other when trying to connect to
// remote nodes and once in a while NewStream returns an error 'both yamux endpoints are clients'.
//
// Note that in case an existing TCP connection underneath to `peerID` exists, that connection is utilized for creating a new stream.
// The multiaddr.Multiaddr return value represents the addresses of `peerID` we dial while trying to create a stream to it.
func (builder *ProtocolBuilder) CreateStream(ctx context.Context, peerID peer.ID, maxAttempts int) (libp2pnet.Stream, []multiaddr.Multiaddr, error) {
	var errs error

	for i := len(builder.unicasts) - 1; i >= 0; i-- {
		s, addrs, err := builder.createStreamWithProtocol(ctx, builder.unicasts[i].ProtocolId(), peerID, maxAttempts)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		// return first successful stream
		return s, addrs, nil
	}

	return nil, nil, fmt.Errorf("could not create stream on any available unicast protocol: %w", errs)
}

func (builder *ProtocolBuilder) createStreamWithProtocol(ctx context.Context,
	protocolID protocol.ID,
	peerID peer.ID,
	maxAttempts int) (libp2pnet.Stream, []multiaddr.Multiaddr, error) {

	var errs error
	var s libp2pnet.Stream
	var retries = 0
	var dialAddr []multiaddr.Multiaddr // address on which we dial peerID
	for ; retries < maxAttempts; retries++ {
		select {
		case <-ctx.Done():
			return nil, nil, fmt.Errorf("context done before stream could be created (retry attempt: %d, errors: %w)", retries, errs)
		default:
		}

		// libp2p internally uses swarm dial - https://github.com/libp2p/go-libp2p-swarm/blob/master/swarm_dial.go
		// to connect to a peer. Swarm dial adds a back off each time it fails connecting to a peer. While this is
		// the desired behaviour for pub-sub (1-k style of communication) for 1-1 style we want to retry the connection
		// immediately without backing off and fail-fast.
		// Hence, explicitly cancel the dial back off (if any) and try connecting again

		// cancel the dial back off (if any), since we want to connect immediately
		network := builder.host.Network()
		dialAddr = network.Peerstore().Addrs(peerID)
		if swm, ok := network.(*swarm.Swarm); ok {
			swm.Backoff().Clear(peerID)
		}

		// if this is a retry attempt, wait for some time before retrying
		if retries > 0 {
			// choose a random interval between 0 to 5
			// (to ensure that this node and the target node don't attempt to reconnect at the same time)
			r := rand.Intn(maxConnectAttemptSleepDuration)
			time.Sleep(time.Duration(r) * time.Millisecond)
		}

		err := builder.host.Connect(ctx, peer.AddrInfo{ID: peerID})
		if err != nil {

			// if the connection was rejected due to invalid node id, skip the re-attempt
			if strings.Contains(err.Error(), "failed to negotiate security protocol") {
				return s, dialAddr, fmt.Errorf("invalid node id: %w", err)
			}

			// if the connection was rejected due to allowlisting, skip the re-attempt
			if errors.Is(err, swarm.ErrGaterDisallowedConnection) {
				return s, dialAddr, fmt.Errorf("target node is not on the approved list of nodes: %w", err)
			}

			errs = multierror.Append(errs, err)
			continue
		}

		// creates stream using stream factory
		s, err = builder.host.NewStream(ctx, peerID, protocolID)
		if err != nil {
			// if the stream creation failed due to invalid protocol id, skip the re-attempt
			if strings.Contains(err.Error(), "protocol not supported") {
				return nil, dialAddr, fmt.Errorf("remote node is running on a different spork: %w, protocol attempted: %s", err, protocolID)
			}
			errs = multierror.Append(errs, err)
			continue
		}

		break
	}

	if retries == maxAttempts {
		return s, dialAddr, errs
	}

	return s, dialAddr, nil
}
