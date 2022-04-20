# VMProtectSDK-Golang
``UnOfficial VMProtectSDK for Golang``

````
Due to the particularity of Golang, VMP is not compatible with it. 

For example, there is no 0 at the end of the Golang string, 

and ABI is non-standard(GO x64 ABI: RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11)

If use CGO, VMP does not recognize Marker and the string need to encrypt,

After a while of research, I solved the compatibility of some VMProtect with Golang.

Now VMProtect can recognize the address of VMProtectBegin and VMProtectEnd.

MarkerName and the string need to encrypt all can be detected(work well on mac and win).

Because of I didn't have Web License Manager, so the correlation function has not been implemented.

Most functions are Bind.


* Test work well on window 11 and Mac OS 12.3.1 (Go 1.18.1)

* But please note that due to the use of some unconventional methods, it may be unsafe.



Tip:

1.\x00 or \000 must be added after string ,like VMProtect.BeginUltra("Marker\x00").

2.Don't use the -gcflags "-N -l" command to compile , Otherwise VMP cannot recognize the Marker.

3.Must use -ldflags "-s -w" to strip the Symbol，Otherwise VMP cannot recognize the file.

4.Refer to the files in the example folder and modify the GoPath.

5.If necessary, add scripts in script.lua for additional protection.

6.Linux has not been tested. Maybe, but some of the code needs to be modified.

7."VMProtect.SetSerialNumber" This function has a probability 
  that the program will crash before it is not be protected (only on Windows).

8.You must use VMProtect.GoString to convert char to string,not C.GoString.

9.You can't write the command "#cgo windows LDFLAGS: -L${SRCDIR} -lVMProtectSDK64" in the Go file,
  Instead, it should be written in the compilation script, Otherwise VMP cannot recognize the Marker.
````
