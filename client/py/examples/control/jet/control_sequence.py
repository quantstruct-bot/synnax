#  Copyright 2024 Synnax Labs, Inc.
#
#  Use of this software is governed by the Business Source License included in the file
#  licenses/BSL.txt.
#
#  As of the Change Date specified in that file, in accordance with the Business Source
#  License, use of this software will be governed by the Apache License, Version 2.0,
#  included in the file licenses/APL.txt.

import dataclasses
import synnax as sy
from synnax.control.controller import Controller

from common import (
    FUEL_PUMP_CMD,
    FUEL_VALVE_1_CMD,
    FUEL_VALVE_2_CMD,
    FUEL_RES_VALVE_1_CMD,
    FUEL_RES_VALVE_2_CMD,
    AIR_VALVE_1_CMD,
    AIR_VALVE_2_CMD,
    BLEED_VALVE_CMD,
    SPARK_PLUG_CMD,
    STARTER_MOTOR_CMD,
    IGNITION_CMD,
    N1_SPEED,
    N2_SPEED,
    FLAME,
    COMBUSTION_TC_1,
    EXHAUST_TC,
)

@dataclasses.dataclass
class EngineParameters:
    n1_idle_target: float = 2000.0  # RPM
    n2_idle_target: float = 5000.0  # RPM
    n1_max: float = 5000.0  # RPM
    n2_max: float = 15000.0  # RPM
    max_exhaust_tc: float = 800.0  # Celsius
    max_combustion_tc_1: float = 1000.0  # Celsius
    startup_timeout: float = 30.0  # seconds
    cooldown_timeout: float = 60.0  # seconds

client = sy.Synnax()

auto_logs = client.channels.create(
    name="auto_logs",
    data_type=sy.DataType.STRING,
    virtual=True,
    retrieve_if_name_exists=True,
)

start_auto_cmd = client.channels.create(
    name="start_auto_cmd",
    data_type=sy.DataType.UINT8,
    virtual=True,
    retrieve_if_name_exists=True,
)

def log(aut: Controller, msg: str):
    s = f"{sy.TimeStamp.now().datetime().strftime('%H:%M:%S.%f')}  {msg}"
    aut.set(auto_logs.key, s)
    print(s)

def execute_startup(aut: Controller, params: EngineParameters) -> bool:
    """Execute engine startup sequence"""
    log(aut, "Waiting for operator")
    aut.wait_until(lambda s: s.get(start_auto_cmd.name, 0) == 1)

    log(aut, "Starting engine startup sequence")
    
    # Open air valves
    aut.set({
        AIR_VALVE_1_CMD: True,
        AIR_VALVE_2_CMD: True,
    })
    aut.sleep(2)
    
    # Start fuel pump and open fuel reservoir valves
    aut.set({
        FUEL_PUMP_CMD: True,
        FUEL_RES_VALVE_1_CMD: True,
        FUEL_RES_VALVE_2_CMD: True,
    })
    aut.sleep(2)
    
    # Start motor and wait for N1 to reach 20% of idle
    log(aut, "Starting motor and waiting for N1 spool-up")
    aut[STARTER_MOTOR_CMD] = True
    
    if not aut.wait_until(
        lambda s: s[N2_SPEED] >= params.n2_idle_target * 0.16,
        timeout=params.startup_timeout
    ):
        log(aut, "N2 failed to reach target speed")
        return False
    
    # Begin ignition sequence
    log(aut, "Beginning ignition sequence")
    aut.set({
        SPARK_PLUG_CMD: True,
        IGNITION_CMD: True,
        FUEL_VALVE_1_CMD: True,
        FUEL_VALVE_2_CMD: True,
    })
    
    # Wait for flame
    if not aut.wait_until(
        lambda s: s[FLAME] > 0,
        timeout=params.startup_timeout
    ):
        log(aut, "Failed to achieve ignition")
        return False
    
    log(aut, "Engine lit - waiting for idle")
    
    # Wait for idle speeds
    if not aut.wait_until(
        lambda s: s[N1_SPEED] >= params.n1_idle_target and s[N2_SPEED] >= params.n2_idle_target,
        timeout=params.startup_timeout
    ):
        log(aut, "Failed to reach idle")
        return False
    
    # Shutdown starter motor and ignition
    aut.set({
        STARTER_MOTOR_CMD: False,
        SPARK_PLUG_CMD: False,
        IGNITION_CMD: False,
    })
    
    log(aut, "Engine startup complete")
    return True

