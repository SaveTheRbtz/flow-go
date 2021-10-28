// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package ingestion

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/atomic"

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/metrics"
	mockmodule "github.com/onflow/flow-go/module/mock"
	netint "github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/mocknetwork"
	"github.com/onflow/flow-go/utils/unittest"
)

func TestIngestionEngine(t *testing.T) {
	suite.Run(t, new(IngestionSuite))
}

type IngestionSuite struct {
	IngestionCoreSuite

	con    *mocknetwork.Conduit
	net    *mocknetwork.Network
	cancel context.CancelFunc

	ingest *Engine
}

func (s *IngestionSuite) SetupTest() {
	s.IngestionCoreSuite.SetupTest()

	s.con = &mocknetwork.Conduit{}

	// set up network module mock
	s.net = &mocknetwork.Network{}
	s.net.On("Register", engine.ReceiveGuarantees, mock.Anything).Return(
		func(channel netint.Channel, engine netint.Engine) netint.Conduit {
			return s.con
		},
		nil,
	)
<<<<<<< HEAD
	ref.On("Epochs").Return(suite.query)
	suite.query.On("Current").Return(suite.epoch)
	suite.epoch.On("Clustering").Return(clusters, nil)

	state.On("AtBlockID", mock.Anything).Return(ref)
	ref.On("Identity", mock.Anything).Return(
		func(nodeID flow.Identifier) *flow.Identity {
			identity, _ := suite.refIdentities.ByNodeID(nodeID)
			return identity
		},
		func(nodeID flow.Identifier) error {
			_, ok := suite.refIdentities.ByNodeID(nodeID)
			if !ok {
				return protocol.IdentityNotFoundError{NodeID: nodeID}
			}
			return nil
		},
	)

	// we use the first consensus node as our local identity
	me.On("NodeID").Return(suite.con1ID)
	me.On("NotMeFilter").Return(filter.Not(filter.HasNodeID(suite.con1ID)))

	// we need to return the head as it's also used as reference block
	headers.On("ByBlockID", head.ID()).Return(&head, nil)

	// only used for metrics, nobody cares
	pool.On("Size").Return(uint(0))

	ingest := &Engine{
		tracer:  tracer,
		metrics: metrics,
		spans:   metrics,
		mempool: metrics,
		state:   state,
		headers: headers,
		me:      me,
		pool:    pool,
		con:     con,
	}

	suite.head = &head
	suite.final = final
	suite.ref = ref
	suite.headers = headers
	suite.pool = pool
	suite.conduit = con
	suite.ingest = ingest
}

