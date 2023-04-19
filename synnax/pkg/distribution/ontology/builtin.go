// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package ontology

import (
	"context"
	"github.com/synnaxlabs/synnax/pkg/distribution/ontology/schema"
	"github.com/synnaxlabs/x/iter"
	"github.com/synnaxlabs/x/observe"
)

// BuiltIn is a resource type that is built into the ontology.
const BuiltIn Type = "builtin"

var (
	// RootID is the root resource in the ontology. All other resources are reachable by
	// traversing the ontology from the root.
	RootID       = ID{Type: BuiltIn, Key: "root"}
	rootResource = Resource{ID: RootID, Name: "root"}
)

type builtinService struct {
	observe.Noop[[]Resource]
}

var _ Service = (*builtinService)(nil)

// Schema implements Service.
func (b *builtinService) Schema() *Schema {
	return &Schema{
		Type:   BuiltIn,
		Fields: map[string]schema.Field{},
	}
}

// RetrieveResource implements Service.
func (b *builtinService) RetrieveResource(_ context.Context, key string) (Resource, error) {
	switch key {
	case "root":
		return rootResource, nil
	default:
		panic("[ontology] - builtin entity not found")
	}
}

// Iterate implements Service.
func (b *builtinService) Iterate(f func(Resource) error) error {
	return iter.ForEachUntilError([]Resource{rootResource}, f)
}
