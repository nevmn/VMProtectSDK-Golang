cd $(dirname "$0")
GO111MODULE=off CGO_LDFLAGS="-Wl,-map,vm.map -w -s" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
go build -a -ldflags "-s -w" -o vm
