#  Copyright 2023 sy Labs, Inc.
#
#  Use of this software is governed by the Business Source License included in the file
#  licenses/BSL.txt.
#
#  As of the Change Date specified in that file, in accordance with the Business Source
#  License, use of this software will be governed by the Apache License, Version 2.0,
#  included in the file licenses/APL.txt.

import pytest

import synnax as sy
import numpy as np


class TestChannelClient:
    @pytest.fixture(scope="class")
    @pytest.mark.channel
    def two_channels(self, client: sy.Synnax) -> list[sy.Channel]:
        return client.channels.create(
            [
                sy.Channel(
                    name="test",
                    rate=1 * sy.Rate.HZ,
                    data_type=sy.DataType.FLOAT64,
                ),
                sy.Channel(
                    name="test2",
                    rate=1 * sy.Rate.HZ,
                    data_type=sy.DataType.FLOAT64,
                ),
            ]
        )

    @pytest.mark.channel
    def test_write_read(self, client: sy.Synnax):
        """Should create a channel and write then read from it"""
        channel = client.channels.create(
            sy.Channel(
                name="test",
                rate=1 * sy.Rate.HZ,
                data_type=sy.DataType.INT64,
            )
        )
        input = np.array([1, 2, 3, 4, 5], dtype=np.int64)
        start = 1 * sy.TimeSpan.SECOND
        channel.write(start, input)
        data, tr = channel.read(start, (start + len(input)) * sy.TimeSpan.SECOND)
        assert len(input) == len(data)
        assert input[0] == data[0]
        assert tr.start == start
        assert tr.end == start + (len(input) - 1) * sy.TimeSpan.SECOND + 1

    @pytest.mark.channel
    def test_create_list(self, two_channels: list[sy.Channel]):
        """Should create a list of valid channels"""
        assert len(two_channels) == 2
        for channel in two_channels:
            assert channel.name.startswith("test")
            assert channel.key != ""

    @pytest.mark.channel
    def test_create_single(self, client: sy.Synnax):
        """Should create a single valid channel"""
        channel = client.channels.create(
            sy.Channel(
                name="test",
                rate=1 * sy.Rate.HZ,
                data_type=sy.DataType.FLOAT64,
            )
        )
        assert channel.name == "test"
        assert channel.key != ""
        assert channel.data_type == sy.DataType.FLOAT64
        assert channel.rate == 1 * sy.Rate.HZ

    @pytest.mark.channel
    def test_create_from_kwargs(self, client: sy.Synnax):
        """Should create a single valid channel"""
        channel = client.channels.create(
            name="test",
            rate=1 * sy.Rate.HZ,
            data_type=sy.DataType.FLOAT64,
        )
        assert channel.name == "test"
        assert channel.key != ""
        assert channel.data_type == sy.DataType.FLOAT64
        assert channel.rate == 1 * sy.Rate.HZ

    @pytest.mark.channel
    def test_create_invalid_nptype(self, client: sy.Synnax):
        """Should throw a Validation Error when passing invalid numpy data type"""
        with pytest.raises(TypeError):
            client.channels.create(data_type=np.csingle)

    @pytest.mark.channel
    def test_retrieve_by_key(
        self, two_channels: list[sy.Channel], client: sy.Synnax
    ) -> None:
        """Should retrieve channels using a list of keys"""
        res_channels = client.channels.retrieve(
            [channel.key for channel in two_channels]
        )
        assert len(res_channels) == 2
        for i, channel in enumerate(res_channels):
            assert two_channels[i].key == channel.key
            assert isinstance(two_channels[i].data_type.density, sy.Density)

    @pytest.mark.channel
    def test_retrieve_by_key_not_found(self, client: sy.Synnax):
        """Should raise QueryError when key not found"""
        with pytest.raises(sy.QueryError):
            client.channels.retrieve("1-100000")

    @pytest.mark.channel
    def test_retrieve_by_list_of_names(
        self, two_channels: list[sy.Channel], client: sy.Synnax
    ) -> None:
        """Should retrieve channels using list of names"""
        res_channels = client.channels.retrieve(["test", "test2"])
        assert len(res_channels) >= 2
        for channel in res_channels:
            assert channel.name in ["test", "test2"]

    @pytest.mark.channel
    def test_retrieve_by_single_name(
        self, channel: sy.Channel, client: sy.Synnax
    ) -> None:
        """Should retrieve channel using a name string"""
        res_channel = client.channels.retrieve("test")
        assert res_channel.name == "test"

    @pytest.mark.channel
    def test_retrieve_list_not_found(self, client: sy.Synnax):
        """Should retrieve an empty list when can't find channels"""
        fake_names = ["fake1", "fake2", "fake3"]
        results = client.channels.retrieve(fake_names)
        assert len(results) == 0

    @pytest.mark.channel
    def test_retrieve_include_not_found_true(self, client: sy.Synnax):
        """Should return list of unfound channels when include_not_found is true"""
        fake_names = ["test", "test2", "fake3"]
        res_channels, not_found = client.channels.retrieve(
            fake_names, include_not_found=True
        )
        assert len(res_channels) == 2
        print(type(not_found))
        assert len(not_found) == 3
