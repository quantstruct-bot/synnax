// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#include "lib_manager.h"
#include <iostream>

#ifdef _WIN32
ni::LibraryManager::LibraryManager() : ni_daqmx_library_handle(nullptr), ni_syscfg_library_handle(nullptr),
    DAQmxCreateDIChan(nullptr),
    DAQmxCreateDOChan(nullptr),
    DAQmxCfgSampClkTiming(nullptr),
    DAQmxStartTask(nullptr),
    DAQmxStopTask(nullptr),
    DAQmxClearTask(nullptr),
    DAQmxReadAnalogF64(nullptr),
    DAQmxReadDigitalLines(nullptr),
    DAQmxWriteDigitalLines(nullptr),
    DAQmxGetExtendedErrorInfo(nullptr),
    DAQmxCreateLinScale(nullptr),
    DAQmxCreateMapScale(nullptr),
    DAQmxCreatePolynomialScale(nullptr),
    DAQmxCreateTableScale(nullptr),
    DAQmxCalculateReversePolyCoeff(nullptr),
    DAQmxCreateTask(nullptr),
    DAQmxCreateAIVoltageChan(nullptr),
    DAQmxCreateAIVoltageRMSChan(nullptr),
    DAQmxCreateAIVoltageChanWithExcit(nullptr),
    DAQmxCreateAIAccel4WireDCVoltageChan(nullptr),
    DAQmxCreateAIAccelChan(nullptr),
    DAQmxCreateAIAccelChargeChan(nullptr),
    DAQmxCreateAIBridgeChan(nullptr),
    DAQmxCreateAIChargeChan(nullptr),
    DAQmxCreateAICurrentChan(nullptr),
    DAQmxCreateAICurrentRMSChan(nullptr),
    DAQmxCreateAIForceBridgePolynomialChan(nullptr),
    DAQmxCreateAIForceBridgeTableChan(nullptr),
    DAQmxCreateAIForceBridgeTwoPointLinChan(nullptr),
    DAQmxCreateAIForceIEPEChan(nullptr),
    DAQmxCreateAIFreqVoltageChan(nullptr),
    DAQmxCreateAIMicrophoneChan(nullptr),
    DAQmxCreateAIPosEddyCurrProxProbeChan(nullptr),
    DAQmxCreateAIPosLVDTChan(nullptr),
    DAQmxCreateAIPosRVDTChan(nullptr),
    DAQmxCreateAIRTDChan(nullptr),
    DAQmxCreateAIResistanceChan(nullptr),
    DAQmxCreateAIRosetteStrainGageChan(nullptr),
    DAQmxCreateAIStrainGageChan(nullptr),
    DAQmxCreateAITempBuiltInSensorChan(nullptr),
    DAQmxCreateAIThrmcplChan(nullptr),
    DAQmxCreateAIThrmstrChanIex(nullptr),
    DAQmxCreateAIThrmstrChanVex(nullptr),
    DAQmxCreateAITorqueBridgePolynomialChan(nullptr),
    DAQmxCreateAITorqueBridgeTableChan(nullptr),
    DAQmxCreateAITorqueBridgeTwoPointLinChan(nullptr),
    DAQmxCreateAIVelocityIEPEChan(nullptr),
    NISysCfgInitializeSession(nullptr),
    NISysCfgCreateFilter(nullptr),
    NISysCfgSetFilterProperty(nullptr),
    NISysCfgCloseHandle(nullptr),
    NISysCfgFindHardware(nullptr),
    NISysCfgNextResource(nullptr),
    NISysCfgGetResourceProperty(nullptr),
    NISysCfgGetResourceIndexedProperty(nullptr){}


ni::LibraryManager::~LibraryManager() {
    unloadLibrary();
}

ni::LibraryManager& ni::LibraryManager::getInstance() {
    static LibraryManager instance;
    return instance;
}

