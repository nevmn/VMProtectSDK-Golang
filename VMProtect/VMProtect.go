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

const (
	SERIAL_STATE_SUCCESS                = 0
	SERIAL_STATE_FLAG_CORRUPTED         = 0x00000001
	SERIAL_STATE_FLAG_INVALID           = 0x00000002
	SERIAL_STATE_FLAG_BLACKLISTED       = 0x00000004
	SERIAL_STATE_FLAG_DATE_EXPIRED      = 0x00000008
	SERIAL_STATE_FLAG_RUNNING_TIME_OVER = 0x00000010
	SERIAL_STATE_FLAG_BAD_HWID          = 0x00000020
	SERIAL_STATE_FLAG_MAX_BUILD_EXPIRED = 0x00000040
)

//go:linkname vmprotectBegin VMProtectBegin
//go:noescape
func vmprotectBegin(*string) unsafe.Pointer

//go:linkname vmprotectBeginVirtualization VMProtectBeginVirtualization
//go:noescape
func vmprotectBeginVirtualization(*string) unsafe.Pointer

//go:linkname vmprotectBeginMutation VMProtectBeginMutation
//go:noescape
func vmprotectBeginMutation(*string) unsafe.Pointer

//go:linkname vmprotectBeginUltra VMProtectBeginUltra
//go:noescape
func vmprotectBeginUltra(*string) unsafe.Pointer

//go:linkname vmprotectBeginUltraLockByKey VMProtectBeginUltraLockByKey
//go:noescape
func vmprotectBeginUltraLockByKey(*string) unsafe.Pointer

//go:linkname vmprotectBeginVirtualizationLockByKey VMProtectBeginVirtualizationLockByKey
//go:noescape
func vmprotectBeginVirtualizationLockByKey(*string) unsafe.Pointer

//go:linkname End VMProtectEnd
//go:noescape
func End()

//go:linkname vmprotectBeginD VMProtectBegin
//go:noescape
func vmprotectBeginD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectBeginVirtualizationD VMProtectBeginVirtualization
//go:noescape
func vmprotectBeginVirtualizationD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectBeginMutationD VMProtectBeginMutation
//go:noescape
func vmprotectBeginMutationD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectBeginUltraD VMProtectBeginUltra
//go:noescape
func vmprotectBeginUltraD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectBeginUltraLockByKeyD VMProtectBeginUltraLockByKey
//go:noescape
func vmprotectBeginUltraLockByKeyD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectBeginVirtualizationLockByKeyD VMProtectBeginVirtualizationLockByKey
//go:noescape
func vmprotectBeginVirtualizationLockByKeyD(string, *string, string) unsafe.Pointer

//go:linkname vmprotectDecryptStringA VMProtectDecryptStringA
//go:noescape
func vmprotectDecryptStringA(*string) *C.char

//go:linkname vmprotectDecryptStringAD VMProtectDecryptStringA
//go:noescape
func vmprotectDecryptStringAD(string, *string, string) *C.char

//go:linkname call runtime.asmcgocall
//go:noescape
func call(fn, arg unsafe.Pointer) int32

//go:linkname callbool runtime.asmcgocall
//go:noescape
func callbool(fn, arg unsafe.Pointer) bool

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
	b := []byte(serial)
	cserial := (*C.char)(unsafe.Pointer(&b[0]))
	return int(call(C.VMProtectSetSerialNumber, unsafe.Pointer(cserial)))
}

func GetSerialNumberState() int {
	return int(call(C.VMProtectGetSerialNumberState, unsafe.Pointer(nil)))
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

func Begin(MarkerName string) unsafe.Pointer {
	if runtime.GOOS == "windows" {
		return vmprotectBegin(&MarkerName)
	} else {
		return vmprotectBeginD("", nil, MarkerName)
	}
}

func BeginVirtualization(MarkerName string) unsafe.Pointer {
	if runtime.GOOS == "windows" {
		return vmprotectBeginVirtualization(&MarkerName)
	} else {
		return vmprotectBeginVirtualizationD("", nil, MarkerName)
	}
}

func BeginMutation(MarkerName string) unsafe.Pointer {
	if runtime.GOOS == "windows" {
		return vmprotectBeginMutation(&MarkerName)
	} else {
		return vmprotectBeginMutationD("", nil, MarkerName)
	}
}

func BeginUltra(MarkerName string) unsafe.Pointer {
	if runtime.GOOS == "windows" {
		return vmprotectBeginUltra(&MarkerName)
	} else {
		return vmprotectBeginUltraD("", nil, MarkerName)
	}
}

func BeginUltraLockByKey(MarkerName string) unsafe.Pointer {
	if runtime.GOOS == "windows" {
		return vmprotectBeginUltraLockByKey(&MarkerName)
	} else {
		return vmprotectBeginUltraLockByKeyD("", nil, MarkerName)
	}
}

func BeginVirtualizationLockByKey(MarkerName string) unsafe.Pointer {
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
	return bool(callbool(C.VMProtectIsDebuggerPresent, unsafe.Pointer(&CheckKernelMode)))
}

func IsVirtualMachinePresent() bool {
	return bool(callbool(C.VMProtectIsVirtualMachinePresent, unsafe.Pointer(nil)))
}

func IsProtected() bool {
	return bool(callbool(C.VMProtectIsProtected, unsafe.Pointer(nil)))
}

func IsValidImageCRC() bool {
	return bool(callbool(C.VMProtectIsValidImageCRC, unsafe.Pointer(nil)))
}
