#!/usr/bin/env bash
go clean && GOOS=linux GOARCH=amd64 go build -o jarvis_server_linux_amd64.bin
#scp jarvis_server_linux_amd64.bin 192.168.130.42:~/
