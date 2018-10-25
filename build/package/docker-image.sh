#!/usr/bin/env bash

BASE_PATH="`dirname \"$0\"`"

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/theservice "${BASE_PATH}/../../"

docker build -t theservice -f "${BASE_PATH}/Dockerfile.scratch" .