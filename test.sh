#!/usr/bin/env bash

git clone https://github.com/robertmin1/Debugger && cd "!$:t:r"
export GOBIN="$PWD"
go install WordGenerator.go
go run main.go