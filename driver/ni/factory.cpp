// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

/// std
#include <vector>

/// external.
#include "glog/logging.h"
#include "driver/ni/ni.h"
#include "nlohmann/json.hpp"

/// internal
#include "nidaqmx/nidaqmx_prod.h"
#include "nisyscfg/nisyscfg_prod.h"

ni::Factory::Factory(
    const std::shared_ptr<DAQmx> &dmx,
    const std::shared_ptr<SysCfg> &syscfg
): dmx(dmx), syscfg(syscfg) {
}

bool ni::Factory::check_health(
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Task &task
) {
    if (this->dmx != nullptr && this->syscfg != nullptr) return true;
    ctx->set_state({
        .task = task.key,
        .variant = "error",
        .details = {
            "message",
            "Cannot create the task because the National Instruments DAQMX and System Configuration libraries are not installed on this system."
        }
    });
    return false;
}

std::shared_ptr<ni::Factory> ni::Factory::create() {
    auto [syscfg, syscfg_err] = SysCfgProd::load();
    auto [dmx, dmx_err] = DAQmxProd::load();
    if (syscfg_err || dmx_err) {
        LOG(ERROR) << syscfg_err;
        LOG(ERROR) << dmx_err;
    }
    return std::make_shared<ni::Factory>(dmx, syscfg);
}

std::pair<std::unique_ptr<task::Task>, bool> ni::Factory::configure_task(
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Task &task
) {
    if (!this->check_health(ctx, task)) return {nullptr, false};
    if (task.type == "ni_scanner")
        return {ni::ScannerTask::configure(this->syscfg, ctx, task), true};
    if (task.type == "ni_analog_read" || task.type == "ni_digital_read")
        return {ni::ReaderTask::configure(this->dmx, ctx, task), true};
    if (task.type == "ni_analog_write")
        return {ni::AnalogWriterTask::configure(this->dmx, ctx, task), true};
    if (task.type == "ni_digital_write")
        return {ni::DigitalWriterTask::configure(this->dmx, ctx, task), true};
    LOG(ERROR) << "[ni] Unknown task type: " << task.type << std::endl;
    return {nullptr, false};
}


std::vector<std::pair<synnax::Task, std::unique_ptr<task::Task> > >
ni::Factory::configure_initial_tasks(
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Rack &rack
) {
    std::vector<std::pair<synnax::Task, std::unique_ptr<task::Task> > > tasks;

    auto [existing, err] = rack.tasks.list();
    if (err) {
        LOG(ERROR) << "[ni] Failed to list existing tasks: " << err;
        return tasks;
    }

    bool hasScanner = false;
    for (const auto &t: existing)
        if (t.type == "ni_scanner") hasScanner = true;

    if (!hasScanner) {
        auto sy_task = synnax::Task(
            rack.key,
            "ni scanner",
            "ni_scanner",
            "",
            true
        );
        auto err = rack.tasks.create(sy_task);
        LOG(INFO) << "[ni] created scanner task with key: " << sy_task.key;
        if (err) {
            LOG(ERROR) << "[ni] Failed to create scanner task: " << err;
            return tasks;
        }
        auto [task, ok] = configure_task(ctx, sy_task);
        if (!ok) {
            LOG(ERROR) << "[ni] Failed to configure scanner task: " << err;
            return tasks;
        }
        tasks.emplace_back(
            std::pair<synnax::Task, std::unique_ptr<task::Task> >({
                sy_task,
                std::move(
                    task)
            }));
    }
    return tasks;
}
