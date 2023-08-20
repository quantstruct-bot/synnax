// Copyrght 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { ReactElement } from "react";

import { CrudeXY } from "@synnaxlabs/x";
import { Handle, Position } from "reactflow";

import { Color } from "@/color";
import { CSS, ColorSwatch, SwatchProps, Input, InputNumberProps, Space } from "@/core";
import { Tank, TankProps } from "@/core/vis/Tank/Tank";
import { componentRenderProp } from "@/util/renderProp";
import {
  PIDElementFormProps,
  PIDElementSpec,
  StatefulPIDElementProps,
} from "@/vis/PID/PIDElement";

import "@/vis/PID/TankPIDElement/TankPIDElement.css";

export interface TankPIDElementProps extends Omit<TankProps, "telem"> {
  label: string;
}

const { Left, Right, Top, Bottom } = Position;

const TankPIDElement = ({
  selected,
  editable,
  position,
  className,
  ...props
}: StatefulPIDElementProps<TankPIDElementProps>): ReactElement => {
  return (
    <div className={CSS(className, CSS.B("tank-pid-element"), CSS.selected(selected))}>
      {editable && (
        <>
          <Handle position={Left} type="source" id="a" style={{ top: "25%" }} />
          <Handle position={Left} type="source" id="c" style={{ top: "75%" }} />
          <Handle position={Right} type="source" id="e" style={{ top: "25%" }} />
          <Handle position={Right} type="source" id="g" style={{ top: "75%" }} />
          <Handle position={Top} type="source" id="j" />
          <Handle position={Bottom} type="source" id="l" />
        </>
      )}
      <Tank {...props}></Tank>
    </div>
  );
};

const DIMENSIONS_DRAG_SCALE: CrudeXY = { y: 2, x: 0.25 };

const TankPIDElementForm = ({
  value,
  onChange,
}: PIDElementFormProps<TankPIDElementProps>): ReactElement => {
  const handleWidthChange = (width: number): void =>
    onChange({ ...value, dimensions: { ...value.dimensions, width } });
  const handleHeightChange = (height: number): void =>
    onChange({ ...value, dimensions: { ...value.dimensions, height } });
  const handleLabelChange = (label: string): void => onChange({ ...value, label });
  const handleColorChange = (color: Color.Color): void =>
    onChange({ ...value, color: color.hex });

  return (
    <>
      <Input.Item<string>
        label="Label"
        value={value.label}
        onChange={handleLabelChange}
      />

      <Space direction="horizonatal">
        <Input.Item<number, number, InputNumberProps>
          label="Width"
          value={value.dimensions.width}
          onChange={handleWidthChange}
          dragScale={DIMENSIONS_DRAG_SCALE}
        >
          {componentRenderProp(Input.Number)}
        </Input.Item>

        <Input.Item<number, number, InputNumberProps>
          label="Height"
          value={value.dimensions.height}
          onChange={handleHeightChange}
          dragScale={DIMENSIONS_DRAG_SCALE}
        >
          {componentRenderProp(Input.Number)}
        </Input.Item>
        <Input.Item<Color.Crude, Color.Color, SwatchProps>
          label="Color"
          onChange={handleColorChange}
          value={value.color}
        >
          {/* @ts-expect-error */}
          {componentRenderProp(ColorSwatch)}
        </Input.Item>
      </Space>
    </>
  );
};

const TankPIDElementPreview = (): ReactElement => {
  return <Tank color={ZERO_PROPS.color} dimensions={{ width: 30, height: 40 }}></Tank>;
};

const ZERO_PROPS = {
  dimensions: { width: 100, height: 250 },
  label: "Tank",
  color: "#ffffff",
};

export const TankPIDElementSpec: PIDElementSpec<TankPIDElementProps> = {
  type: "tank",
  title: "Tank",
  initialProps: ZERO_PROPS,
  Element: TankPIDElement,
  Form: TankPIDElementForm,
  Preview: TankPIDElementPreview,
};
