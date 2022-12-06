#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

Check() {
    if grep -i "A breakpoint has been hit" dump.txt; then
        exit 0
    else
        exit 1
    fi
}

git clone https://github.com/robertmin1/Debugger && cd Debugger/testdata
export GOBIN="$PWD"
go install WordGenerator.go
cd ..
go mod init Debugger.go
go mod tidy

go run Debugger.go >> dump.txt & sleep 10; Check
