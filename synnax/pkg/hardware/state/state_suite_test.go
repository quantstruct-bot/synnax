// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package state_test

import (
	"context"
	"github.com/synnaxlabs/synnax/pkg/distribution"
	"github.com/synnaxlabs/synnax/pkg/distribution/mock"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx  = context.Background()
	_b   *mock.Builder
	dist distribution.Distribution
)

var _ = BeforeSuite(func() {
	_b = mock.NewBuilder()
	dist = _b.New(ctx)
})

var _ = AfterSuite(func() {
	Expect(_b.Close()).To(Succeed())
	Expect(_b.Cleanup()).To(Succeed())
})

func TestState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "State Suite")
}
