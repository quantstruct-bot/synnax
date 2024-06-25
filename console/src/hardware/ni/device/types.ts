// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { task } from "@synnaxlabs/client";
import { z } from "zod";

// VENDOR

const VENDORS = ["NI", "other"] as const;

export const vendorZ = z.enum(VENDORS);

export type Vendor = z.infer<typeof vendorZ>;

// DEVICE DIGEST

const propertiesDigestZ = z.object({
  key: z.string(),
  enriched: z.boolean(),
});

export type PropertiesDigest = z.infer<typeof propertiesDigestZ>;

export const configurablePropertiesZ = z.object({
  name: z.string(),
  identifier: z.string().min(2).max(12),
});

const commandStatePairZ = z.object({
  command: z.number(),
  state: z.number(),
});

const enrichedPropertiesDigestZ = propertiesDigestZ.extend({
  identifier: z.string().min(2).max(12),
  analogInput: z.object({
    portCount: z.number(),
    index: z.number(),
    channels: z.record(z.string(), z.number()),
  }),
  analogOutput: z.object({
    portCount: z.number(),
  }),
  digitalInputOutput: z.object({
    portCount: z.number(),
    lineCounts: z.number().array(),
  }),
  digitalInput: z.object({
    portCount: z.number(),
    lineCounts: z.number().array(),
  }),
  digitalOutput: z.object({
    portCount: z.number(),
    lineCounts: z.number().array(),
    stateIndex: z.number(),
    channels: z.record(z.string(), commandStatePairZ),
  }),
});

export type Properties = z.infer<typeof enrichedPropertiesDigestZ>;
