// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package cesium

import (
	"context"
	"github.com/synnaxlabs/cesium/internal/core"
	"github.com/synnaxlabs/cesium/internal/unary"
	"github.com/synnaxlabs/cesium/internal/virtual"
	"github.com/synnaxlabs/x/confluence"
	"github.com/synnaxlabs/x/errors"
	"github.com/synnaxlabs/x/telem"
	"io"
	"sync"
	"sync/atomic"
)

type (
	Channel    = core.Channel
	ChannelKey = core.ChannelKey
	Frame      = core.Frame
)

func NewFrame(keys []core.ChannelKey, series []telem.Series) Frame {
	return core.NewFrame(keys, series)
}

var (
	errDBClosed        = core.EntityClosed("cesium.db")
	ErrChannelNotFound = core.ErrChannelNotFound
)

// DB is a cesium database for storing channel-based time series data.
type DB struct {
	*options
	// relay is used for real-time streaming of written sensor values.
	relay *relay
	// mu is the primary database lock. It guards regions of the database that cannot
	// be accessed or modified concurrently.
	mu struct {
		sync.RWMutex
		// unaryDBs tracks the databases the fixed sample size, unary channels that store
		// data within a file system.
		unaryDBs map[ChannelKey]unary.DB
		// virtualDBs tracks the databases for the virtual channels that are used for
		// streaming purposes only.
		virtualDBs map[ChannelKey]virtual.DB
		// controlDigests is used to manage the lifecycle of the control handoff tracking
		// mechanism i.e. a writer that communicates control changes to the relay.
		controlDigests struct {
			// key is the key of the channel assigned to the control digests.
			key ChannelKey
			// inlet is the inlet of the control digests writer.
			inlet confluence.Inlet[WriterRequest]
			// outlet is the outlet of the control digests writer.
			outlet confluence.Outlet[WriterResponse]
		}
	}
	// closed marks whether the DB has been closed or not.
	closed *atomic.Bool
	// shutdown shuts down database goroutines spawned through internal processes
	// (e.g. GC).
	shutdown io.Closer
}

// Write writes the frame to database at the specified start time.
func (db *DB) Write(
	ctx context.Context,
	start telem.TimeStamp,
	frame Frame,
) (err error) {
	if db.closed.Load() {
		return errDBClosed
	}
	_, span := db.T.Bench(ctx, "write")
	defer func() { err = span.EndWith(err) }()
	var w *Writer
	w, err = db.OpenWriter(ctx, WriterConfig{Start: start, Channels: frame.Keys})
	if err != nil {
		return
	}
	defer func() { err = w.Close() }()
	w.Write(frame)
	w.Commit()
	return
}

// WriteArray writes a series into the specified channel at the specified start time.
func (db *DB) WriteArray(
	ctx context.Context,
	key core.ChannelKey,
	start telem.TimeStamp,
	series telem.Series,
) error {
	if db.closed.Load() {
		return errDBClosed
	}
	return db.Write(ctx, start, core.NewFrame([]core.ChannelKey{key}, []telem.Series{series}))
}

// Read reads from the database at the specified time range and outputs a frame.
func (db *DB) Read(ctx context.Context, tr telem.TimeRange, keys ...core.ChannelKey) (frame Frame, err error) {
	if db.closed.Load() {
		return frame, errDBClosed
	}
	_, span := db.T.Bench(ctx, "read")
	defer func() { err = span.EndWith(err) }()
	iter, err := db.OpenIterator(IteratorConfig{Channels: keys, Bounds: tr})
	if err != nil {
		return
	}
	defer func() { err = iter.Close() }()
	if !iter.SeekFirst() {
		return
	}
	for iter.Next(telem.TimeSpanMax) {
		frame = frame.AppendFrame(iter.Value())
	}
	return
}

// Close closes the database.
// Close is not safe to call with any other DB methods concurrently.
// Note that if this method is called while writers are still open on channels in the
// database, a deadlock is caused since the signal context is closed while the writers
// attempt to send to relay.
// If there is an error in closing the cesium database, the database will be marked as
// closed regardless of whether an error occurred.
func (db *DB) Close() error {
	if !db.closed.CompareAndSwap(false, true) {
		return nil
	}

	c := errors.NewCatcher(errors.WithAggregation())
	// Crucial to close control controlDigests here before closing the signal context so
	// writes can still use the signal context to send frames to relay.
	//
	// This function acquires the mutex lock internally, so there's no need to lock
	// it here.
	db.closeControlDigests()
	// Shut down without locking mutex to allow existing goroutines (e.g. GC) that
	// require a mutex lock to exit.
	c.Exec(db.shutdown.Close)
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, u := range db.mu.unaryDBs {
		c.Exec(u.Close)
	}

	return c.Error()
}
