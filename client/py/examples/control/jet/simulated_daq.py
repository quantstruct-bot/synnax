#  Copyright 2024 Synnax Labs, Inc.
#
#  Use of this software is governed by the Business Source License included in the file
#  licenses/BSL.txt.
#
#  As of the Change Date specified in that file, in accordance with the Business Source
#  License, use of this software will be governed by the Apache License, Version 2.0,
#  included in the file licenses/APL.txt.

import json
import os
from common import (
    DAQ_TIME,
    VALVES,
    SENSORS,
    FLAME,
    FUEL_PUMP_STATE,
    FUEL_VALVE_1_STATE,
    FUEL_VALVE_2_STATE,
    FUEL_RES_VALVE_1_STATE,
    FUEL_RES_VALVE_2_STATE,
    AIR_VALVE_1_STATE,
    AIR_VALVE_2_STATE,
    BLEED_VALVE_STATE,
    SPARK_PLUG_STATE,
    STARTER_MOTOR_STATE,
    IGNITION_STATE,
    FUEL_RES_PT,
    AIR_SUPPLY_PT,
    FAN_SPEED,
    OIL_PT,
    COMPRESSOR_SPEED,
    AMBIENT_PT,
    AMBIENT_TC,
    FAN_INLET_TC,
    COMPRESSOR_INLET_TC,
    COMPRESSOR_INLET_PT,
    COMPRESSOR_SHAFT_SPEED,
    COMBUSTION_PT,
    COMBUSTION_TC,
    EXHAUST_TC,
    FUEL_TC,
    FUEL_PT,
    FUEL_FLOW,
    OIL_TC,
    OIL_FLOW,
    STARTER_MOTOR_CURRENT,
    VIBRATION,
    THRUST,
    N1_SPEED,
    N2_SPEED,
)

import numpy as np
import synnax as sy


# Initialize client
client = sy.Synnax()

# Create DAQ time channel
daq_time = client.channels.create(
    name=DAQ_TIME, 
    is_index=True, 
    retrieve_if_name_exists=True
)

# Create valve command and state channels
for cmd, state in VALVES.items():
    client.channels.create(
        [
            sy.Channel(
                name=cmd, 
                data_type=sy.DataType.UINT8, 
                virtual=True
            ),
            sy.Channel(
                name=state, 
                data_type=sy.DataType.UINT8, 
                index=daq_time.key
            ),
        ],
        retrieve_if_name_exists=True,
    )

# Create sensor channels
for sensor in SENSORS:
    data_type = sy.DataType.UINT8 if sensor == FLAME else sy.DataType.FLOAT32
    s = client.channels.create(
        name=sensor,
        data_type=data_type,
        index=daq_time.key,
        retrieve_if_name_exists=True,
    )
    print(f"Created channel: {s.name} with key: {s.key}")

# Initial state dictionary with all sensor values
DAQ_STATE = {
    # Valve states
    FUEL_PUMP_STATE: 0,
    FUEL_VALVE_1_STATE: 0,
    FUEL_VALVE_2_STATE: 0,
    FUEL_RES_VALVE_1_STATE: 0,
    FUEL_RES_VALVE_2_STATE: 0,
    AIR_VALVE_1_STATE: 0,
    AIR_VALVE_2_STATE: 0,
    BLEED_VALVE_STATE: 0,
    SPARK_PLUG_STATE: 0,
    STARTER_MOTOR_STATE: 0,
    IGNITION_STATE: 0,
    
    # Sensor initial values
    FUEL_RES_PT: 0,
    AIR_SUPPLY_PT: 0,
    FAN_SPEED: 0,
    OIL_PT: 0,
    COMPRESSOR_SPEED: 0,
    AMBIENT_PT: 14.7,  # Standard atmospheric pressure
    AMBIENT_TC: 25,    # Room temperature
    FAN_INLET_TC: 25,
    COMPRESSOR_INLET_TC: 25,
    COMPRESSOR_INLET_PT: 14.7,
    COMPRESSOR_SHAFT_SPEED: 0,
    FLAME: 0,
    COMBUSTION_PT: 14.7,
    COMBUSTION_TC: 25,
    EXHAUST_TC: 25,
    FUEL_TC: 25,
    FUEL_PT: 0,
    FUEL_FLOW: 0,
    OIL_TC: 25,
    OIL_PT: 0,
    OIL_FLOW: 0,
    STARTER_MOTOR_CURRENT: 0,
    VIBRATION: 0,
    THRUST: 0,
    N1_SPEED: 0,
    N2_SPEED: 0,
}

