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
	"github.com/cockroachdb/errors"
	"github.com/synnaxlabs/cesium/internal/core"
	"github.com/synnaxlabs/cesium/internal/index"
	"github.com/synnaxlabs/cesium/internal/unary"
	"github.com/synnaxlabs/x/errutil"
	"github.com/synnaxlabs/x/query"
	"github.com/synnaxlabs/x/telem"
	"sync"
)

var (
	// ChannelNotFound is returned when a channel or a range of data cannot be found in the DB.
	ChannelNotFound  = errors.Wrap(query.NotFound, "[DB] - channel not found")
	ErrDiscontinuous = index.ErrDiscontinuous
)

type (
	Channel    = core.Channel
	ChannelKey = core.ChannelKey
	Frame      = core.Frame
)

func NewFrame(keys []core.ChannelKey, arrays []telem.Array) Frame { return core.NewFrame(keys, arrays) }

type DB struct {
	*options
	mu    sync.RWMutex
	relay *relay
	dbs   map[uint32]unary.DB
}

// Write implements DB.
func (db *DB) Write(ctx context.Context, start telem.TimeStamp, frame Frame) error {
	_, span := db.T.Debug(ctx, "write")
	defer span.End()
	w, err := db.NewWriter(ctx, WriterConfig{Start: start, Channels: frame.Keys})
	if err != nil {
		return err
	}
	w.Write(frame)
	w.Commit()
	return w.Close()
}

// WriteArray implements DB.
func (db *DB) WriteArray(ctx context.Context, start telem.TimeStamp, key core.ChannelKey, arr telem.Array) error {
	return db.Write(ctx, start, core.NewFrame([]core.ChannelKey{key}, []telem.Array{arr}))
}

// Read implements DB.
func (db *DB) Read(ctx context.Context, tr telem.TimeRange, keys ...core.ChannelKey) (frame Frame, err error) {
	iter, err := db.NewIterator(IteratorConfig{Channels: keys, Bounds: tr})
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

// Close implements DB.
func (db *DB) Close() error {
	c := errutil.NewCatch(errutil.WithAggregation())
	for _, u := range db.dbs {
		c.Exec(u.Close)
	}
	return nil
}
