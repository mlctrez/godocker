package main

import (
	"github.com/mlctrez/godocker/pkg/server"
	"log"
)

func main() {
	log.Fatal(server.New().Start())
}
