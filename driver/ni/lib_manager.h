// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#pragma once

#include "nisyscfg.h"
#include "daqmx.h"

#ifdef _WIN32
#include <windows.h>
#else
#include <cstdint>
#endif

namespace ni {

// Ni-DAQmx function pointers 
typedef int32 (*DAQmxCreateDIChan_t)(TaskHandle, const char[], const char[], int32);
typedef int32 (*DAQmxCreateDOChan_t)(TaskHandle, const char[], const char[], int32);
typedef int32 (*DAQmxCfgSampClkTiming_t)(TaskHandle, const char[], float64, int32, int32, uInt64);
typedef int32 (*DAQmxStartTask_t)(TaskHandle);
typedef int32 (*DAQmxStopTask_t)(TaskHandle);
typedef int32 (*DAQmxClearTask_t)(TaskHandle);
typedef int32 (*DAQmxReadAnalogF64_t)(TaskHandle, int32, float64, int32, float64[], uInt32, int32*, bool32*);
typedef int32 (*DAQmxReadDigitalLines_t)(TaskHandle, int32, float64, int32, uInt8[], uInt32, int32*, int32*, bool32*);
typedef int32 (*DAQmxWriteDigitalLines_t)(TaskHandle, int32, bool32, float64, int32, const uInt8[], int32*, bool32*);
typedef int32 (*DAQmxGetExtendedErrorInfo_t)(char[], uInt32);
typedef int32 (*DAQmxCreateLinScale_t)(const char[], float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateMapScale_t)(const char[], float64, float64, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreatePolynomialScale_t)(const char[], const float64[], uInt32, const float64[], uInt32, int32, const char[]);
typedef int32 (*DAQmxCreateTableScale_t)(const char[], const float64[], uInt32, const float64[], uInt32, int32, const char[]);
typedef int32 (*DAQmxCalculateReversePolyCoeff_t)(const float64[], uInt32, float64, float64, int32, int32, float64[]);
typedef int32 (*DAQmxCreateTask_t)(const char[], TaskHandle*);
typedef int32 (*DAQmxCreateAIVoltageChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIVoltageRMSChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIVoltageChanWithExcit_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, int32, int32, float64, bool32, const char[]);
typedef int32 (*DAQmxCreateAIAccel4WireDCVoltageChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, int32, float64, bool32, const char[]);
typedef int32 (*DAQmxCreateAIAccelChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIAccelChargeChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIBridgeChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIChargeChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAICurrentChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAICurrentRMSChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIForceBridgePolynomialChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, const float64[], uInt32, int32, int32, const char[]);
typedef int32 (*DAQmxCreateAIForceBridgeTableChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, int32, const float64[], uInt32, int32, const char[]);
typedef int32 (*DAQmxCreateAIForceBridgeTwoPointLinChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIForceIEPEChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIFreqVoltageChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, float64, float64, const char[]);
typedef int32 (*DAQmxCreateAIMicrophoneChan_t)(TaskHandle, const char[], const char[], int32, int32, float64, float64, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIPosEddyCurrProxProbeChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIPosLVDTChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, float64, int32, int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIPosRVDTChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, float64, int32, int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIRTDChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, int32, float64, float64);
typedef int32 (*DAQmxCreateAIResistanceChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIRosetteStrainGageChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, float64, const int32[], uInt32, int32, int32, float64, float64, float64, float64);
typedef int32 (*DAQmxCreateAIStrainGageChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, float64, const char[]);
typedef int32 (*DAQmxCreateAITempBuiltInSensorChan_t)(TaskHandle, const char[], const char[], int32);
typedef int32 (*DAQmxCreateAIThrmcplChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, const char[]);
typedef int32 (*DAQmxCreateAIThrmstrChanIex_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64);
typedef int32 (*DAQmxCreateAIThrmstrChanVex_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, float64, float64);
typedef int32 (*DAQmxCreateAITorqueBridgePolynomialChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, const float64[], uInt32, int32, int32, const char[]);
typedef int32 (*DAQmxCreateAITorqueBridgeTableChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, const float64[], uInt32, int32, const float64[], uInt32, int32, const char[]);
typedef int32 (*DAQmxCreateAITorqueBridgeTwoPointLinChan_t)(TaskHandle, const char[], const char[], float64, float64, int32, int32, int32, float64, float64, float64, float64, int32, float64, float64, int32, const char[]);
typedef int32 (*DAQmxCreateAIVelocityIEPEChan_t)(TaskHandle, const char[], const char[], int32, float64, float64, int32, float64, int32, int32, float64, const char[]);

// NI-SysCfg function pointers
typedef NISysCfgStatus (*NISysCfgInitializeSession_t)(const char *targetName, const char *username, const char *password, NISysCfgLocale language, NISysCfgBool forcePropertyRefresh, unsigned int connectTimeoutMsec, NISysCfgEnumExpertHandle *expertEnumHandle, NISysCfgSessionHandle *sessionHandle);
typedef NISysCfgStatus (*NISysCfgCreateFilter_t)(NISysCfgSessionHandle sessionHandle, NISysCfgFilterHandle *filterHandle);
typedef NISysCfgStatus (*NISysCfgSetFilterProperty_t)(NISysCfgFilterHandle filterHandle, NISysCfgFilterProperty propertyID, ...);
typedef NISysCfgStatus (*NISysCfgCloseHandle_t)(void *syscfgHandle);
typedef NISysCfgStatus (*NISysCfgFindHardware_t)(NISysCfgSessionHandle sessionHandle, NISysCfgFilterMode filterMode, NISysCfgFilterHandle filterHandle, const char *expertNames, NISysCfgEnumResourceHandle *resourceEnumHandle);
typedef NISysCfgStatus (*NISysCfgNextResource_t)(NISysCfgSessionHandle sessionHandle, NISysCfgEnumResourceHandle resourceEnumHandle, NISysCfgResourceHandle *resourceHandle);
typedef NISysCfgStatus (*NISysCfgGetResourceProperty_t)(NISysCfgResourceHandle resourceHandle, NISysCfgResourceProperty propertyID, void *value);
typedef NISysCfgStatus (*NISysCfgGetResourceIndexedProperty_t)(NISysCfgResourceHandle resourceHandle, NISysCfgIndexedProperty propertyID, unsigned int index, void *value);


class LibraryManager {
public:
    static LibraryManager& getInstance();
    bool loadLibrary();
    void unloadLibrary();
    bool doLibrariesExist();