# Simulation constants
IDLE_N1_SPEED = 20000.0   # RPM - Increased from 2000
MAX_N1_SPEED = 25000.0    # RPM - Increased from 5000
IDLE_N2_SPEED = 45000.0   # RPM - Increased from 10000
MAX_N2_SPEED = 55000.0    # RPM - Increased from 15000
MAX_THRUST = 1000.0      # lbf
AMBIENT_TEMP = 25.0     # Celsius
AMBIENT_PRESS = 14.7    # PSI

def update_sensors_starting(state):
    """Update sensor values during engine start"""
    if state[STARTER_MOTOR_STATE]:
        # N1 (fan) spool-up to ~15% of max for light-off
        target_n1 = MAX_N1_SPEED * 0.15
        current_n1 = state[N1_SPEED]
        if current_n1 < target_n1:
            state[N1_SPEED] = min(current_n1 + 800.0, target_n1)
            
        # N2 (core) spools up faster and to a higher speed
        target_n2 = MAX_N2_SPEED * 0.15
        current_n2 = state[N2_SPEED]
        if current_n2 < target_n2:
            state[N2_SPEED] = min(current_n2 + 2000.0, target_n2)
            
        # Check for ignition conditions
        if not state[FLAME]:  # Only check if not already lit
            conditions = {
                "Spark Plug": (bool(state[SPARK_PLUG_STATE]), f"Spark Plug: {state[SPARK_PLUG_STATE]}"),
                "Ignition": (bool(state[IGNITION_STATE]), f"Ignition: {state[IGNITION_STATE]}"),
                "Fuel Valve": (bool(state[FUEL_VALVE_1_STATE]), f"Fuel Valve: {state[FUEL_VALVE_1_STATE]}"),
                "Fuel Pump": (bool(state[FUEL_PUMP_STATE]), f"Fuel Pump: {state[FUEL_PUMP_STATE]}"),
                "N2 Speed": (state[N2_SPEED] >= MAX_N2_SPEED * 0.15, 
                            f"{state[N2_SPEED]:.0f} RPM >= {MAX_N2_SPEED * 0.15:.0f} RPM")
            }
            
            # Log any unmet conditions
            unmet = [details[1] for name, details in conditions.items() 
                    if not details[0]]
            
            if unmet:
                print(f"Ignition conditions not met: {', '.join(unmet)}")
            else:
                state[FLAME] = 1.0
                print("All conditions met - flame lit!")
            
        # Update dependent parameters
        state[COMPRESSOR_SPEED] = state[N2_SPEED]
        state[FAN_SPEED] = state[N1_SPEED]
        
        # Starter motor current increases with speed
        state[STARTER_MOTOR_CURRENT] = 10.0 * (state[N1_SPEED] / target_n1)
    
    return state

def update_sensors_running(state):
    """Update sensor values during normal engine operation"""
    # Check if fuel supply is cut off - if so, extinguish flame
    if (state[FLAME] and 
        (not state[FUEL_PUMP_STATE] or 
         not state[FUEL_VALVE_1_STATE] or 
         not state[FUEL_VALVE_2_STATE])):
        state[FLAME] = 0.0
        print("Flame extinguished - fuel supply cut off")
    
    if state[FLAME] > 0:  # Only accelerate if flame is lit
        # N1 acceleration
        target_n1 = IDLE_N1_SPEED
        current_n1 = state[N1_SPEED]
        if current_n1 < target_n1:
            state[N1_SPEED] = min(current_n1 + 1500.0, target_n1)
            
        # N2 accelerates faster than N1
        target_n2 = IDLE_N2_SPEED
        current_n2 = state[N2_SPEED]
        if current_n2 < target_n2:
            state[N2_SPEED] = min(current_n2 + 3000.0, target_n2)
        # Update temperatures based on N2 speed
        n2_ratio = current_n2 / MAX_N2_SPEED
        state[COMBUSTION_TC] = state[AMBIENT_TC] + (1200 * n2_ratio)
        state[EXHAUST_TC] = state[COMBUSTION_TC] * 0.9
        
        # Update pressures based on N2
        state[COMBUSTION_PT] = state[AMBIENT_PT] * (1 + (4.0 * n2_ratio))
        
        # Update thrust based on both N1 and N2
        state[THRUST] = MAX_THRUST * (n2_ratio * 0.7 + (current_n1 / MAX_N1_SPEED) * 0.3)
    else:
        # When flame is out, temperatures and pressures should decay much faster
        state[COMBUSTION_TC] = max(state[AMBIENT_TC], state[COMBUSTION_TC] - 25.0)
        state[EXHAUST_TC] = max(state[AMBIENT_TC], state[EXHAUST_TC] - 20.0)
        state[COMBUSTION_PT] = max(state[AMBIENT_PT], state[COMBUSTION_PT] - 1.0)
        state[N1_SPEED] = max(0.0, state[N1_SPEED] - 500.0)
        state[N2_SPEED] = max(0.0, state[N2_SPEED] - 800.0)
        state[THRUST] = 0.0
    
    # Update dependent parameters
    state[COMPRESSOR_SPEED] = state[N2_SPEED]
    state[FAN_SPEED] = state[N1_SPEED]
    
    return state

