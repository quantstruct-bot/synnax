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
	"io/fs"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/synnaxlabs/x/errors"
	"github.com/synnaxlabs/x/signal"
	"github.com/synnaxlabs/x/telem"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

type GCConfig struct {
	// MaxGoroutine is the maximum number of GoRoutines that can be launched for
	// each try of garbage collection.
	MaxGoroutine int64

	// GCTryInterval is the interval of time between two tries of garbage collection
	// are started.
	GCTryInterval time.Duration

	// GCThreshold is the minimum tombstone proportion of the Filesize to trigger a GC.
	// Must be in (0, 1].
	// Note: Setting this value to 0 will have NO EFFECT as it is the default value.
	// instead, set it to a very small number greater than 0.
	// [OPTIONAL] Default: 0.2
	GCThreshold float32
}

var DefaultGCConfig = GCConfig{
	MaxGoroutine:  10,
	GCTryInterval: 30 * time.Second,
	GCThreshold:   0.2,
}

func keyToDirName(ch ChannelKey) string {
	return strconv.Itoa(int(ch))
}

const deleteChannelSuffix = "-DELETE-"

func deleteChannelKey(ch ChannelKey) string {
	return keyToDirName(ch) + deleteChannelSuffix + uuid.New().String()
}

// DeleteChannel deletes a channel by its key.
// This method returns an error if there are other channels depending on the current
// channel, or if the current channel is being written to or read from.
// DeleteChannel is idempotent.
func (db *DB) DeleteChannel(ch ChannelKey) error {
	if db.closed.Load() {
		return errDBClosed
	}
	// Rename the file first, so we can avoid hogging the mutex while deleting the
	// directory, which may take a longer time.
	// Rename the file to have a random suffix in case the channel is repeatedly created
	// and deleted.
	var (
		oldName = keyToDirName(ch)
		newName = deleteChannelKey(ch)
	)
	if err := (func() error {
		db.mu.Lock()
		defer db.mu.Unlock()
		if err := db.removeChannel(ch); err != nil {
			return err
		}
		err := db.fs.Rename(oldName, newName)
		return errors.Skip(err, fs.ErrNotExist)
	})(); err != nil {
		return err
	}

	return db.fs.Remove(newName)
}

// DeleteChannels deletes many channels by their keys.
// This operation is not guaranteed to be atomic: it is possible some channels in chs
// are deleted and some are not.
func (db *DB) DeleteChannels(chs []ChannelKey) (err error) {
	if db.closed.Load() {
		return errDBClosed
	}
	var (
		indexChannels       = make([]ChannelKey, 0, len(chs))
		directoriesToRemove = make([]string, 0, len(chs))
	)
	db.mu.Lock()
	// This 'defer' statement does a best-effort removal of all renamed directories
	// to ensure that all DBs deleted from db.unaryDBs and db.virtualDBs are also deleted
	// on FS.
	defer func() {
		db.mu.Unlock()
		c := errors.NewCatcher(errors.WithAggregation())
		for _, name := range directoriesToRemove {
			c.Exec(func() error { return db.fs.Remove(name) })
		}
		err = errors.Combine(err, c.Error())
	}()

	// Do a pass first to remove all non-index channels
	for _, ch := range chs {
		uDB, uOk := db.unaryDBs[ch]

		if !uOk || uDB.Channel().IsIndex {
			if uDB.Channel().IsIndex {
				indexChannels = append(indexChannels, ch)
			}
			continue
		}

		if err = db.removeChannel(ch); err != nil {
			return
		}

		// Rename the files first, so we can avoid hogging the mutex while deleting
		// the directory, which may take a longer time.
		oldName := keyToDirName(ch)
		newName := deleteChannelKey(ch)
		if err = db.fs.Rename(oldName, newName); err != nil {
			return
		}

		directoriesToRemove = append(directoriesToRemove, newName)
	}

	// Do another pass to remove all index channels
	for _, ch := range indexChannels {
		if err = db.removeChannel(ch); err != nil {
			return
		}

		oldName := keyToDirName(ch)
		newName := deleteChannelKey(ch)
		err = db.fs.Rename(oldName, newName)
		if err != nil {
			return
		}

		directoriesToRemove = append(directoriesToRemove, newName)
	}

	return
}

