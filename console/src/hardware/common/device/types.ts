// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { z } from "zod";

const IDENTIFIER_MESSAGE = "Identifier must be between 2-12 characters";

export const identifierZ = z
  .string()
  .min(2, IDENTIFIER_MESSAGE)
  .max(12, IDENTIFIER_MESSAGE);

export type Identifier = z.infer<typeof identifierZ>;
