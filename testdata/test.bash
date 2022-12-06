#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

git clone https://github.com/robertmin1/Debugger && cd Debugger/testdata
export GOBIN="$PWD"
go install WordGenerator.go

timeout 10s go run main.go
