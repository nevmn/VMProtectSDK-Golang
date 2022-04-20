cd $(dirname "$0")
GO111MODULE=off CGO_LDFLAGS="-L. -lVMProtectSDK64 -Wl,-Map,vm.map" CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 \
go build -a -ldflags "-s -w" -o vm.exe
