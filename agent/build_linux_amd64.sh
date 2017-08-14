#!/usr/bin/env bash
go clean && GOOS=linux GOARCH=amd64 go build -o jarvis_agent_linux_amd64.bin
scp jarvis_agent_linux_amd64.bin k2data@192.168.130.42:~/jarvis-agent

