// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { PropsWithChildren, useEffect } from "react";

import { z } from "zod";

import { Aether } from "@/aether";
import { control } from "@/telem/control/aether";

export interface ControllerProps
  extends z.input<typeof control.controllerStateZ>,
    PropsWithChildren {
  onStateChange?: (state: control.State) => void;
}

export const Controller = Aether.wrap<ControllerProps>(
  control.Controller.TYPE,
  ({ aetherKey, authority, acquireTrigger: propsTrigger, children, onStateChange }) => {
    const [{ path }, { state, acquireTrigger }, setState] = Aether.use({
      aetherKey,
      type: control.Controller.TYPE,
      schema: control.controllerStateZ,
      initialState: { authority, acquireTrigger: propsTrigger },
    });
    useEffect(() => {
      if (state != null) onStateChange?.(state);
    }, [state, onStateChange]);
    if (acquireTrigger !== propsTrigger)
      setState((p) => ({ ...p, acquireTrigger: propsTrigger }));

    return <Aether.Composite path={path}>{children}</Aether.Composite>;
  }
);
