#!/usr/bin/env bash

set -euo pipefail

git clone https://github.com/robertmin1/Debugger && cd Debugger
export GOBIN="$PWD"
go install WordGenerator.go
timeout 10s go run main.go
