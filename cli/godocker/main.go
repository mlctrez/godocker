package main

import (
	"log"

	"github.com/mlctrez/godocker/pkg/server"
)

func main() {
	err := server.New().Start()
	if err != nil {
		log.Fatal(err)
	}
}
