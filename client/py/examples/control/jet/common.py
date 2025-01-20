
# Time channel
DAQ_TIME = "daq_time"

# Command and state channels
FUEL_PUMP_CMD = "fuel_pump_cmd"
FUEL_PUMP_STATE = "fuel_pump_state"
FUEL_VALVE_1_CMD = "fuel_valve_1_cmd"
FUEL_VALVE_1_STATE = "fuel_valve_1_state"
FUEL_VALVE_2_CMD = "fuel_valve_2_cmd"
FUEL_VALVE_2_STATE = "fuel_valve_2_state"
FUEL_RES_VALVE_1_CMD = "fuel_res_valve_1_cmd"
FUEL_RES_VALVE_1_STATE = "fuel_res_valve_1_state"
FUEL_RES_VALVE_2_CMD = "fuel_res_valve_2_cmd"
FUEL_RES_VALVE_2_STATE = "fuel_res_valve_2_state"
AIR_VALVE_1_CMD = "air_valve_1_cmd"
AIR_VALVE_1_STATE = "air_valve_1_state"
AIR_VALVE_2_CMD = "air_valve_2_cmd"
AIR_VALVE_2_STATE = "air_valve_2_state"
BLEED_VALVE_CMD = "bleed_valve_cmd"
BLEED_VALVE_STATE = "bleed_valve_state"
SPARK_PLUG_CMD = "spark_plug_cmd"
SPARK_PLUG_STATE = "spark_plug_state"
STARTER_MOTOR_CMD = "starter_motor_cmd"
STARTER_MOTOR_STATE = "starter_motor_state"
IGNITION_CMD = "ignition_cmd"
IGNITION_STATE = "ignition_state"

# Sensor channels
FUEL_RES_PT = "fuel_res_pt"
AIR_SUPPLY_PT = "air_supply_pt"
OIL_PT = "oil_pt"
AMBIENT_PT = "ambient_pt"
AMBIENT_TC = "ambient_tc"
FAN_INLET_TC = "fan_inlet_tc"
COMPRESSOR_INLET_TC = "compressor_inlet_tc"
COMPRESSOR_INLET_PT = "compressor_inlet_pt"
FLAME = "flame"
COMBUSTION_PT = "combustion_pt"
COMBUSTION_TC_1 = "combustion_tc_1"
COMBUSTION_TC_2 = "combustion_tc_2"
EXHAUST_TC = "exhaust_tc"
EXHAUST_FLOW = "exhaust_flow"
FUEL_TC = "fuel_tc"
FUEL_PT = "fuel_pt"
FUEL_FLOW = "fuel_flow"
OIL_TC = "oil_tc"
OIL_PT = "oil_pt"
OIL_FLOW = "oil_flow"
STARTER_MOTOR_CURRENT = "starter_motor_current"
VIBRATION = "vibration"
THRUST = "thrust"
N1_SPEED = "n1_speed"
N2_SPEED = "n2_speed"

# Group channels by type
VALVES = {
    FUEL_PUMP_CMD: FUEL_PUMP_STATE,
    FUEL_VALVE_1_CMD: FUEL_VALVE_1_STATE,
    FUEL_VALVE_2_CMD: FUEL_VALVE_2_STATE,
    FUEL_RES_VALVE_1_CMD: FUEL_RES_VALVE_1_STATE,
    FUEL_RES_VALVE_2_CMD: FUEL_RES_VALVE_2_STATE,
    AIR_VALVE_1_CMD: AIR_VALVE_1_STATE,
    AIR_VALVE_2_CMD: AIR_VALVE_2_STATE,
    BLEED_VALVE_CMD: BLEED_VALVE_STATE,
    SPARK_PLUG_CMD: SPARK_PLUG_STATE,
    STARTER_MOTOR_CMD: STARTER_MOTOR_STATE,
    IGNITION_CMD: IGNITION_STATE,
}

SENSORS = [
    FUEL_RES_PT,
    AIR_SUPPLY_PT,
    OIL_PT,
    AMBIENT_PT,
    AMBIENT_TC,
    FAN_INLET_TC,
    COMPRESSOR_INLET_TC,
    COMPRESSOR_INLET_PT,
    FLAME,
    COMBUSTION_PT,
    COMBUSTION_TC_1,
    COMBUSTION_TC_2,
    EXHAUST_TC,
    FUEL_TC,
    EXHAUST_FLOW,
    FUEL_PT,
    FUEL_FLOW,
    OIL_TC,
    OIL_PT,
    OIL_FLOW,
    STARTER_MOTOR_CURRENT,
    VIBRATION,
    THRUST,
    N1_SPEED,
    N2_SPEED,
]