def update_sensors_shutdown(state):
    """Update sensor values during shutdown sequence"""
    # Gradually decrease speeds
    state[N1_SPEED] = max(0.0, state[N1_SPEED] - 500.0)
    state[N2_SPEED] = max(0.0, state[N2_SPEED] - 800.0)
    
    # Update dependent parameters
    state[COMPRESSOR_SPEED] = state[N2_SPEED]  # Compressor is on N2 shaft
    state[FAN_SPEED] = state[N1_SPEED]        # Fan is on N1 shaft
    
    # Update temperatures and pressures
    state[COMBUSTION_TC] = max(state[AMBIENT_TC], state[COMBUSTION_TC] - 25.0)
    state[EXHAUST_TC] = max(state[AMBIENT_TC], state[EXHAUST_TC] - 20.0)
    state[COMBUSTION_PT] = max(state[AMBIENT_PT], state[COMBUSTION_PT] - 1.0)
    state[THRUST] = 0.0
    
    return state

def introduce_randomness(state: dict) -> dict:
    """Add small random variations to sensor values"""
    for sensor in SENSORS:
        if sensor != FLAME:  # Don't add noise to boolean values
            current = state[sensor]
            # Add 1% random variation to non-zero values
            if current != 0:
                state[sensor] += current * np.random.normal(0, 0.01)
    return state

loop = sy.Loop(sy.Rate.HZ * 30, precise=True)

# open a CSV file
# delete the file if it exists
if os.path.exists("data.csv"):
    os.remove("data.csv")

with open("data.csv", "w") as f:
    with client.open_streamer([cmd for cmd in VALVES.keys()]) as streamer:
        with client.open_writer(
            sy.TimeStamp.now(),
            channels=[
                *[state for state in VALVES.values()],  # All valve state channels
                *SENSORS,  # All sensor channels
                DAQ_TIME,  # Time channel
            ],
            enable_auto_commit=True,
        ) as writer:
            while loop.wait():
                try:
                    # Read command changes
                    while True:
                        frame = streamer.read(0)
                        if frame is None:
                            break
                        for ch in frame.channels:
                            # Map command channel to state channel in DAQ_STATE
                            if ch in VALVES:
                                state_ch = VALVES[ch]
                                DAQ_STATE[state_ch] = frame[ch][0]
                    
                    # Update sensor values based on current state
                    if DAQ_STATE[STARTER_MOTOR_STATE]:
                        DAQ_STATE = update_sensors_starting(DAQ_STATE)
                    elif DAQ_STATE[FLAME]:
                        DAQ_STATE = update_sensors_running(DAQ_STATE)
                    else:
                        DAQ_STATE = update_sensors_shutdown(DAQ_STATE)
                    
                    # Add noise and random variations
                    DAQ_STATE = introduce_randomness(DAQ_STATE)
                    
                    # Create a new dict with only the channels we're writing
                    write_state = {
                        DAQ_TIME: sy.TimeStamp.now(),
                        **{state: DAQ_STATE[state] for state in VALVES.values()},
                        **{sensor: DAQ_STATE[sensor] for sensor in SENSORS}
                    }
                    
                    writer.write(write_state)
             
                except Exception as e:
                    print(e)
                    raise e
                


