package main

import (
	"github.com/mlctrez/godocker/pkg/server"
	"log"
)

func main() {
	err := server.New().Start()
	if err != nil {
		log.Fatal(err)
	}
}