// removeChannel removes ch from db.unaryDBs or db.virtualDBs. If the key does not exist
// or if there is an open entity on the specified database.
func (db *DB) removeChannel(ch ChannelKey) error {
	if uDB, uOk := db.unaryDBs[ch]; uOk {
		if uDB.Channel().IsIndex {
			for otherDBKey, otherDB := range db.unaryDBs {
				if otherDBKey != ch && otherDB.Channel().Index == uDB.Channel().Key {
					return errors.Newf(
						"cannot delete channel %v "+
							"because it indexes data in channel %v",
						uDB.Channel(),
						otherDB.Channel(),
					)
				}
			}
		}

		if err := uDB.Close(); err != nil {
			return err
		}
		delete(db.unaryDBs, ch)
		return nil
	}
	if vDB, vOk := db.virtualDBs[ch]; vOk {
		if err := vDB.Close(); err != nil {
			return err
		}
		delete(db.virtualDBs, ch)
		return nil
	}

	return nil
}

// DeleteTimeRange deletes a time range of data in the database in the given channels
// This method return an error if the channel to be deleted is an index channel and
// there are other channels depending on it in the time range.
// DeleteTimeRange is idempotent, but when the channel does not exist, it returns
// ErrChannelNotFound.
func (db *DB) DeleteTimeRange(
	ctx context.Context,
	chs []ChannelKey,
	tr telem.TimeRange,
) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	var (
		indexChannels = make([]ChannelKey, 0, len(chs))
		dataChannels  = make([]ChannelKey, 0, len(chs))
	)

	for _, ch := range chs {
		udb, uok := db.unaryDBs[ch]
		if !uok {
			if vdb, vok := db.virtualDBs[ch]; vok {
				return errors.Newf(
					"cannot delete time range from virtual channel %v",
					vdb.Channel(),
				)
			}
			return errors.Wrapf(ErrChannelNotFound, "channel key %d not found", ch)
		}

		// Cannot delete an index channel that other channels rely on.
		if udb.Channel().IsIndex {
			indexChannels = append(indexChannels, ch)
			continue
		}

		dataChannels = append(dataChannels, ch)
	}

	for _, ch := range dataChannels {
		udb := db.unaryDBs[ch]
		if err := udb.Delete(ctx, tr); err != nil {
			return err
		}
	}

	for _, ch := range indexChannels {
		udb := db.unaryDBs[ch]
		// Cannot delete an index channel that other channels rely on.
		for otherDBKey, otherDB := range db.unaryDBs {
			if otherDBKey == ch || otherDB.Channel().Index != ch {
				continue
			}
			hasOverlap, err := otherDB.HasDataFor(ctx, tr)
			if err != nil || hasOverlap {
				return errors.Newf(
					"cannot delete index channel %v "+
						"with channel %v depending on it on the time range %s",
					udb.Channel(),
					otherDB.Channel(),
					tr,
				)
			}
		}

		if err := udb.Delete(ctx, tr); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) garbageCollect(ctx context.Context, maxRoutines int64) (err error) {
	var (
		_, span      = db.T.Debug(ctx, "garbage_collect")
		sCtx, cancel = signal.Isolated()
		sem          = semaphore.NewWeighted(maxRoutines)
	)
	db.mu.RLock()
	defer func() {
		db.mu.RUnlock()
		// If error is not nil here,it means we failed to acquire the semaphore,
		// so we cancel the wrapped context to stop all the goroutines.
		if err != nil {
			cancel()
		}
		err = errors.Combine(err, sCtx.Wait())
		span.End()
	}()
	for _, uDB := range db.unaryDBs {
		if err = sem.Acquire(ctx, 1); err != nil {
			return err
		}
		uDB := uDB
		sCtx.Go(func(ctx context.Context) error {
			defer sem.Release(1)
			return uDB.GarbageCollect(ctx)
		}, signal.RecoverWithErrOnPanic())
	}
	return
}

func (db *DB) startGC(sCtx signal.Context) {
	signal.GoTick(sCtx, db.gcCfg.GCTryInterval, func(ctx context.Context, time time.Time) error {
		if err := db.garbageCollect(ctx, db.gcCfg.MaxGoroutine); err != nil {
			db.L.Error("garbage collection error", zap.Error(err))
		}
		return nil
	}, signal.WithRetryOnPanic(10), signal.RecoverWithoutErrOnPanic())
}
