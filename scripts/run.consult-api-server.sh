#!/usr/bin/env bash

pushd ../cmd/consult-api-server
PORT=9999 go run main.go
popd