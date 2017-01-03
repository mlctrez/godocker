#!/usr/bin/env bash

mkdir -p bin

CERTFILE=/etc/ssl/certs/ca-certificates.crt

# grab a copy of the latest ca-certificates bundle from the golang docker image
docker run --rm -v bin:/certcopy golang cp $CERTFILE /certcopy/

if [ $? -eq 0 ]
then
    echo "copied $CERTFILE to local directory"
else
    echo "error copying cert file $CERTFILE"
    exit 1
fi

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o bin/godocker ./cli/godocker/main.go

docker build --tag godocker .
