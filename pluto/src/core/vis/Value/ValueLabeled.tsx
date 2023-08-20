// Copyrght 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { ReactElement, useState } from "react";

import { Box, Direction, XY } from "@synnaxlabs/x";

import { Color, CrudeColor } from "@/color";
import { CSS } from "@/css";
import { useResize } from "@/hooks";
import { PackProps, Space, Text } from "@/core/std";
import { Theming } from "@/theming/main";
import { UseTypographyReturn } from "@/theming/main/font";
import { ValueCore, ValueCoreProps } from "@/core/vis/Value/ValueCore";

import "@/core/vis/Value/ValueLabeled.css";

export interface ValueLabeledProps
  extends Omit<ValueCoreProps, "box">,
    Omit<PackProps, "color" | "onChange"> {
  position?: XY;
  zoom?: number;
  label: string;
  onLabelChange?: (label: string) => void;
  color?: CrudeColor;
  textColor?: CrudeColor;
}

export const ValueLabeled = ({
  label,
  onLabelChange,
  level = "p",
  direction = "y",
  position,
  className,
  children,
  textColor,
  color,
  zoom = 1,
  ...props
}: ValueLabeledProps): ReactElement => {
  const font = Theming.useTypography(level);
  const [box, setBox] = useState<Box>(Box.ZERO);

  const valueBoxHeight = (font.lineHeight + 2) * font.baseSize + 2;
  const resizeRef = useResize(setBox, {});

  const adjustedBox = adjustBox(
    new Direction(direction),
    zoom,
    box,
    valueBoxHeight,
    font,
    position
  );

  return (
    <Space
      className={CSS(className, CSS.B("value-labeled"))}
      align="center"
      ref={resizeRef}
      direction={direction}
      {...props}
    >
      <Text.MaybeEditable value={label} onChange={onLabelChange} level={level} />
      <div
        className={CSS.B("value")}
        style={{
          height: valueBoxHeight,
          borderColor: Color.cssString(color),
        }}
      >
        {children}
        <ValueCore color={textColor} level={level} {...props} box={adjustedBox} />
      </div>
    </Space>
  );
};

const adjustBox = (
  direction: Direction,
  zoom: number,
  box: Box,
  valueBoxHeight: number,
  font: UseTypographyReturn,
  position?: XY
): Box => {
  if (direction.isX) {
    return new Box(
      (position?.x ?? box.left) + box.width / zoom - 100,
      position?.y ?? box.top,
      100,
      valueBoxHeight
    );
  }
  return new Box(
    position?.x ?? box.left,
    (position?.y ?? box.top) + box.height / zoom - valueBoxHeight,
    box.width / zoom,
    valueBoxHeight
  );
};
