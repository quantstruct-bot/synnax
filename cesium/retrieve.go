package cesium

import (
	"context"
	"github.com/arya-analytics/cesium/internal/channel"
	"github.com/arya-analytics/cesium/internal/core"
	"github.com/arya-analytics/cesium/internal/persist"
	"github.com/arya-analytics/cesium/internal/segment"
	"github.com/arya-analytics/x/alamos"
	"github.com/arya-analytics/x/confluence"
	"github.com/arya-analytics/x/confluence/plumber"
	"github.com/arya-analytics/x/errutil"
	"github.com/arya-analytics/x/kv"
	"github.com/arya-analytics/x/query"
	"github.com/arya-analytics/x/queue"
	"github.com/arya-analytics/x/signal"
	"github.com/arya-analytics/x/telem"
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"time"
)

// |||||| CONFIGURATION ||||||

const (
	// retrievePersistMaxRoutines is the maximum number of goroutines the retrieve query persist.Persist can use.
	retrievePersistMaxRoutines = persist.DefaultNumWorkers
	// retrieveDebounceFlushInterval is the interval at which retrieve debounce queue will flush if the number of
	// retrieve operations is below the threshold.
	retrieveDebounceFlushInterval = 5 * time.Millisecond
	// retrieveDebounceFlushThreshold is the number of retrieve operations that must be in the debounce queue before
	// it flushes
	retrieveDebounceFlushThreshold = 100
)

type retrieveConfig struct {
	// exp is used to track metrics for the Retrieve query. See retrieveMetrics for more.
	exp alamos.Experiment
	// fs is the file system for reading segment data from.
	fs core.FS
	// kv is the key-value store for reading segment metadata from.
	kv kv.DB
	// logger is where retrieve operations will log their progress.
	logger *zap.Logger
	// debounce sets the debounce parameters for retrieve operations.
	// this is mostly here for optimizing performance under varied conditions.
	debounce queue.DebounceConfig
	// persist used for setting the parameters for persist.Persist thar reads
	// segment data from disk.
	persist persist.Config
}

func mergeRetrieveConfigDefaults(cfg *retrieveConfig) {

	// |||||| PERSIST ||||||

	if cfg.persist.NumWorkers == 0 {
		cfg.persist.NumWorkers = retrievePersistMaxRoutines
	}
	if cfg.persist.Logger == nil {
		cfg.persist.Logger = cfg.logger
	}

	// |||||| DEBOUNCE ||||||

	if cfg.debounce.Interval == 0 {
		cfg.debounce.Interval = retrieveDebounceFlushInterval
	}
	if cfg.debounce.Threshold == 0 {
		cfg.debounce.Threshold = retrieveDebounceFlushThreshold
	}
}

// |||||| STREAM ||||||

type ResponseVariant uint8

const (
	// AckResponse is a response that indicates that an iteration request was acknowledged.
	AckResponse ResponseVariant = iota + 1
	// DataResponse is a response that indicates that an iteration request returned data.
	DataResponse
)

// IteratorResponse is a response containing segments satisfying a Retrieve Query as well as any errors
// encountered during the retrieval.
type IteratorResponse struct {
	Counter  int
	Variant  ResponseVariant
	Ack      bool
	Err      error
	Segments []segment.Segment
}

// |||||| QUERY ||||||

type Retrieve struct {
	query.Query
	kve     kv.DB
	ops     confluence.Inlet[[]retrieveOperationUnary]
	metrics retrieveMetrics
	logger  *zap.Logger
}

// WhereChannels sets the channels to retrieve data for.
// If no keys are provided, will return an ErrInvalidQuery error.
func (r Retrieve) WhereChannels(keys ...ChannelKey) Retrieve { channel.SetKeys(r, keys...); return r }

// WhereTimeRange sets the time range to retrieve data from.
func (r Retrieve) WhereTimeRange(tr TimeRange) Retrieve { telem.SetTimeRange(r, tr); return r }

// Sync is an option that only applies to queries than open a StreamIterator.
// If set to true, the StreamIterator will wait for operations for a particular command
// to complete before returning.
func (r Retrieve) Sync() Retrieve { setSync(r, true); return r }