def execute_shutdown(aut: Controller, params: EngineParameters) -> bool:
    """Execute engine shutdown sequence"""
    log(aut, "Beginning engine shutdown sequence")
    
    # Close fuel valves
    aut.set({
        FUEL_VALVE_1_CMD: False,
        FUEL_VALVE_2_CMD: False,
        FUEL_PUMP_CMD: False,
    })
    
    # Wait for flame out
    if not aut.wait_until(
        lambda s: not s[FLAME],
        timeout=params.startup_timeout
    ):
        log(aut, "Flame failed to extinguish")
        return False
    
    log(aut, "Flame out - beginning cooldown")
    
    # Wait for temperature cooldown
    if not aut.wait_until(
        lambda s: s[COMBUSTION_TC_1] <= 100 and s[EXHAUST_TC] <= 100,
        timeout=params.cooldown_timeout
    ):
        log(aut, "Temperatures remained high")
        return False
    
    # Close remaining valves
    aut.set({
        AIR_VALVE_1_CMD: False,
        AIR_VALVE_2_CMD: False,
        FUEL_RES_VALVE_1_CMD: False,
        FUEL_RES_VALVE_2_CMD: False,
    })
    
    log(aut, "Engine shutdown complete")
    return True

def execute_test(params: EngineParameters = EngineParameters()) -> sy.Range:
    """Execute complete engine test sequence"""
    with client.control.acquire(
        "Engine Test",
        write=[
            FUEL_PUMP_CMD, FUEL_VALVE_1_CMD, FUEL_VALVE_2_CMD,
            FUEL_RES_VALVE_1_CMD, FUEL_RES_VALVE_2_CMD,
            AIR_VALVE_1_CMD, AIR_VALVE_2_CMD, BLEED_VALVE_CMD,
            SPARK_PLUG_CMD, STARTER_MOTOR_CMD, IGNITION_CMD,
            auto_logs.name,
        ],
        read=[N1_SPEED, N2_SPEED, FLAME, COMBUSTION_TC_1, EXHAUST_TC, start_auto_cmd.name],
        write_authorities=[250],
    ) as ctrl:
        try:
            # Create range for test
            rng = client.ranges.create(
                name="Engine Test",
                time_range=sy.TimeRange(sy.TimeStamp.now(), sy.TimeStamp.now()),
            )
            
            # Execute startup
            start = sy.TimeStamp.now()
            if not execute_startup(ctrl, params):
                log(ctrl, "Startup failed - aborting test")
                execute_shutdown(ctrl, params)
                return rng
            
            # Create startup subrange
            rng.create_sub_range(
                name="Startup",
                time_range=sy.TimeRange(start, sy.TimeStamp.now()),
                color="#bada55",
            )
            
            # Run for 30 seconds at idle
            log(ctrl, "Running at idle for 30 seconds")
            idle_start = sy.TimeStamp.now()
            ctrl.sleep(30)
            idle_end = sy.TimeStamp.now()
            
            # Create idle subrange
            rng.create_sub_range(
                name="Idle",
                time_range=sy.TimeRange(idle_start, idle_end),
                color="#00ff00",
            )
            
            # Shutdown
            shutdown_start = sy.TimeStamp.now()
            if not execute_shutdown(ctrl, params):
                log(ctrl, "Shutdown failed")
            
            # Create shutdown subrange
            rng.create_sub_range(
                name="Shutdown",
                time_range=sy.TimeRange(shutdown_start, sy.TimeStamp.now()),
                color="#ff0000",
            )
            
            return rng
            
        except Exception as e:
            log(ctrl, f"Test failed with error: {e}")
            execute_shutdown(ctrl, params)
            raise e

if __name__ == "__main__":
    execute_test()
