#!/usr/bin/env bash
go clean && GOOS=linux GOARCH=amd64 go build -o jarvis_agent_linux_amd64
