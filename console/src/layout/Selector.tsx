// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import "@/layout/Selector.css";

import { Button, Eraser, Text } from "@synnaxlabs/pluto";
import { Align } from "@synnaxlabs/pluto/align";
import { type ReactElement } from "react";

import { CSS } from "@/css";
import { PlacerArgs, usePlacer } from "@/layout/hooks";
import { RendererProps } from "@/layout/slice";

export interface Selectable {
  key: string;
  title: string;
  icon: ReactElement;
  create: (layoutKey: string) => PlacerArgs;
}

export interface SelectorProps extends Align.SpaceProps, RendererProps {
  layouts?: Selectable[];
}

const Base = ({
  layoutKey,
  direction,
  layouts,
  visible: _,
  focused: __,
  ...props
}: SelectorProps): ReactElement => {
  const place = usePlacer();

  return (
    <Eraser.Eraser>
      <Align.Center
        className={CSS.B("vis-layout-selector")}
        size="large"
        {...props}
        wrap
      >
        <Text.Text level="h4" shade={6} weight={400}>
          Select a Component Type
        </Text.Text>
        <Align.Space
          direction="x"
          wrap
          style={{ width: "500px" }}
          justify="center"
          size={2.5}
        >
          {layouts?.map(({ key, title, icon, create }) => (
            <Button.Button
              key={key}
              variant="outlined"
              onClick={() => place(create(layoutKey))}
              startIcon={icon}
              style={{ flexBasis: "200px" }}
            >
              {title}
            </Button.Button>
          ))}
        </Align.Space>
      </Align.Center>
    </Eraser.Eraser>
  );
};

export const createSelectorComponent = (
  layouts: Selectable[],
): ((props: SelectorProps) => ReactElement) => {
  const C = (props: SelectorProps): ReactElement => (
    <Base layouts={layouts} {...props} />
  );
  C.displayName = "LayoutSelector";
  return C;
};
