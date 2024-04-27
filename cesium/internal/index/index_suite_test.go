// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package index_test

import (
	"context"
	"github.com/synnaxlabs/x/testutil"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx                       = context.Background()
	rootPath                  = "index-testdata"
	fileSystems, cleanUp, err = testutil.FileSystems()
)

func TestIndex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Index Suite")
}

var _ = BeforeSuite(func() { Expect(err).ToNot(HaveOccurred()) })

var _ = AfterSuite(func() { Expect(cleanUp()).To(Succeed()) })
