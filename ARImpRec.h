#pragma once

#ifdef __cplusplus
extern "C" {
#endif

// 类型定义
typedef unsigned int Cardinal;
typedef int Integer;
typedef char* PChar;
typedef Cardinal* PCardinal;
typedef void* HMODULE;
typedef Cardinal IRSaveOEPToFile;

// ARImpRec.dll 导出函数声明
__declspec(dllimport) Integer __stdcall UnpackPdataSection@12(PChar MSNameOfProtected, PChar MSNameOfDumped, PChar MSWarning);
__declspec(dllimport) Integer __stdcall GetNameFileOptimized@8(PChar MSFileNameOrig, PChar MSFileNameOptimized);
__declspec(dllimport) Integer __stdcall RebuildSectionsFromArmadillo@12(PChar MSNameOfProtected, PChar MSNameOfDumped, PChar MSWarning);
__declspec(dllimport) Integer __stdcall TryGetImportedFunction@24(Cardinal IRProcessId, Cardinal IRVAddress, PCardinal* IROrdinal, PCardinal* IRHint, PChar IRFunctionName, PChar IRModule);
__declspec(dllimport) Integer __stdcall SearchAndRebuildImportsNoNewSection@28(Cardinal IRProcessId, PChar IRNameOfDumped, Cardinal IROEP, IRSaveOEPToFile IRSaveOEPToFile, PCardinal* IRIATRVA, PCardinal* IRIATSize, PChar IRWarning);
__declspec(dllimport) Integer __stdcall SearchAndRebuildImportsIATOptimized@28(Cardinal IRProcessId, PChar IRNameOfDumped, Cardinal IROEP, IRSaveOEPToFile IRSaveOEPToFile, PCardinal* IRIATRVA, PCardinal* IRIATSize, PChar IRWarning);
__declspec(dllimport) Integer __stdcall SearchAndRebuildImports@28(Cardinal IRProcessId, PChar IRNameOfDumped, Cardinal IROEP, IRSaveOEPToFile IRSaveOEPToFile, PCardinal* IRIATRVA, PCardinal* IRIATSize, PChar IRWarning);
__declspec(dllimport) Integer __stdcall GetProcNameAndOrdinal@20(HMODULE IRHModule, Cardinal IRAddress, PCardinal* IROrdinal, PCardinal* IRHint, PChar IRProcName);
__declspec(dllimport) Cardinal __stdcall GetProcOrdinal@8(HMODULE IRHModule, Cardinal IRAddress);
__declspec(dllimport) Integer __stdcall GetProcName@16(HMODULE IRHModule, Cardinal IRAddress, PCardinal* IRHint, PChar IRProcName);
__declspec(dllimport) Integer __stdcall GetAllVAddressesOfImports@16(Cardinal IRProcessId, Cardinal IROEP, PCardinal* IRVAddressImports, Integer IRNumberOfImports);

#ifdef __cplusplus
}
#endif
