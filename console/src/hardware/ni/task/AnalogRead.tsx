// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { NotFoundError, QueryError, task as clientTask } from "@synnaxlabs/client";
import { Icon } from "@synnaxlabs/media";
import {
  Align,
  Button,
  Form,
  Header,
  List,
  Menu,
  Status,
  Synnax,
  Text,
} from "@synnaxlabs/pluto";
import { binary, deep, id, primitiveIsZero, unique } from "@synnaxlabs/x";
import { useMutation } from "@tanstack/react-query";
import { type FC, type ReactElement, useCallback, useState } from "react";

import { CSS } from "@/css";
import { Common } from "@/hardware/common";
import { Device } from "@/hardware/ni/device";
import {
  ANALOG_INPUT_FORMS,
  SelectChannelTypeField,
} from "@/hardware/ni/task/ChannelForms";
import {
  AI_CHANNEL_TYPE_NAMES,
  type AIChannel,
  type AIChannelType,
  ANALOG_READ_TYPE,
  type AnalogReadConfig,
  analogReadConfigZ,
  type AnalogReadDetails,
  type AnalogReadType,
  ZERO_AI_CHANNELS,
  ZERO_ANALOG_READ_PAYLOAD,
} from "@/hardware/ni/task/types";
import { useCopyToClipboard } from "@/hooks/useCopyToClipboard";
import { type Layout } from "@/layout";

export const ANALOG_READ_LAYOUT: Common.Task.LayoutBaseState = {
  ...Common.Task.LAYOUT,
  type: ANALOG_READ_TYPE,
  name: ZERO_ANALOG_READ_PAYLOAD.name,
  icon: "Logo.NI",
  key: ANALOG_READ_TYPE,
};

export const ANALOG_READ_SELECTABLE: Layout.Selectable = {
  key: ANALOG_READ_TYPE,
  title: "NI Analog Read Task",
  icon: <Icon.Logo.NI />,
  create: (key) => ({ ...ANALOG_READ_LAYOUT, key }),
};

interface ChannelDetailsProps {
  selectedChannelIndex?: number | null;
}

const ChannelDetails = ({
  selectedChannelIndex,
}: ChannelDetailsProps): ReactElement => {
  const ctx = Form.useContext();
  const copy = useCopyToClipboard();
  const handleCopyChannelDetails = () => {
    if (selectedChannelIndex == null) return;
    copy(
      binary.JSON_CODEC.encodeString(
        ctx.get(`config.channels.${selectedChannelIndex}`).value,
      ),
      "Channel details",
    );
  };

  return (
    <Align.Space className={CSS.B("channel-form")} direction="y" grow>
      <Header.Header level="h4">
        <Header.Title weight={500} wrap={false}>
          Details
        </Header.Title>
        <Header.Actions>
          <Button.Icon
            tooltip="Copy channel details as JSON"
            tooltipLocation="left"
            variant="text"
            onClick={handleCopyChannelDetails}
          >
            <Icon.JSON style={{ color: "var(--pluto-gray-l7)" }} />
          </Button.Icon>
        </Header.Actions>
      </Header.Header>
      <Align.Space className={CSS.B("details")}>
        {selectedChannelIndex != null && (
          <ChannelForm selectedChannelIndex={selectedChannelIndex} />
        )}
      </Align.Space>
    </Align.Space>
  );
};

interface ChannelFormProps {
  selectedChannelIndex: number;
}

const ChannelForm = ({ selectedChannelIndex }: ChannelFormProps): ReactElement => {
  const prefix = `config.channels.${selectedChannelIndex}`;
  const type = Form.useFieldValue<AIChannelType>(`${prefix}.type`, true);
  if (type == null) return <></>;
  const TypeForm = ANALOG_INPUT_FORMS[type];
  if (selectedChannelIndex == -1) return <></>;
  return (
    <>
      <Align.Space direction="y" className={CSS.B("channel-form-content")} empty>
        <SelectChannelTypeField path={prefix} inputProps={{ allowNone: false }} />
        <TypeForm prefix={prefix} />
      </Align.Space>
    </>
  );
};

interface ChannelListProps {
  path: string;
  onSelect: (keys: string[], index: number) => void;
  selected: string[];
  snapshot?: boolean;
  onTare: (keys: number[]) => void;
  state?: clientTask.State<{ running?: boolean; message?: string }>;
}

