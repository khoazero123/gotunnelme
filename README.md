# Commands

## Get

    go get github.com/khoazero123/gotunnelme
    go run github.com/khoazero123/gotunnelme/main.go

## Build for Windows

    set GOARCH=amd64
    set GOOS=windows
    go build -o bin/lt.exe -ldflags "-s -w" github.com/khoazero123/gotunnelme

## Build for Linux

    set GOARCH=amd64
    set GOOS=linux
    go build -o bin/lt -ldflags "-s -w" github.com/khoazero123/gotunnelme
