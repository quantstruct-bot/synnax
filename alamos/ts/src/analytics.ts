// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { Meta } from "@/meta";

type CaptureFn = (event: string, properties: Record<string, unknown>) => void;

export interface AnalyticsProps {
  capture: CaptureFn;
}

const ZERO_ANALYTICS_PROPS: AnalyticsProps = { capture: () => {} };

export class Analytics {
  private capture_: CaptureFn;
  meta: Meta = Meta.NOOP;

  static readonly NOOP = new Analytics(ZERO_ANALYTICS_PROPS);

  constructor(p: AnalyticsProps = ZERO_ANALYTICS_PROPS) {
    const { capture } = p;
    this.capture_ = capture;
  }

  capture(event: string, properties: Record<string, unknown>): void {
    if (this.meta.noop) return;
    this.capture_(`${this.meta.path}.${event}`, properties);
  }

  child(meta: Meta): Analytics {
    const a = new Analytics({ capture: this.capture_ });
    a.meta = meta;
    return a;
  }
}
