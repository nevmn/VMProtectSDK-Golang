package main

import (
	"VMProtect"
)

func main() {
	VMProtect.BeginUltra("Marker\x00")
	str := VMProtect.GoString(VMProtect.DecryptStringA("This is a decrypted string\x00"))
	serial := "SerialNumber\x00"

	println(str)
	println("HWID: ", VMProtect.GetCurrentHWID())
	println("IsProtected: ", VMProtect.IsProtected())
	println("IsDebuggerPresent: ", VMProtect.IsDebuggerPresent(true))
	println("IsVirtualMachinePresent: ", VMProtect.IsVirtualMachinePresent())
	println("IsValidImageCRC: ", VMProtect.IsValidImageCRC())
	println("SetSerialNumber: ", VMProtect.SetSerialNumber(serial))
	if VMProtect.GetSerialNumberState() == VMProtect.SERIAL_STATE_SUCCESS {
		println("-- Registered --")
	}
	println("User: ", VMProtect.GetUser())
	println("Email: ", VMProtect.GetEmail())
	println("ExpireDate: ", VMProtect.GetExpireDate())
	println("MaxBuildDate: ", VMProtect.GetMaxBuild())
	VMProtect.End()
}
