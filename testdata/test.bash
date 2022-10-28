#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

timeout 10s go run main.go