    // Function pointers
    DAQmxCreateDIChan_t DAQmxCreateDIChan;
    DAQmxCreateDOChan_t DAQmxCreateDOChan;
    DAQmxCfgSampClkTiming_t DAQmxCfgSampClkTiming;
    DAQmxStartTask_t DAQmxStartTask;
    DAQmxStopTask_t DAQmxStopTask;
    DAQmxClearTask_t DAQmxClearTask;
    DAQmxReadAnalogF64_t DAQmxReadAnalogF64;
    DAQmxReadDigitalLines_t DAQmxReadDigitalLines;
    DAQmxWriteDigitalLines_t DAQmxWriteDigitalLines;
    DAQmxGetExtendedErrorInfo_t DAQmxGetExtendedErrorInfo;
    DAQmxCreateLinScale_t DAQmxCreateLinScale;
    DAQmxCreateMapScale_t DAQmxCreateMapScale;
    DAQmxCreatePolynomialScale_t DAQmxCreatePolynomialScale;
    DAQmxCreateTableScale_t DAQmxCreateTableScale;
    DAQmxCalculateReversePolyCoeff_t DAQmxCalculateReversePolyCoeff;
    DAQmxCreateTask_t DAQmxCreateTask;
    DAQmxCreateAIVoltageChan_t DAQmxCreateAIVoltageChan;
    DAQmxCreateAIVoltageRMSChan_t DAQmxCreateAIVoltageRMSChan;
    DAQmxCreateAIVoltageChanWithExcit_t DAQmxCreateAIVoltageChanWithExcit;
    DAQmxCreateAIAccel4WireDCVoltageChan_t DAQmxCreateAIAccel4WireDCVoltageChan;
    DAQmxCreateAIAccelChan_t DAQmxCreateAIAccelChan;
    DAQmxCreateAIAccelChargeChan_t DAQmxCreateAIAccelChargeChan;
    DAQmxCreateAIBridgeChan_t DAQmxCreateAIBridgeChan;
    DAQmxCreateAIChargeChan_t DAQmxCreateAIChargeChan;
    DAQmxCreateAICurrentChan_t DAQmxCreateAICurrentChan;
    DAQmxCreateAICurrentRMSChan_t DAQmxCreateAICurrentRMSChan;
    DAQmxCreateAIForceBridgePolynomialChan_t DAQmxCreateAIForceBridgePolynomialChan;
    DAQmxCreateAIForceBridgeTableChan_t DAQmxCreateAIForceBridgeTableChan;
    DAQmxCreateAIForceBridgeTwoPointLinChan_t DAQmxCreateAIForceBridgeTwoPointLinChan;
    DAQmxCreateAIForceIEPEChan_t DAQmxCreateAIForceIEPEChan;
    DAQmxCreateAIFreqVoltageChan_t DAQmxCreateAIFreqVoltageChan;
    DAQmxCreateAIMicrophoneChan_t DAQmxCreateAIMicrophoneChan;
    DAQmxCreateAIPosEddyCurrProxProbeChan_t DAQmxCreateAIPosEddyCurrProxProbeChan;
    DAQmxCreateAIPosLVDTChan_t DAQmxCreateAIPosLVDTChan;
    DAQmxCreateAIPosRVDTChan_t DAQmxCreateAIPosRVDTChan;
    DAQmxCreateAIRTDChan_t DAQmxCreateAIRTDChan;
    DAQmxCreateAIResistanceChan_t DAQmxCreateAIResistanceChan;
    DAQmxCreateAIRosetteStrainGageChan_t DAQmxCreateAIRosetteStrainGageChan;
    DAQmxCreateAIStrainGageChan_t DAQmxCreateAIStrainGageChan;
    DAQmxCreateAITempBuiltInSensorChan_t DAQmxCreateAITempBuiltInSensorChan;
    DAQmxCreateAIThrmcplChan_t DAQmxCreateAIThrmcplChan;
    DAQmxCreateAIThrmstrChanIex_t DAQmxCreateAIThrmstrChanIex;
    DAQmxCreateAIThrmstrChanVex_t DAQmxCreateAIThrmstrChanVex;
    DAQmxCreateAITorqueBridgePolynomialChan_t DAQmxCreateAITorqueBridgePolynomialChan;
    DAQmxCreateAITorqueBridgeTableChan_t DAQmxCreateAITorqueBridgeTableChan;
    DAQmxCreateAITorqueBridgeTwoPointLinChan_t DAQmxCreateAITorqueBridgeTwoPointLinChan;
    DAQmxCreateAIVelocityIEPEChan_t DAQmxCreateAIVelocityIEPEChan;

    NISysCfgInitializeSession_t NISysCfgInitializeSession;
    NISysCfgCreateFilter_t NISysCfgCreateFilter;
    NISysCfgSetFilterProperty_t NISysCfgSetFilterProperty;
    NISysCfgCloseHandle_t NISysCfgCloseHandle;
    NISysCfgFindHardware_t NISysCfgFindHardware;
    NISysCfgNextResource_t NISysCfgNextResource;
    NISysCfgGetResourceProperty_t NISysCfgGetResourceProperty;
    NISysCfgGetResourceIndexedProperty_t NISysCfgGetResourceIndexedProperty;

private:
    LibraryManager();
    ~LibraryManager();

#ifdef _WIN32
    HMODULE ni_daqmx_library_handle;
    HMODULE ni_syscfg_library_handle;
#endif

    LibraryManager(const LibraryManager&) = delete;
    LibraryManager& operator=(const LibraryManager&) = delete;
};

}