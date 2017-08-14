#!/usr/bin/env bash
go clean && go build -o agent.bin && clear && ./agent.bin --master=localhost:2999 $@
