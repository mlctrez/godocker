#!/usr/bin/env bash

BINDIR=$(pwd)/bin

mkdir -p bin

echo "copying golang cert bundle"
# grab a copy of the latest ca-certificates bundle from the golang docker image
docker run --rm -v $BINDIR:/certcopy golang cp /etc/ssl/certs/ca-certificates.crt /certcopy/ca-bundle-golang.crt

if [ $? -ne 0 ]
then
    echo "error copying golang docker ca bundle"
    exit 1
fi

echo "copying amazonlinux cert bundle"
docker run --rm -v $BINDIR:/certcopy amazonlinux cp /etc/ssl/certs/ca-bundle.crt /certcopy/ca-bundle-amazonlinux.crt

if [ $? -ne 0 ]
then
    echo "error copying amazonlinux docker ca bundle"
    exit 1
fi

echo "building godocker binary"

# should use this probably, but it messes with the caching of the precompiled binaries
# CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-w' -o bin/godocker ./cli/godocker/main.go

# prime the cache - permissions of /usr/local/go/pkg/linux_amd64 may need to be changed to allow this
# GOOS=linux go install -a github.com/mlctrez/godocker/...

GOOS=linux go build -o bin/godocker ./cli/godocker/main.go

docker build --tag godocker .