const availablePortFinder = (channels: AIChannel[]): (() => number) => {
  const exclude = new Set(channels.map((v) => v.port));
  return () => {
    let i = 0;
    while (exclude.has(i)) i++;
    exclude.add(i);
    return i;
  };
};

const ChannelList = ({
  path,
  snapshot,
  selected,
  onSelect,
  state,
  onTare,
}: ChannelListProps): ReactElement => {
  const { value, push, remove } = Form.useFieldArray<AIChannel>({ path });
  const handleAdd = (): void => {
    const key = id.id();
    push({
      ...deep.copy(ZERO_AI_CHANNELS.ai_voltage),
      port: availablePortFinder(value)(),
      key,
    });
    onSelect([key], value.length);
  };
  const menuProps = Menu.useContextMenu();
  return (
    <Align.Space className={CSS.B("channels")} grow empty>
      <Common.Task.ChannelListHeader onAdd={handleAdd} snapshot={snapshot} />
      <Menu.ContextMenu
        menu={({ keys }: Menu.ContextMenuMenuProps): ReactElement => (
          <Common.Task.ChannelListContextMenu
            path={path}
            keys={keys}
            value={value}
            remove={remove}
            onSelect={onSelect}
            snapshot={snapshot}
            onTare={onTare}
            allowTare={state?.details?.running === true}
            onDuplicate={(indices) => {
              const pf = availablePortFinder(value);
              push(
                indices.map((i) => ({
                  ...deep.copy(value[i]),
                  channel: 0,
                  port: pf(),
                  key: id.id(),
                })),
              );
            }}
          />
        )}
        {...menuProps}
      >
        <List.List<string, AIChannel>
          data={value}
          emptyContent={
            <Common.Task.ChannelListEmptyContent
              onAdd={handleAdd}
              snapshot={snapshot}
            />
          }
        >
          <List.Selector<string, AIChannel>
            value={selected}
            allowNone={false}
            allowMultiple
            onChange={(keys, { clickedIndex }) =>
              clickedIndex != null && onSelect(keys, clickedIndex)
            }
            replaceOnSingle
          >
            <List.Core<string, AIChannel> grow>
              {({ key: i, ...props }) => (
                <ChannelListItem
                  {...props}
                  key={i}
                  path={path}
                  snapshot={snapshot}
                  state={state}
                  onTare={(key) => onTare([key])}
                />
              )}
            </List.Core>
          </List.Selector>
        </List.List>
      </Menu.ContextMenu>
    </Align.Space>
  );
};

const ChannelListItem = ({
  path: basePath,
  snapshot = false,
  onTare,
  state,
  ...props
}: List.ItemProps<string, AIChannel> & {
  path: string;
  snapshot?: boolean;
  onTare?: (channelKey: number) => void;
  state?: clientTask.State<{ running?: boolean; message?: string }>;
}): ReactElement => {
  const ctx = Form.useContext();
  const path = `${basePath}.${props.index}`;
  const portValid = Form.useFieldValid(`${path}.port`);

  // TODO: fix bug so I can refactor this to original code
  const channels = Form.useChildFieldValues<AIChannel[]>({ path: basePath });
  if (channels == null || props?.index == null) return <></>;
  const childValues = channels[props.index];
  // const childValues = Form.useChildFieldValues<AIChan>({ path, optional: true });

  if (childValues == null) return <></>;
  const showTareButton = childValues.channel != null && onTare != null;
  const tareIsDisabled =
    !childValues.enabled || snapshot || state?.details?.running !== true;
  return (
    <List.ItemFrame
      {...props}
      entry={childValues}
      justify="spaceBetween"
      align="center"
    >
      <Align.Space direction="y" size="small">
        <Align.Space direction="x">
          <Text.Text
            level="p"
            weight={500}
            shade={6}
            style={{ width: "3rem" }}
            color={portValid ? undefined : "var(--pluto-error-z)"}
          >
            {childValues.port}
          </Text.Text>
          <Text.Text level="p" weight={500} shade={9}>
            {AI_CHANNEL_TYPE_NAMES[childValues.type]}
          </Text.Text>
        </Align.Space>
      </Align.Space>
      <Align.Pack direction="x" align="center" size="small">
        {showTareButton && (
          <Common.Task.TareButton
            disabled={tareIsDisabled}
            onClick={() => onTare(childValues.channel)}
          />
        )}
        <Common.Task.EnableDisableButton
          value={childValues.enabled}
          onChange={(v) => ctx.set(`${path}.enabled`, v)}
          snapshot={snapshot}
        />
      </Align.Pack>
    </List.ItemFrame>
  );
};

