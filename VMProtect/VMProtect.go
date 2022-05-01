package VMProtect

/*
#include <stdbool.h>
#include "VMProtectSDK.h"
*/
import "C"

import (
	"runtime"
	"strconv"
	"unsafe"
)

//go:linkname vmprotectBegin VMProtectBegin
//go:noescape
func vmprotectBegin(*string) C.void

//go:linkname vmprotectBeginVirtualization VMProtectBeginVirtualization
//go:noescape
func vmprotectBeginVirtualization(*string) C.void

//go:linkname vmprotectBeginMutation VMProtectBeginMutation
//go:noescape
func vmprotectBeginMutation(*string) C.void

//go:linkname vmprotectBeginUltra VMProtectBeginUltra
//go:noescape
func vmprotectBeginUltra(*string) C.void

//go:linkname vmprotectBeginUltraLockByKey VMProtectBeginUltraLockByKey
//go:noescape
func vmprotectBeginUltraLockByKey(*string) C.void

//go:linkname vmprotectBeginVirtualizationLockByKey VMProtectBeginVirtualizationLockByKey
//go:noescape
func vmprotectBeginVirtualizationLockByKey(*string) C.void

//go:linkname End VMProtectEnd
//go:noescape
func End()

//go:linkname vmprotectBeginD VMProtectBegin
//go:noescape
func vmprotectBeginD(string, *string, string) C.void

//go:linkname vmprotectBeginVirtualizationD VMProtectBeginVirtualization
//go:noescape
func vmprotectBeginVirtualizationD(string, *string, string) C.void

//go:linkname vmprotectBeginMutationD VMProtectBeginMutation
//go:noescape
func vmprotectBeginMutationD(string, *string, string) C.void

//go:linkname vmprotectBeginUltraD VMProtectBeginUltra
//go:noescape
func vmprotectBeginUltraD(string, *string, string) C.void

//go:linkname vmprotectBeginUltraLockByKeyD VMProtectBeginUltraLockByKey
//go:noescape
func vmprotectBeginUltraLockByKeyD(string, *string, string) C.void

//go:linkname vmprotectBeginVirtualizationLockByKeyD VMProtectBeginVirtualizationLockByKey
//go:noescape
func vmprotectBeginVirtualizationLockByKeyD(string, *string, string) C.void

//go:linkname GetSerialNumberState VMProtectGetSerialNumberState
//go:noescape
func GetSerialNumberState() int

//go:linkname vmprotectDecryptStringA VMProtectDecryptStringA
//go:noescape
func vmprotectDecryptStringA(*string) *C.char

//go:linkname vmprotectDecryptStringAD VMProtectDecryptStringA
//go:noescape
func vmprotectDecryptStringAD(string, *string, string) *C.char

//go:linkname vmprotectSetSerialNumber VMProtectSetSerialNumber
func vmprotectSetSerialNumber(*string) int

//go:linkname vmprotectSetSerialNumberD VMProtectSetSerialNumber
func vmprotectSetSerialNumberD(string, *string, string) int

//go:linkname IsProtected VMProtectIsProtected
//go:noescape
func IsProtected() bool

//go:linkname IsValidImageCRC VMProtectIsValidImageCRC
//go:noescape
func IsValidImageCRC() bool

func GoString(cchar *C.char) string {
	return C.GoString(cchar)
}

func GetCurrentHWID() (hwid string) {
	nSize := C.VMProtectGetCurrentHWID(nil, 0)
	b := make([]byte, nSize)
	hw := (*C.char)(unsafe.Pointer(&b))
	C.VMProtectGetCurrentHWID(hw, nSize)
	hwid = C.GoStringN(hw, nSize)
	return hwid
}

func SetSerialNumber(serial string) int {
	if runtime.GOOS == "windows" {
		return vmprotectSetSerialNumber(&serial)
	} else {
		return vmprotectSetSerialNumberD("", nil, serial)
	}
}

func GetUser() (user string) {
	var sd C.VMProtectSerialNumberData
	if C.VMProtectGetSerialNumberData(&sd, C.sizeof_VMProtectSerialNumberData) {
		for _, v := range sd.wUserName {
			if v != 0 {
				user += string(v)
			}
		}
		return user
	}
	return
}

func GetEmail() (email string) {
	var sd C.VMProtectSerialNumberData
	if C.VMProtectGetSerialNumberData(&sd, C.sizeof_VMProtectSerialNumberData) {
		for _, v := range sd.wEMail {
			if v != 0 {
				email += string(v)
			}
		}
		return email
	}
	return
}

func GetExpireDate() (date string) {
	var sd C.VMProtectSerialNumberData
	if C.VMProtectGetSerialNumberData(&sd, C.sizeof_VMProtectSerialNumberData) {
		year := *(*uint16)(unsafe.Pointer(&sd.dtExpire.wYear))
		month := *(*uint8)(unsafe.Pointer(&sd.dtExpire.bMonth))
		day := *(*uint8)(unsafe.Pointer(&sd.dtExpire.bDay))
		date := strconv.Itoa(int(year)) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(int(day))
		return date
	}
	return
}

func GetMaxBuild() (date string) {
	var sd C.VMProtectSerialNumberData
	if C.VMProtectGetSerialNumberData(&sd, C.sizeof_VMProtectSerialNumberData) {
		year := *(*uint16)(unsafe.Pointer(&sd.dtMaxBuild.wYear))
		month := *(*uint8)(unsafe.Pointer(&sd.dtMaxBuild.bMonth))
		day := *(*uint8)(unsafe.Pointer(&sd.dtMaxBuild.bDay))
		date := strconv.Itoa(int(year)) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(int(day))
		return date
	}
	return
}

func Begin(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBegin(&MarkerName)
	} else {
		return vmprotectBeginD("", nil, MarkerName)
	}
}

func BeginVirtualization(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBeginVirtualization(&MarkerName)
	} else {
		return vmprotectBeginVirtualizationD("", nil, MarkerName)
	}
}

func BeginMutation(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBeginMutation(&MarkerName)
	} else {
		return vmprotectBeginMutationD("", nil, MarkerName)
	}
}

func BeginUltra(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBeginUltra(&MarkerName)
	} else {
		return vmprotectBeginUltraD("", nil, MarkerName)
	}
}

func BeginUltraLockByKey(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBeginUltraLockByKey(&MarkerName)
	} else {
		return vmprotectBeginUltraLockByKeyD("", nil, MarkerName)
	}
}

func BeginVirtualizationLockByKey(MarkerName string) C.void {
	if runtime.GOOS == "windows" {
		return vmprotectBeginVirtualizationLockByKey(&MarkerName)
	} else {
		return vmprotectBeginVirtualizationLockByKeyD("", nil, MarkerName)
	}
}

func DecryptStringA(EncryptStr string) (DecryptStr *C.char) {
	if runtime.GOOS == "windows" {
		return vmprotectDecryptStringA(&EncryptStr)
	} else {
		return vmprotectDecryptStringAD("", nil, EncryptStr)
	}
}

func IsDebuggerPresent(CheckKernelMode bool) bool {
	return bool(C.VMProtectIsDebuggerPresent(C.bool(CheckKernelMode)))
}

func IsVirtualMachinePresent() bool {
	return bool(C.VMProtectIsVirtualMachinePresent())
}
