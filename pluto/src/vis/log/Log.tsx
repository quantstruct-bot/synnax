// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { box } from "@synnaxlabs/x";
import { ReactElement, useCallback, useEffect, useRef } from "react";
import { z } from "zod";

import { Aether } from "@/aether";
import { useCombinedRefs, useSyncedRef } from "@/hooks";
import { useMemoDeepEqualProps } from "@/memo";
import { Canvas } from "@/vis/canvas";
import { log } from "@/vis/log/aether";

export interface LogProps extends Omit<z.input<typeof log.logState>, "region"> {}

export const Log = Aether.wrap<LogProps>(
  "Log",
  ({ aetherKey, font, color, telem }): ReactElement | null => {
    const memoProps = useMemoDeepEqualProps({ font, color, telem });
    const elRef = useRef<HTMLDivElement | null>(null);
    const [, { scrollPosition, totalHeight }, setState] = Aether.use({
      aetherKey,
      type: log.Log.TYPE,
      schema: log.logState,
      initialState: {
        region: box.ZERO,
        scrollPosition: null,
        totalHeight: 0,
        ...memoProps,
      },
    });

    const scrollPosRef = useSyncedRef(scrollPosition);
    const snapRef = useRef<number | null>(null);
    useEffect(() => {
      if (elRef.current == null || snapRef.current != null) return;
      elRef.current.scrollTop = elRef.current.scrollHeight ?? 0;
    }, [totalHeight]);

    useEffect(() => {
      setState((s) => ({ ...s, ...memoProps }));
    }, [memoProps, setState]);

    const resizeRef = Canvas.useRegion(
      useCallback(
        (b) => {
          if (snapRef.current == null && elRef.current != null)
            elRef.current.scrollTop = elRef.current.scrollHeight;
          setState((s) => ({ ...s, region: b }));
        },
        [setState],
      ),
    );

    const combinedRef = useCombinedRefs(elRef, resizeRef);
    return (
      <div style={{ height: "100%", paddingTop: "1rem", paddingLeft: "1rem" }}>
        <div
          ref={combinedRef}
          style={{ height: "100%", overflowY: "auto" }}
          onScroll={(e) => {
            const el = e.target as HTMLDivElement;
            const elScrollPos = el.scrollTop + el.clientHeight;
            if (elScrollPos == el.scrollHeight) {
              snapRef.current = null;
              if (scrollPosRef.current != null)
                setState((s) => ({ ...s, scrollPosition: null }));
              return;
            }
            if (snapRef.current == null) snapRef.current = el.scrollHeight;
            setState((s) => ({ ...s, scrollPosition: elScrollPos - snapRef.current }));
          }}
        >
          <div style={{ height: totalHeight }} />
        </div>
      </div>
    );
  },
);
