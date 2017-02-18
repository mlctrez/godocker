#!/usr/bin/env bash

BINDIR=$(pwd)/bin

mkdir -p bin

echo "copying amazonlinux cert bundle"

docker run --rm -v $BINDIR:/certcopy amazonlinux cp /etc/ssl/certs/ca-bundle.crt /certcopy/ca-certificates.crt

if [ $? -ne 0 ]
then
    echo "error copying amazonlinux docker ca bundle"
    exit 1
fi

echo "building godocker binary"

CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-w' -o bin/godocker ./cli/godocker/main.go

docker build --tag godocker .
