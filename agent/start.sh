#!/usr/bin/env bash
go clean && go build && clear && ./agent --master=10.1.10.99:2999 $@