func (suite *IngestionSuite) TestOnGuaranteeNewFromCollection() {

	guarantee := suite.validGuarantee()

	// the guarantee is not part of the memory pool yet
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(true)

	// submit the guarantee as if it was sent by a collection node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().NoError(err, "should not error on new guarantee from collection node")

	// check that the guarantee has been added to the mempool
	suite.pool.AssertCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeUnstaked() {

	guarantee := suite.validGuarantee()

	// the guarantee is not part of the memory pool yet
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(true)

	// we are not staked
	suite.finalIdentities = suite.finalIdentities.Filter(filter.Not(filter.HasNodeID(suite.con1ID)))

	// submit the guarantee
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().NoError(err, "should not error on guarantee when unstaked")

	// the guarantee should be added to the mempool
	suite.pool.AssertCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeNewFromConsensus() {

	guarantee := suite.validGuarantee()

	// the guarantee is not part of the memory pool yet
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(true)

	// submit the guarantee as if it was sent by a consensus node
	err := suite.ingest.onGuarantee(suite.con1ID, guarantee)
	suite.Assert().NoError(err, "should not error on new guarantee from consensus node")

	// check that the guarantee has been added to the mempool
	suite.pool.AssertCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeOld() {

	guarantee := suite.validGuarantee()

	// the guarantee is part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(true)
	suite.pool.On("Add", guarantee).Return(true)

	// submit the guarantee as if it was sent by a collection node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().NoError(err, "should not error on old guarantee")

	// check that the guarantee has been added to the mempool
	suite.pool.AssertNotCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeNotAdded() {

	guarantee := suite.validGuarantee()

	// the guarantee is not already part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(false)

	// submit the guarantee as if it was sent by a collection node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().NoError(err, "should not error when guarantee was already added")

	// check that the guarantee has been added to the mempool
	suite.pool.AssertCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeNoGuarantor() {

	// create a guarantee signed by the collection node and referencing the
	// current head of the protocol state
	guarantee := suite.validGuarantee()
	guarantee.SignerIDs = nil

	// the guarantee is part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(false)

	// submit the guarantee as if it was sent by a consensus node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().Error(err, "should error with missing guarantor")
	suite.Assert().True(engine.IsInvalidInputError(err))

	// check that the guarantee has been added to the mempool
	suite.pool.AssertNotCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeInvalidRole() {

	// create a guarantee signed by the collection node and referencing the
	// current head of the protocol state
	guarantee := suite.validGuarantee()
	guarantee.SignerIDs = append(guarantee.SignerIDs, suite.execID)

	// the guarantee is part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(false)

	// submit the guarantee as if it was sent by a consensus node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().Error(err, "should error with missing guarantor")
	suite.Assert().True(engine.IsInvalidInputError(err))

	// check that the guarantee has been added to the mempool
	suite.pool.AssertNotCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeExpired() {

	// create an alternative block
	header := unittest.BlockHeaderFixture()
	header.Height = suite.head.Height - flow.DefaultTransactionExpiry - 1
	suite.headers.On("ByBlockID", header.ID()).Return(&header, nil)

	// create a guarantee signed by the collection node and referencing the
	// current head of the protocol state
	guarantee := suite.validGuarantee()
	guarantee.ReferenceBlockID = header.ID()

	// the guarantee is part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(false)

	// submit the guarantee as if it was sent by a consensus node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().Error(err, "should error with expired collection")
	suite.Assert().True(engine.IsOutdatedInputError(err))

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

func (suite *IngestionSuite) TestOnGuaranteeInvalidGuarantor() {

	// create a guarantee signed by the collection node and referencing the
	// current head of the protocol state
	guarantee := suite.validGuarantee()
	guarantee.SignerIDs = append(guarantee.SignerIDs, unittest.IdentifierFixture())

	// the guarantee is not part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(false)

	// submit the guarantee as if it was sent by a collection node
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().Error(err, "should error with invalid guarantor")
	suite.Assert().True(engine.IsInvalidInputError(err))

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

// test that just after an epoch boundary we still accept guarantees from collectors
// in clusters from the previous epoch (and collectors which are leaving the network
// at this epoch boundary).
func (suite *IngestionSuite) TestOnGuaranteeEpochEnd() {

	// in the finalized state the collectors has 0 stake but is not ejected
	// this is what happens when we finalize the final block of the epoch during
	// which this node requested to unstake
	colID, ok := suite.finalIdentities.ByNodeID(suite.collID)
	suite.Require().True(ok)
	colID.Stake = 0

	guarantee := suite.validGuarantee()

	// the guarantee is not part of the memory pool
	suite.pool.On("Has", guarantee.ID()).Return(false)
	suite.pool.On("Add", guarantee).Return(true)

	// submit the guarantee as if it was sent by the collection node which
	// is leaving at the current epoch boundary
	err := suite.ingest.onGuarantee(suite.collID, guarantee)
	suite.Assert().NoError(err, "should not error with collector from ending epoch")

	// check that the guarantee has been added to the mempool
	suite.pool.AssertExpectations(suite.T())

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
=======

	// setup my own identity
	me := &mockmodule.Local{}
	me.On("NodeID").Return(s.conID) // we use the first consensus node as our local identity

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	signalerCtx, _ := irrecoverable.WithSignaler(ctx)

	metrics := metrics.NewNoopCollector()
	ingest, err := New(unittest.Logger(), metrics, s.net, me, s.core)
	require.NoError(s.T(), err)
	s.ingest = ingest
	s.ingest.Start(signalerCtx)
	<-s.ingest.Ready()
}

func (s *IngestionSuite) TearDownTest() {
	s.cancel()
	<-s.ingest.Done()
>>>>>>> 02def6ea5f686f5a6c5cfddcc230cc3e66e1d802
}

func (s *IngestionSuite) TestSubmittingMultipleEntries() {
	originID := s.collID
	count := uint64(15)

	processed := atomic.NewUint64(0)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < int(count); i++ {
			guarantee := s.validGuarantee()
			s.pool.On("Has", guarantee.ID()).Return(false)
			s.pool.On("Add", guarantee).Run(func(args mock.Arguments) {
				processed.Add(1)
			}).Return(true)

<<<<<<< HEAD
	suite.pool.AssertNotCalled(suite.T(), "Add", guarantee)

	// we should not propagate the guarantee
	suite.conduit.AssertNotCalled(suite.T(), "Multicast", guarantee, mock.Anything, mock.Anything)
	suite.conduit.AssertNotCalled(suite.T(), "Publish", guarantee, mock.Anything)
}

// validGuarantee returns a valid collection guarantee based on the suite state.
func (suite *IngestionSuite) validGuarantee() *flow.CollectionGuarantee {
	guarantee := unittest.CollectionGuaranteeFixture()
	guarantee.SignerIDs = []flow.Identifier{suite.collID}
	guarantee.ReferenceBlockID = suite.head.ID()
	return guarantee
=======
			// execute the vote submission
			_ = s.ingest.Process(engine.ProvideCollections, originID, guarantee)
		}
		wg.Done()
	}()

	wg.Wait()

	require.Eventually(s.T(), func() bool {
		return processed.Load() == count
	}, time.Millisecond*200, time.Millisecond*20)

	s.pool.AssertExpectations(s.T())
>>>>>>> 02def6ea5f686f5a6c5cfddcc230cc3e66e1d802
}