bool ni::LibraryManager::loadLibrary() {
    this->ni_daqmx_library_handle  = LoadLibrary("NIDAQmx.dll");
    this->ni_syscfg_library_handle = LoadLibrary("nisyscfg.dll");
    if (!ni_daqmx_library_handle) {
        std::cerr << "Error loading NIDAQmx library" << std::endl;
        return false;
    } 

    if(!ni_syscfg_library_handle) {
        std::cerr << "Error loading NISysCfg library" << std::endl;
        return false;
    }

    DAQmxCreateDIChan = (DAQmxCreateDIChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateDIChan");
    DAQmxCreateDOChan = (DAQmxCreateDOChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateDOChan");
    DAQmxCfgSampClkTiming = (DAQmxCfgSampClkTiming_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCfgSampClkTiming");
    DAQmxStartTask = (DAQmxStartTask_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxStartTask");
    DAQmxStopTask = (DAQmxStopTask_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxStopTask");
    DAQmxClearTask = (DAQmxClearTask_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxClearTask");
    DAQmxReadAnalogF64 = (DAQmxReadAnalogF64_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxReadAnalogF64");
    DAQmxReadDigitalLines = (DAQmxReadDigitalLines_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxReadDigitalLines");
    DAQmxWriteDigitalLines = (DAQmxWriteDigitalLines_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxWriteDigitalLines");
    DAQmxGetExtendedErrorInfo = (DAQmxGetExtendedErrorInfo_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxGetExtendedErrorInfo");
    DAQmxCreateLinScale = (DAQmxCreateLinScale_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateLinScale");
    DAQmxCreateMapScale = (DAQmxCreateMapScale_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateMapScale");
    DAQmxCreatePolynomialScale = (DAQmxCreatePolynomialScale_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreatePolynomialScale");
    DAQmxCreateTableScale = (DAQmxCreateTableScale_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateTableScale");
    DAQmxCalculateReversePolyCoeff = (DAQmxCalculateReversePolyCoeff_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCalculateReversePolyCoeff");
    DAQmxCreateTask = (DAQmxCreateTask_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateTask");
    DAQmxCreateAIVoltageChan = (DAQmxCreateAIVoltageChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIVoltageChan");
    DAQmxCreateAIVoltageRMSChan = (DAQmxCreateAIVoltageRMSChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIVoltageRMSChan");
    DAQmxCreateAIVoltageChanWithExcit = (DAQmxCreateAIVoltageChanWithExcit_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIVoltageChanWithExcit");
    DAQmxCreateAIAccel4WireDCVoltageChan = (DAQmxCreateAIAccel4WireDCVoltageChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIAccel4WireDCVoltageChan");
    DAQmxCreateAIAccelChan = (DAQmxCreateAIAccelChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIAccelChan");
    DAQmxCreateAIAccelChargeChan = (DAQmxCreateAIAccelChargeChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIAccelChargeChan");
    DAQmxCreateAIBridgeChan = (DAQmxCreateAIBridgeChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIBridgeChan");
    DAQmxCreateAIChargeChan = (DAQmxCreateAIChargeChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIChargeChan");
    DAQmxCreateAICurrentChan = (DAQmxCreateAICurrentChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAICurrentChan");
    DAQmxCreateAICurrentRMSChan = (DAQmxCreateAICurrentRMSChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAICurrentRMSChan");
    DAQmxCreateAIForceBridgePolynomialChan = (DAQmxCreateAIForceBridgePolynomialChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIForceBridgePolynomialChan");
    DAQmxCreateAIForceBridgeTableChan = (DAQmxCreateAIForceBridgeTableChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIForceBridgeTableChan");
    DAQmxCreateAIForceBridgeTwoPointLinChan = (DAQmxCreateAIForceBridgeTwoPointLinChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIForceBridgeTwoPointLinChan");
    DAQmxCreateAIForceIEPEChan = (DAQmxCreateAIForceIEPEChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIForceIEPEChan");
    DAQmxCreateAIFreqVoltageChan = (DAQmxCreateAIFreqVoltageChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIFreqVoltageChan");
    DAQmxCreateAIMicrophoneChan = (DAQmxCreateAIMicrophoneChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIMicrophoneChan");
    DAQmxCreateAIPosEddyCurrProxProbeChan = (DAQmxCreateAIPosEddyCurrProxProbeChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIPosEddyCurrProxProbeChan");
    DAQmxCreateAIPosLVDTChan = (DAQmxCreateAIPosLVDTChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIPosLVDTChan");
    DAQmxCreateAIPosRVDTChan = (DAQmxCreateAIPosRVDTChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIPosRVDTChan");
    DAQmxCreateAIRTDChan = (DAQmxCreateAIRTDChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIRTDChan");
    DAQmxCreateAIResistanceChan = (DAQmxCreateAIResistanceChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIResistanceChan");
    DAQmxCreateAIRosetteStrainGageChan = (DAQmxCreateAIRosetteStrainGageChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIRosetteStrainGageChan");
    DAQmxCreateAIStrainGageChan = (DAQmxCreateAIStrainGageChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIStrainGageChan");
    DAQmxCreateAITempBuiltInSensorChan = (DAQmxCreateAITempBuiltInSensorChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAITempBuiltInSensorChan");
    DAQmxCreateAIThrmcplChan = (DAQmxCreateAIThrmcplChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIThrmcplChan");
    DAQmxCreateAIThrmstrChanIex = (DAQmxCreateAIThrmstrChanIex_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIThrmstrChanIex");
    DAQmxCreateAIThrmstrChanVex = (DAQmxCreateAIThrmstrChanVex_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIThrmstrChanVex");
    DAQmxCreateAITorqueBridgePolynomialChan = (DAQmxCreateAITorqueBridgePolynomialChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAITorqueBridgePolynomialChan");
    DAQmxCreateAITorqueBridgeTableChan = (DAQmxCreateAITorqueBridgeTableChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAITorqueBridgeTableChan");
    DAQmxCreateAITorqueBridgeTwoPointLinChan = (DAQmxCreateAITorqueBridgeTwoPointLinChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAITorqueBridgeTwoPointLinChan");
    DAQmxCreateAIVelocityIEPEChan = (DAQmxCreateAIVelocityIEPEChan_t)GetProcAddress(ni_daqmx_library_handle, "DAQmxCreateAIVelocityIEPEChan");

    // NISysCfg
    NISysCfgInitializeSession = (NISysCfgInitializeSession_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgInitializeSession");
    NISysCfgCreateFilter = (NISysCfgCreateFilter_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgCreateFilter");
    NISysCfgSetFilterProperty = (NISysCfgSetFilterProperty_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgSetFilterProperty");
    NISysCfgCloseHandle = (NISysCfgCloseHandle_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgCloseHandle");
    NISysCfgFindHardware = (NISysCfgFindHardware_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgFindHardware");
    NISysCfgNextResource = (NISysCfgNextResource_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgNextResource");
    NISysCfgGetResourceProperty = (NISysCfgGetResourceProperty_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgGetResourceProperty");
    NISysCfgGetResourceIndexedProperty = (NISysCfgGetResourceIndexedProperty_t)GetProcAddress(ni_syscfg_library_handle, "NISysCfgGetResourceIndexedProperty");


    if (!DAQmxCreateDIChan || !DAQmxCreateDOChan || !DAQmxCfgSampClkTiming || !DAQmxStartTask || !DAQmxStopTask || !DAQmxClearTask || !DAQmxReadAnalogF64 ||
        !DAQmxReadDigitalLines || !DAQmxWriteDigitalLines || !DAQmxGetExtendedErrorInfo || !DAQmxCreateLinScale || !DAQmxCreateMapScale || !DAQmxCreatePolynomialScale ||
        !DAQmxCreateTableScale || !DAQmxCalculateReversePolyCoeff || !DAQmxCreateTask || !DAQmxCreateAIVoltageChan || !DAQmxCreateAIVoltageRMSChan || !DAQmxCreateAIVoltageChanWithExcit ||
        !DAQmxCreateAIAccel4WireDCVoltageChan || !DAQmxCreateAIAccelChan || !DAQmxCreateAIAccelChargeChan || !DAQmxCreateAIBridgeChan || !DAQmxCreateAIChargeChan ||
        !DAQmxCreateAICurrentChan || !DAQmxCreateAICurrentRMSChan || !DAQmxCreateAIForceBridgePolynomialChan || !DAQmxCreateAIForceBridgeTableChan || !DAQmxCreateAIForceBridgeTwoPointLinChan ||
        !DAQmxCreateAIForceIEPEChan || !DAQmxCreateAIFreqVoltageChan || !DAQmxCreateAIMicrophoneChan || !DAQmxCreateAIPosEddyCurrProxProbeChan || !DAQmxCreateAIPosLVDTChan ||
        !DAQmxCreateAIPosRVDTChan || !DAQmxCreateAIRTDChan || !DAQmxCreateAIResistanceChan || !DAQmxCreateAIRosetteStrainGageChan || !DAQmxCreateAIStrainGageChan ||
        !DAQmxCreateAITempBuiltInSensorChan || !DAQmxCreateAIThrmcplChan || !DAQmxCreateAIThrmstrChanIex || !DAQmxCreateAIThrmstrChanVex || !DAQmxCreateAITorqueBridgePolynomialChan ||
        !DAQmxCreateAITorqueBridgeTableChan || !DAQmxCreateAITorqueBridgeTwoPointLinChan || !DAQmxCreateAIVelocityIEPEChan) {
        std::cerr << "Error getting function pointers from NIDAQmx library" << std::endl;
        FreeLibrary(ni_daqmx_library_handle);
        FreeLibrary(ni_syscfg_library_handle);
        ni_daqmx_library_handle = nullptr;
        ni_syscfg_library_handle = nullptr;
        return false;
    }


    if(!NISysCfgInitializeSession || !NISysCfgCreateFilter || !NISysCfgSetFilterProperty || !NISysCfgCloseHandle || !NISysCfgFindHardware ||
        !NISysCfgNextResource || !NISysCfgGetResourceProperty || !NISysCfgGetResourceIndexedProperty) {
        std::cerr << "Error getting function pointers from NISysCfg library" << std::endl;
        FreeLibrary(ni_daqmx_library_handle);
        FreeLibrary(ni_syscfg_library_handle);
        ni_daqmx_library_handle = nullptr;
        ni_syscfg_library_handle = nullptr;
        return false;
    }

    return true;
}

void ni::LibraryManager::unloadLibrary() {
    if(ni_daqmx_library_handle) {
        FreeLibrary(ni_daqmx_library_handle);
        ni_daqmx_library_handle = nullptr;
    }
    if(ni_syscfg_library_handle) {
        FreeLibrary(ni_syscfg_library_handle);
        ni_syscfg_library_handle = nullptr;
    }
}

#else // Non-Windows implementation
ni::LibraryManager::LibraryManager() :
    DAQmxCreateDIChan(nullptr),
    DAQmxCreateDOChan(nullptr),
    DAQmxCfgSampClkTiming(nullptr),
    DAQmxStartTask(nullptr),
    DAQmxStopTask(nullptr),
    DAQmxClearTask(nullptr),
    DAQmxReadAnalogF64(nullptr),
    DAQmxReadDigitalLines(nullptr),
    DAQmxWriteDigitalLines(nullptr),
    DAQmxGetExtendedErrorInfo(nullptr),
    DAQmxCreateLinScale(nullptr),
    DAQmxCreateMapScale(nullptr),
    DAQmxCreatePolynomialScale(nullptr),
    DAQmxCreateTableScale(nullptr),
    DAQmxCalculateReversePolyCoeff(nullptr),
    DAQmxCreateTask(nullptr),
    DAQmxCreateAIVoltageChan(nullptr),
    DAQmxCreateAIVoltageRMSChan(nullptr),
    DAQmxCreateAIVoltageChanWithExcit(nullptr),
    DAQmxCreateAIAccel4WireDCVoltageChan(nullptr),
    DAQmxCreateAIAccelChan(nullptr),
    DAQmxCreateAIAccelChargeChan(nullptr),
    DAQmxCreateAIBridgeChan(nullptr),
    DAQmxCreateAIChargeChan(nullptr),
    DAQmxCreateAICurrentChan(nullptr),
    DAQmxCreateAICurrentRMSChan(nullptr),
    DAQmxCreateAIForceBridgePolynomialChan(nullptr),
    DAQmxCreateAIForceBridgeTableChan(nullptr),
    DAQmxCreateAIForceBridgeTwoPointLinChan(nullptr),
    DAQmxCreateAIForceIEPEChan(nullptr),
    DAQmxCreateAIFreqVoltageChan(nullptr),
    DAQmxCreateAIMicrophoneChan(nullptr),
    DAQmxCreateAIPosEddyCurrProxProbeChan(nullptr),
    DAQmxCreateAIPosLVDTChan(nullptr),
    DAQmxCreateAIPosRVDTChan(nullptr),
    DAQmxCreateAIRTDChan(nullptr),
    DAQmxCreateAIResistanceChan(nullptr),
    DAQmxCreateAIRosetteStrainGageChan(nullptr),
    DAQmxCreateAIStrainGageChan(nullptr),
    DAQmxCreateAITempBuiltInSensorChan(nullptr),
    DAQmxCreateAIThrmcplChan(nullptr),
    DAQmxCreateAIThrmstrChanIex(nullptr),
    DAQmxCreateAIThrmstrChanVex(nullptr),
    DAQmxCreateAITorqueBridgePolynomialChan(nullptr),
    DAQmxCreateAITorqueBridgeTableChan(nullptr),
    DAQmxCreateAITorqueBridgeTwoPointLinChan(nullptr),
    DAQmxCreateAIVelocityIEPEChan(nullptr),
    NISysCfgInitializeSession(nullptr),
    NISysCfgCreateFilter(nullptr),
    NISysCfgSetFilterProperty(nullptr),
    NISysCfgCloseHandle(nullptr),
    NISysCfgFindHardware(nullptr),
    NISysCfgNextResource(nullptr),
    NISysCfgGetResourceProperty(nullptr),
    NISysCfgGetResourceIndexedProperty(nullptr){}


ni::LibraryManager::~LibraryManager() {}

ni::LibraryManager& ni::LibraryManager::getInstance() {
    static LibraryManager instance;
    return instance;
}

bool ni::LibraryManager::loadLibrary() {
    // On non-Windows systems, return success and set stubs
    DAQmxCreateDIChan = [](TaskHandle, const char[], const char[], int32_t) { return 0; };
    DAQmxCreateDOChan = [](TaskHandle, const char[], const char[], int32_t) { return 0; };
    DAQmxCfgSampClkTiming = [](TaskHandle, const char[], float64, int32_t, int32_t, uInt64) { return 0; };
    DAQmxStartTask = [](TaskHandle) { return 0; };
    DAQmxStopTask = [](TaskHandle) { return 0; };
    DAQmxClearTask = [](TaskHandle) { return 0; };
    DAQmxReadAnalogF64 = [](TaskHandle, int32_t, float64, int32_t, float64[], uInt32, int32_t*, bool32*) { return 0; };
    DAQmxReadDigitalLines = [](TaskHandle, int32_t, float64, int32_t, uInt8[], uInt32, int32_t*, int32_t*, bool32*) { return 0; };
    DAQmxWriteDigitalLines = [](TaskHandle, int32_t, bool32, float64, int32_t, const uInt8[], int32_t*, bool32*) { return 0; };
    DAQmxGetExtendedErrorInfo = [](char[], uInt32) { return 0; };
    DAQmxCreateLinScale = [](const char[], float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateMapScale = [](const char[], float64, float64, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreatePolynomialScale = [](const char[], const float64[], uInt32, const float64[], uInt32, int32_t, const char[]) { return 0; };
    DAQmxCreateTableScale = [](const char[], const float64[], uInt32, const float64[], uInt32, int32_t, const char[]) { return 0; };
    DAQmxCalculateReversePolyCoeff = [](const float64[], uInt32, float64, float64, int32_t, int32_t, float64[]) { return 0; };
    DAQmxCreateTask = [](const char[], TaskHandle*) { return 0; };
    DAQmxCreateAIVoltageChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIVoltageRMSChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIVoltageChanWithExcit = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, int32_t, int32_t, float64, bool32, const char[]) { return 0; };
    DAQmxCreateAIAccel4WireDCVoltageChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, float64, int32_t, int32_t, float64, bool32, const char[]) { return 0; };
    DAQmxCreateAIAccelChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, float64, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIAccelChargeChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIBridgeChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIChargeChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAICurrentChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAICurrentRMSChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIForceBridgePolynomialChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, float64, float64, const float64[], uInt32, const float64[], uInt32, int32_t, int32_t, const char[]) { return 0; };
    DAQmxCreateAIForceBridgeTableChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, float64, float64, const float64[], uInt32, int32_t, const float64[], uInt32, int32_t, const char[]) { return 0; };
    DAQmxCreateAIForceBridgeTwoPointLinChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, float64, float64, float64, float64, int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIForceIEPEChan = [](TaskHandle, const char[], const char[], int32_t, float64, float64, int32_t, float64, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIFreqVoltageChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, float64, float64, const char[]) { return 0; };
    DAQmxCreateAIMicrophoneChan = [](TaskHandle, const char[], const char[], int32_t, int32_t, float64, float64, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIPosEddyCurrProxProbeChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIPosLVDTChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, float64, int32_t, int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIPosRVDTChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, float64, int32_t, int32_t, float64, float64, int32_t, const char[]) { return 0; };
    DAQmxCreateAIRTDChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, int32_t, float64, float64) { return 0; };
    DAQmxCreateAIResistanceChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, int32_t, int32_t, float64, const char[]) { return 0; };
    DAQmxCreateAIRosetteStrainGageChan = [](TaskHandle, const char[], const char[], float64, float64, int32_t, float64, const int32[], uInt32, int32, int32, float64, float64, float64, float64) { return 0; };
    DAQmxCreateAIStrainGageChan = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, float64, const char[]) { return 0; };
    DAQmxCreateAITempBuiltInSensorChan = [](TaskHandle, const char[], const char[], int32) { return 0; };
    DAQmxCreateAIThrmcplChan = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, const char[]) { return 0; };
    DAQmxCreateAIThrmstrChanIex = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64) { return 0; };
    DAQmxCreateAIThrmstrChanVex = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, float64, float64) { return 0; };
    DAQmxCreateAITorqueBridgePolynomialChan = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, const float64[], uInt32, int32, int32, const char[]) { return 0; };
    DAQmxCreateAITorqueBridgeTableChan = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, int32, const float64[], uInt32, int32, const char[]) { return 0; };
    DAQmxCreateAITorqueBridgeTwoPointLinChan = [](TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, int32, float64, float64, int32, const char[]) { return 0; };
    DAQmxCreateAIVelocityIEPEChan = [](TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, int32, float64, const char[]) { return 0; };

    NISysCfgInitializeSession = [](const char*, const char*, const char*, NISysCfgLocale, NISysCfgBool, unsigned int, NISysCfgEnumExpertHandle*, NISysCfgSessionHandle*) { return 0; };
    NISysCfgCreateFilter = [](NISysCfgSessionHandle, NISysCfgFilterHandle*) { return 0; };
    NISysCfgSetFilterProperty = [](NISysCfgFilterHandle, NISysCfgFilterProperty, ...) { return 0; };
    NISysCfgCloseHandle = [](void*) { return 0; };
    NISysCfgFindHardware = [](NISysCfgSessionHandle, NISysCfgFilterMode, NISysCfgFilterHandle, const char*, NISysCfgEnumResourceHandle*) { return 0; };
    NISysCfgNextResource = [](NISysCfgSessionHandle, NISysCfgEnumResourceHandle, NISysCfgResourceHandle*) { return 0; };
    NISysCfgGetResourceProperty = [](NISysCfgResourceHandle, NISysCfgResourceProperty, void*) { return 0; };
    NISysCfgGetResourceIndexedProperty = [](NISysCfgResourceHandle, NISysCfgIndexedProperty, unsigned int, void*) { return 0; };
    
    return true;
}

void ni::LibraryManager::unloadLibrary() {
    // Nothing to unload on non-Windows systems
}

#endif // _WIN32

