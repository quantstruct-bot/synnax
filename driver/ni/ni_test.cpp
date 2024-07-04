#include <stdio.h>
#include "daqmx.h"
#include "lib_manager.h"

#define DAQmxErrChk(functionCall) if( DAQmxFailed(error=(functionCall)) ) goto Error; else

int main(void)
{
    ni::LibraryManager& m = ni::LibraryManager::getInstance();
    
    int32       error=0;
    TaskHandle  taskHandle=0;
    TaskHandle  taskHandle2=0;
    int32       numRead;
    uInt8       data[8000];
    float64       data2[8000];
    char        errBuff[2048]={'\0'};

    /*********************************************/
    // DAQmx Configure Code
    /*********************************************/
    DAQmxErrChk (m.DAQmxCreateTask("",&taskHandle));

    DAQmxErrChk (m.DAQmxCreateDIChan(taskHandle,"Dev1/port0/line0","",DAQmx_Val_ChanPerLine));

    for(int i = 0; i < 1; i++) {
        DAQmxErrChk (m.DAQmxReadDigitalLines(taskHandle,1000,10.0,DAQmx_Val_GroupByChannel,data,1000,&numRead,NULL,NULL));

        printf("Acquired %d samples\n",(int)numRead);

        for(int i = 0; i < numRead; i++) 
            printf("Sample %d: %d\n", i, data[i]);
    }


    Error:
    if( DAQmxFailed(error) )
        DAQmxGetExtendedErrorInfo(errBuff,2048);
    if( taskHandle!=0 ) {
        /*********************************************/
        // DAQmx Stop Code
        /*********************************************/
        DAQmxStopTask(taskHandle);
        DAQmxClearTask(taskHandle);
    }
    if( DAQmxFailed(error) )
        printf("DAQmx Error: %s\n",errBuff);
    printf("End of program, press Enter key to quit\n");
    getchar();
    return 0;
}
