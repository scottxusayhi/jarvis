#!/usr/bin/env bash
go clean && go build -o agent.bin && clear && ./agent.bin --master=10.1.10.99:2999 $@
