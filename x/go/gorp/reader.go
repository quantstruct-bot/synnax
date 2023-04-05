// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package gorp

import (
	"github.com/samber/lo"
	"github.com/synnaxlabs/x/kv"
)

// Iterator provides a simple wrapper around a kv.Write that decodes a byte-value
// before returning it to the caller. It provides no abstracted utilities for the
// iteration itself, and is focused only on maintaining a nearly identical interface to
// the underlying iterator. To create a new Iterator, call NewIterator.
type Iterator[E any] struct {
	kv.Iterator
	error error
	options
}

// NewIterator wraps the provided iterator. All valid calls to iter.Value are
// decoded into the entry type E.
func NewIterator[E any](wrapped kv.Iterator, opts ...Option) *Iterator[E] {
	return &Iterator[E]{Iterator: wrapped, options: newOptions(opts...)}
}

// Value returns the decoded value from the iterator. Iter.Alive must be true
// for calls to return a valid value.
func (k *Iterator[E]) Value() (entry E) { k.BindValue(&entry); return entry }

func (k *Iterator[E]) BindValue(entry *E) {
	k.error = k.decoder.Decode(k.Iterator.Value(), entry)
}

func (k *Iterator[E]) Error() error {
	return lo.Ternary(k.error != nil, k.error, k.Iterator.Error())
}

func (k *Iterator[E]) Valid() bool {
	return lo.Ternary(k.error != nil, false, k.Iterator.Valid())
}
