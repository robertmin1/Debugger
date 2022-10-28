#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

git clone https://github.com/robertmin1/Debugger && cd Debugger
export GOBIN="$PWD"

timeout 10s go run WordGenerator.go
