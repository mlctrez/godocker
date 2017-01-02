#!/usr/bin/env bash

mkdir -p bin

GOOS=linux GOARCH=amd64 go build -o bin/godocker ./cli/...

docker build --tag godocker .
