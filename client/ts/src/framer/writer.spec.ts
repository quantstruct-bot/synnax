// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { DataType, Rate, TimeRange, TimeSpan, TimeStamp } from "@synnaxlabs/x/telem";
import { describe, expect, test } from "vitest";

import { type channel } from "@/channel";
import { ALWAYS_INDEX_PERSIST_ON_AUTO_COMMIT, WriterMode } from "@/framer/writer";
import { newClient } from "@/setupspecs";
import { randomSeries } from "@/util/telem";

const client = newClient();

const newChannel = async (): Promise<channel.Channel> =>
  await client.channels.create({
    name: "test",
    leaseholder: 1,
    rate: Rate.hz(1),
    dataType: DataType.FLOAT64,
  });

describe("Writer", () => {
  describe("Writer", () => {
    test("basic write", async () => {
      const ch = await newChannel();
      const writer = await client.openWriter({ start: 0, channels: ch.key });
      try {
        await writer.write(ch.key, randomSeries(10, ch.dataType));
        await writer.commit();
      } finally {
        await writer.close();
      }
      expect(true).toBeTruthy();
    });
    test("write to unknown channel key", async () => {
      const ch = await newChannel();
      const writer = await client.openWriter({ start: 0, channels: ch.key });
      await expect(
        writer.write("billy bob", randomSeries(10, DataType.FLOAT64)),
      ).rejects.toThrow("Channel billy bob not found");
      await writer.close();
    });
    test("stream when mode is set ot persist only", async () => {
      const ch = await newChannel();
      const stream = await client.openStreamer(ch.key);
      const writer = await client.openWriter({
        start: 0,
        channels: ch.key,
        mode: WriterMode.PersistOnly,
      });
      try {
        await writer.write(ch.key, randomSeries(10, ch.dataType));
      } finally {
        await writer.close();
      }
      // Simulating a timeout.
      const v = await Promise.race([
        stream.read(),
        new Promise((resolve) => setTimeout(() => resolve(123), 250)),
      ]);
      expect(v).toEqual(123);
    });
    test("write with auto commit on", async () => {
      const ch = await newChannel();
      const writer = await client.openWriter({
        start: 0,
        channels: ch.key,
        enableAutoCommit: true,
      });
      try {
        await writer.write(ch.key, randomSeries(10, ch.dataType));
      } finally {
        await writer.close();
      }
      expect(true).toBeTruthy();

      const f = await client.read(new TimeRange(0, TimeStamp.seconds(10)), ch.key);
      expect(f.length).toEqual(10);
    });
    test("write with auto commit and alwaysPersist", async () => {
      const ch = await newChannel();
      const writer = await client.openWriter({
        start: 0,
        channels: ch.key,
        enableAutoCommit: true,
        autoIndexPersistInterval: ALWAYS_INDEX_PERSIST_ON_AUTO_COMMIT,
      });
      try {
        await writer.write(ch.key, randomSeries(10, ch.dataType));
      } finally {
        await writer.close();
      }
      expect(true).toBeTruthy();
    });
    test("write with auto commit and a set interval", async () => {
      const ch = await newChannel();
      const writer = await client.openWriter({
        start: 0,
        channels: ch.key,
        enableAutoCommit: true,
        autoIndexPersistInterval: TimeSpan.milliseconds(100),
      });
      try {
        await writer.write(ch.key, randomSeries(10, ch.dataType));
      } finally {
        await writer.close();
      }
      expect(true).toBeTruthy();
    });
    test("write with errOnUnauthorized", async () => {
      const ch = await newChannel();
      const w1 = await client.openWriter({
        start: 0,
        channels: ch.key,
      });

      const w2 = await client.openWriter({
        start: 0,
        channels: ch.key,
        errOnUnauthorized: true,
      })
      await w2.write(ch.key, randomSeries(10, ch.dataType))
      await expect(
        w2.close()
      ).rejects.toThrow("unauthorized")

      await w1.close()
    })
  });
  describe("Client", () => {
    test("Client - basic write", async () => {
      const ch = await newChannel();
      const data = randomSeries(10, ch.dataType);
      await client.write(TimeStamp.seconds(1), ch.key, data);
      const res = await client.read(TimeRange.MAX, ch.key);
      expect(res.length).toEqual(data.length);
      expect(res.data).toEqual(data);
    });
  });
});
