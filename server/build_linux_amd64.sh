#!/usr/bin/env bash
go clean && GOOS=linux GOARCH=amd64 go build -o jarvis_master_linux_amd64.bin
cp jarvis_master_linux_amd64.bin docker_release/