const TaskForm: FC<
  Common.Task.FormProps<AnalogReadConfig, AnalogReadDetails, AnalogReadType>
> = ({ task, taskState }) => {
  const [selectedChannels, setSelectedChannels] = useState<string[]>(
    task.config.channels.length ? [task.config.channels[0].key] : [],
  );
  const [selectedChannelIndex, setSelectedChannelIndex] = useState<number | null>(
    task.config.channels.length > 0 ? 0 : null,
  );
  const client = Synnax.use();
  const handleException = Status.useExceptionHandler();
  const handleTare = useMutation({
    onError: (e) => handleException(e, "Failed to tare channels"),
    mutationFn: async (keys: number[]) => {
      if (client == null) return;
      if (!(task instanceof clientTask.Task)) return;
      await task?.executeCommand("tare", { keys });
    },
  }).mutate;
  return (
    <>
      <Align.Space direction="x" className={CSS.B("task-properties")}>
        <Align.Space direction="x">
          <Form.NumericField
            label="Sample Rate"
            path="config.sampleRate"
            inputProps={{ endContent: "Hz" }}
          />
          <Form.NumericField
            label="Stream Rate"
            path="config.streamRate"
            inputProps={{ endContent: "Hz" }}
          />
          <Form.SwitchField path="config.dataSaving" label="Data Saving" />
        </Align.Space>
      </Align.Space>
      <Align.Space
        direction="x"
        className={CSS.B("channel-form-container")}
        bordered
        rounded
        grow
        empty
      >
        <ChannelList
          snapshot={task?.snapshot}
          path="config.channels"
          selected={selectedChannels}
          onSelect={useCallback(
            (v, i) => {
              setSelectedChannels(v);
              setSelectedChannelIndex(i);
            },
            [setSelectedChannels, setSelectedChannelIndex],
          )}
          onTare={handleTare}
          state={taskState}
        />
        <ChannelDetails selectedChannelIndex={selectedChannelIndex} />
      </Align.Space>
    </>
  );
};

export const AnalogReadTask = Common.Task.wrapForm(TaskForm, {
  configSchema: analogReadConfigZ,
  type: ANALOG_READ_TYPE,
  zeroPayload: ZERO_ANALOG_READ_PAYLOAD,
  onConfigure: async (client, config) => {
    const devices = unique.unique(config.channels.map((c) => c.device));
    for (const devKey of devices) {
      const dev = await client.hardware.devices.retrieve<Device.Properties>(devKey);
      dev.properties = Device.enrich(dev.model, dev.properties);
      let modified = false;
      let shouldCreateIndex = primitiveIsZero(dev.properties.analogInput.index);
      if (!shouldCreateIndex)
        try {
          await client.channels.retrieve(dev.properties.analogInput.index);
        } catch (e) {
          if (NotFoundError.matches(e)) shouldCreateIndex = true;
          else throw e;
        }
      if (shouldCreateIndex) {
        modified = true;
        const aiIndex = await client.channels.create({
          name: `${dev.properties.identifier}_ai_time`,
          dataType: "timestamp",
          isIndex: true,
        });
        dev.properties.analogInput.index = aiIndex.key;
        dev.properties.analogInput.channels = {};
      }
      const toCreate: AIChannel[] = [];
      for (const channel of config.channels) {
        if (channel.device !== dev.key) continue;
        // check if the channel is in properties
        const exKey = dev.properties.analogInput.channels[channel.port.toString()];
        if (primitiveIsZero(exKey)) toCreate.push(channel);
        else
          try {
            await client.channels.retrieve(exKey.toString());
          } catch (e) {
            if (QueryError.matches(e)) toCreate.push(channel);
            else throw e;
          }
      }
      if (toCreate.length > 0) {
        modified = true;
        const channels = await client.channels.create(
          toCreate.map((c) => ({
            name: `${dev.properties.identifier}_ai_${c.port}`,
            dataType: "float32",
            index: dev.properties.analogInput.index,
          })),
        );
        channels.forEach(
          (c, i) =>
            (dev.properties.analogInput.channels[toCreate[i].port.toString()] = c.key),
        );
      }
      if (modified) await client.hardware.devices.create(dev);
      config.channels.forEach((c) => {
        if (c.device !== dev.key) return;
        c.channel = dev.properties.analogInput.channels[c.port.toString()];
      });
    }
    return config;
  },
});