func (r Retrieve) SendAcks() Retrieve {
	setSendAcks(r, true)
	return r
}

// Stream streams all segments from the iterator out to the channel. Errors encountered
// during stream construction are returned immediately. Errors encountered during
// segment reads are returns as part of IteratorResponse.
func (r Retrieve) Stream(ctx context.Context, res chan<- IteratorResponse) (err error) {
	iter := r.Iterate()
	iter.OutTo(confluence.NewInlet(res))
	defer func() {
		err = errors.CombineErrors(iter.Close(), err)
	}()
	iter.SeekFirst()
	for iter.First(); iter.Valid(); iter.Next() {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
	}
	return err
}

func (r Retrieve) Iterate() StreamIterator {
	return newIteratorFromRetrieve(r)
}

// |||||| OPTIONS ||||||

// |||| SYNC ||||

const syncOptKey query.OptionKey = "sync"

func setSync(q query.Query, sync bool) {
	q.Set(syncOptKey, sync)
}

func getSync(q query.Query) bool {
	sync, ok := q.Get(syncOptKey)
	if !ok {
		return false
	}
	return sync.(bool)
}

// |||||| SEND ACKS |||||

const sendAcksOptKey query.OptionKey = "sendAcknowledgements"

func setSendAcks(q query.Query, sendAcks bool) {
	q.Set(sendAcksOptKey, sendAcks)
}

func getSendAcks(q query.Query) bool {
	sendAcks, ok := q.Get(sendAcksOptKey)
	if !ok {
		return false
	}
	return sendAcks.(bool)
}

// |||||| QUERY FACTORY ||||||

type retrieveFactory struct {
	kve     kv.DB
	logger  *zap.Logger
	metrics retrieveMetrics
	confluence.AbstractUnarySource[[]retrieveOperationUnary]
	confluence.EmptyFlow
}

func (r retrieveFactory) New() Retrieve {
	return Retrieve{
		Query:   query.New(),
		ops:     r.Out,
		kve:     r.kve,
		logger:  r.logger,
		metrics: r.metrics,
	}
}

func startRetrieve(
	ctx signal.Context,
	cfg retrieveConfig,
) (query.Factory[Retrieve], error) {
	mergeRetrieveConfigDefaults(&cfg)

	pipe := plumber.New()

	//// queue 'debounces' operations so that they can be flushed to disk in efficient
	//// batches.
	//plumber.SetSegment[[]retrieveOperationUnary, []retrieveOperationUnary](
	//	pipe,
	//	"queue",
	//	&queue.Debounce[retrieveOperationUnary]{Config: cfg.debounce},
	//)

	// batch groups operations into batches that optimize sequential IO.
	plumber.SetSegment[[]retrieveOperationUnary, []retrieveOperationSet](
		pipe,
		"batch",
		newRetrieveBatch(),
	)

	// persist executes batched operations on disk.
	plumber.SetSink[[]retrieveOperationSet](
		pipe,
		"persist",
		persist.New[core.FileKey, retrieveOperationSet](cfg.fs, cfg.persist),
	)

	queryFactory := &retrieveFactory{
		kve:     cfg.kv,
		metrics: newRetrieveMetrics(cfg.exp),
		logger:  cfg.logger,
	}

	plumber.SetSource[[]retrieveOperationUnary](pipe, "query", queryFactory)

	c := errutil.NewCatch()

	c.Exec(plumber.UnaryRouter[[]retrieveOperationUnary]{
		SourceTarget: "query",
		SinkTarget:   "batch",
		Capacity:     10,
	}.PreRoute(pipe))

	//methodCounter.Exec(plumber.UnaryRouter[[]retrieveOperationUnary]{
	//	SourceTarget: "queue",
	//	SinkTarget:   "batch",
	//	Capacity:     10,
	//}.PreRoute(pipe))

	c.Exec(plumber.UnaryRouter[[]retrieveOperationSet]{
		SourceTarget: "batch",
		SinkTarget:   "persist",
		Capacity:     10,
	}.PreRoute(pipe))

	if err := c.Error(); err != nil {
		panic(c.Error())
	}

	pipe.Flow(ctx)

	return queryFactory, nil
}
