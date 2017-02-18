# godocker

[![Go Report Card](https://goreportcard.com/badge/github.com/mlctrez/godocker)](https://goreportcard.com/report/github.com/mlctrez/godocker)

An example for building a minimal golang docker image with ca certificate support.

The ca-certs bundle is copied from the amazon linux docker image and placed in the appropriate directory.

Locations of cacert bundles for linux systems can be found in the go source at [crypto/x509/root_linux.go](https://golang.org/src/crypto/x509/root_linux.go)

Better described [here](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)