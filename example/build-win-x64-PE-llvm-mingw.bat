@echo off
set GO111MODULE=off
set CC=E:\llvm-mingw\bin\clang.exe
set CGO_LDFLAGS=-Wl,-pdb=vm.pdb
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -a -ldflags "-s -w" -o vm.exe
pause
