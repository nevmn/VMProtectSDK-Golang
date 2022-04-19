cd $(dirname "$0")
CGO_LDFLAGS="-L. -lVMProtectSDK -Wl,-map,vm.map -w -s" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
go build -a -ldflags "-s -w" 
