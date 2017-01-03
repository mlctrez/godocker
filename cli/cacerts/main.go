package main

import (
	"os"
	"github.com/mlctrez/godocker/pkg/cacerts"
)

func main() {
	p1 := cacerts.New()
	err := p1.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}
	p2 := cacerts.New()
	err = p2.Parse(os.Args[2])

	if err != nil {
		panic(err)
	}
	cacerts.Compare(p1, p2)
}
