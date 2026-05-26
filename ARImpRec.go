package ipmrec

import (
	"unsafe"
)

func (i *Ipmrec) UnpackPdataSection(MSNameOfProtected *int8, MSNameOfDumped *int8, MSWarning *int8) int32 {
	r1, _, _ := getProc("UnpackPdataSection@12").Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return int32(r1)
}

func (i *Ipmrec) GetNameFileOptimized(MSFileNameOrig *int8, MSFileNameOptimized *int8) int32 {
	r1, _, _ := getProc("GetNameFileOptimized@8").Call(uintptr(unsafe.Pointer(MSFileNameOrig)), uintptr(unsafe.Pointer(MSFileNameOptimized)))
	return int32(r1)
}

func (i *Ipmrec) RebuildSectionsFromArmadillo(MSNameOfProtected *int8, MSNameOfDumped *int8, MSWarning *int8) int32 {
	r1, _, _ := getProc("RebuildSectionsFromArmadillo@12").Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return int32(r1)
}

func (i *Ipmrec) TryGetImportedFunction(IRProcessId uint32, IRVAddress uint32, IROrdinal **uint32, IRHint **uint32, IRFunctionName *int8, IRModule *int8) int32 {
	r1, _, _ := getProc("TryGetImportedFunction@24").Call(uintptr(IRProcessId), uintptr(IRVAddress), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRFunctionName)), uintptr(unsafe.Pointer(IRModule)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsNoNewSection(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := getProc("SearchAndRebuildImportsNoNewSection@28").Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsIATOptimized(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := getProc("SearchAndRebuildImportsIATOptimized@28").Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImports(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := getProc("SearchAndRebuildImports@28").Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) GetProcNameAndOrdinal(IRHModule uintptr, IRAddress uint32, IROrdinal **uint32, IRHint **uint32, IRProcName *int8) int32 {
	r1, _, _ := getProc("GetProcNameAndOrdinal@20").Call(IRHModule, uintptr(IRAddress), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return int32(r1)
}

func (i *Ipmrec) GetProcOrdinal(IRHModule uintptr, IRAddress uint32) uint32 {
	r1, _, _ := getProc("GetProcOrdinal@8").Call(IRHModule, uintptr(IRAddress))
	return uint32(r1)
}

func (i *Ipmrec) GetProcName(IRHModule uintptr, IRAddress uint32, IRHint **uint32, IRProcName *int8) int32 {
	r1, _, _ := getProc("GetProcName@16").Call(IRHModule, uintptr(IRAddress), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return int32(r1)
}

func (i *Ipmrec) GetAllVAddressesOfImports(IRProcessId uint32, IROEP uint32, IRVAddressImports **uint32, IRNumberOfImports int32) int32 {
	r1, _, _ := getProc("GetAllVAddressesOfImports@16").Call(uintptr(IRProcessId), uintptr(IROEP), uintptr(unsafe.Pointer(IRVAddressImports)), uintptr(IRNumberOfImports))
	return int32(r1)
}
