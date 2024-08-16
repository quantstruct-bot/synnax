// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#pragma once

#include <utility>
#include <chrono>
#include <stdio.h>
#include <cassert> 
#include <regex>

#include "client/cpp/telem/telem.h"
#include "driver/ni/ni.h"

#include "nlohmann/json.hpp"
#include "glog/logging.h"

using json = nlohmann::json;

void ni::AnalogWriteSink::AnalogWriteSink(
    TaskHandle task_handle,
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Task &task)
    :   task_handle(task_handle),
        ctx(ctx),
        task(task),
        err_info({}) {

    auto config_parser = config::Parser(task.config);
    this->writer_config.task_name = task.name;
    this->parse_config(config_parser);
    if(!config_parser.ok()){
        this->log_error("failed to parse configuration for " + task.name );
        this->ctx->setState({
            .task = this->task.key,
            .variant = "error",
            .details = config_parser.error_json()
        });
        return;
    }
    auto breaker_config = breaker::Config{
        .name = task.name,
        .base_interval = 1 * SECOND,
        .max_retries = 20,
    }

}

int ni::AnalogWriteSink::init(){}

freighter::Error ni::AnalogWriteSink::write(synnax::Frame frame){}

// TODO: code dedup
freighter::Error ni::AnalogWriteSink::stop(){}

// TODO: code dedup
freighter::Error ni::AnalogWriteSink::start(){}

// TODO: code dedup
freighter::Error ni::AnalogWriteSink::cycle(){}

std::vector<synnax::ChannelKey> ni::AnalogWriteSink::get_cmd_channel_keys(){}

void ni::AnalogWriteSink::get_index_keys(){}

bool ni::AnalogWriteSink::ok(){}

void ni::AnalogWriteSink::jsonify_error(std::string){}

void ni::AnalogWriteSink::stoppedWithErr(const freighter::Error &err){}

void ni::AnalogWriteSink::log_error(stdLLstring err_msg);

// TODO: code dedup
void ni::AnalogWriteSink::clear_task(){
}

void ni::AnalogWriteSink::parse_config(config::Parser parser){}

void ni::AnalogWriteSink::check_ni_error(int32 error){}
