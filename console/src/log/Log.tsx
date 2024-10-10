import { Icon } from "@synnaxlabs/media";
import { Log as Core, telem } from "@synnaxlabs/pluto";
import { id } from "@synnaxlabs/x";

import { Layout } from "@/layout";

export type LayoutType = "log";
export const LAYOUT_TYPE = "log";

export const Log: Layout.Renderer = () => {
  const t = telem.sourcePipeline("string", {
    connections: [
      {
        from: "valueStream",
        to: "stringifier",
      },
    ],
    segments: {
      valueStream: telem.streamChannelValue({ channel: 1048635 }),
      stringifier: telem.stringifyNumber({ precision: 2 }),
    },
    outlet: "stringifier",
  });
  return <Core.Log telem={t} />;
};

export const SELECTABLE: Layout.Selectable = {
  key: LAYOUT_TYPE,
  title: "Log",
  icon: <Icon.Log />,
  create: (key) => create({ key }),
};

export const create = (initial: Omit<Partial<Layout.State>, "type">): Layout.State => {
  const key = initial.key ?? id.id();
  return {
    key,
    name: "Log",
    icon: "Log",
    location: "mosaic",
    type: LAYOUT_TYPE,
    windowKey: key,
  };
};
