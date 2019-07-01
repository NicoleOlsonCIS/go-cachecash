package client

import (
	"context"
	"errors"
	"net"
	"testing"

	cachecash "github.com/cachecashproject/go-cachecash"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/ed25519"
)

func TestSchedulerTestSuite(t *testing.T) {
	suite.Run(t, new(SchedulerTestSuite))
}

type SchedulerTestSuite struct {
	suite.Suite

	l *logrus.Logger
}

func (suite *SchedulerTestSuite) SetupTest() {
	l := logrus.New()
	suite.l = l
	suite.l.SetLevel(logrus.DebugLevel)
}

func (suite *SchedulerTestSuite) newMock() (*client, *publisherMock) {
	mock := newPublisherMock()
	return &client{
		l:             suite.l,
		publisherConn: mock,
		cacheConns:    map[cacheID]cacheConnection{},
	}, mock
}

type CROptions struct {
	bundles    uint64
	objectSize uint64
}

func (suite *SchedulerTestSuite) newContentResponse(options ...CROptions) *ccmsg.ContentResponse {
	// Defaults
	opts := CROptions{
		bundles:    0,
		objectSize: 512,
	}
	// Merge explicit choices
	for _, opt := range options {
		if opt.bundles != 0 {
			opts.bundles = opt.bundles
		}
		if opt.objectSize != 0 {
			opts.objectSize = opt.objectSize
		}
	}

	bundles := []*ccmsg.TicketBundle{}

	for i := uint64(0); i < opts.bundles; i++ {
		chunkIdx := uint64(i * 2)
		bundles = append(bundles, &ccmsg.TicketBundle{
			Remainder: &ccmsg.TicketBundleRemainder{
				PuzzleInfo: &ccmsg.ColocationPuzzleInfo{
					Goal: []byte{},
				},
			},
			TicketRequest: []*ccmsg.TicketRequest{
				{
					ChunkIdx: chunkIdx,
					InnerKey: &ccmsg.BlockKey{
						Key: []byte{},
					},
					CachePublicKey: &ccmsg.PublicKey{
						PublicKey: []byte{},
					},
				},
				{
					ChunkIdx: chunkIdx + 1,
					InnerKey: &ccmsg.BlockKey{
						Key: []byte{},
					},
					CachePublicKey: &ccmsg.PublicKey{
						PublicKey: []byte{},
					},
				},
			},
			CacheInfo: []*ccmsg.CacheInfo{
				suite.newCache(net.ParseIP("192.0.2.1"), 1001, []byte{0, 1, 2, 3, 4}),
				suite.newCache(net.ParseIP("192.0.2.2"), 1002, []byte{5, 6, 7, 8, 9}),
			},
			Metadata: &ccmsg.ObjectMetadata{
				ObjectSize: opts.objectSize,
				ChunkSize:  128,
			},
		})
	}
	return &ccmsg.ContentResponse{
		Bundles: bundles,
	}
}

func (suite *SchedulerTestSuite) newCache(ip net.IP, port uint32, pubkey []byte) *ccmsg.CacheInfo {
	return &ccmsg.CacheInfo{
		Pubkey: &ccmsg.PublicKey{
			PublicKey: pubkey,
		},
		Addr: &ccmsg.NetworkAddress{
			Inetaddr: ip,
			Port:     port,
		},
	}
}

func (suite *SchedulerTestSuite) TestSchedulerOneBundle() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 0, Chunks: 2}
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 2, Chunks: 2}
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	group = <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestSchedulerZeroBundles() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 0, Chunks: 2}
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 2, Chunks: 2}
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	group = <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestSchedulerAllBundlesAtOnce() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 2}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 0, Chunks: 2}
	bundleCompletions <- BundleOutcome{Outcome: Completed, ChunkOffset: 2, Chunks: 2}
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	group = <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestSchedulerError() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return((*ccmsg.ContentResponse)(nil), errors.New("this is an error")).Once()

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.Nil(t, group.bundle)
	assert.Equal(t, "failed to fetch chunk-group at chunk offset 0: failed to request bundle from publisher: this is an error", group.err.Error())
	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestRequestBundle() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(&ccmsg.ContentResponse{
		Bundles: []*ccmsg.TicketBundle{},
	}, nil).Once()

	resp, err := cl.requestBundles(context.Background(), "/", 0)
	assert.Nil(t, err, "failed to get bundle")
	assert.Equal(t, []*ccmsg.TicketBundle{}, resp)
}

func (suite *SchedulerTestSuite) TestRequestBundleError() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return((*ccmsg.ContentResponse)(nil), errors.New("this is an error")).Once()

	resp, err := cl.requestBundles(context.Background(), "/", 0)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func (suite *SchedulerTestSuite) TestCacheConnectionError() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("newCacheConnection", cl.l, "192.0.2.1:1001",
		ed25519.PublicKey(([]byte)("\x00\x01\x02\x03\x04"))).Return(
		(*cacheMock)(nil), errors.New("cache connection failure")).Once()

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.NotNil(t, group.err)
	assert.Nil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestChangedChunkCount() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(CROptions{bundles: 1, objectSize: 1024}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.Nil(t, group.err)
	assert.NotNil(t, group.bundle)

	group = <-queue
	assert.NotNil(t, group.err)
	assert.Nil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestSchedulerClientErrorsOneBundle() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	// Today, attempting to retry shuts down the scheduler
	bundleCompletions <- BundleOutcome{Outcome: Retry, ChunkOffset: 0, Chunks: 2}
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.NotNil(t, group.err)
	assert.Nil(t, group.bundle)

	assert.Zero(t, len(queue))
}

func (suite *SchedulerTestSuite) TestSchedulerClientDefersOneBundle() {
	t := suite.T()
	cl, mock := suite.newMock()

	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      0,
		RangeEnd:        0,
		BacklogDepth:    map[string]uint64{},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.On("GetContent", &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            "/",
		RangeBegin:      256,
		RangeEnd:        0,
		BacklogDepth: map[string]uint64{
			"\x00\x01\x02\x03\x04": 0x1,
			"\x05\x06\x07\x08\x09": 0x1,
		},
	}).Return(suite.newContentResponse(CROptions{bundles: 1}), nil).Once()
	mock.makeNewCacheCall(cl.l, "192.0.2.1:1001", "\x00\x01\x02\x03\x04")
	mock.makeNewCacheCall(cl.l, "192.0.2.2:1002", "\x05\x06\x07\x08\x09")

	queue := make(chan *fetchGroup, 128)
	bundleCompletions := make(chan BundleOutcome, 128)
	// Today, attempting to defer shuts down the scheduler
	bundleCompletions <- BundleOutcome{Outcome: Deferred, ChunkOffset: 0, Chunks: 2}
	cl.schedule(context.Background(), "/", queue, bundleCompletions)

	group := <-queue
	assert.NotNil(t, group.err)
	assert.Nil(t, group.bundle)

	assert.Zero(t, len(queue))
}
