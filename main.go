package main

import (
	"log"
	"net"

	"github.com/HillcrestEnigma/mcbuild/connection"
	"github.com/HillcrestEnigma/mcbuild/registry"
)

func main() {
	registry.LoadRegistryData()

	l, err := net.Listen("tcp", ":2000")

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connection.HandleConnection(c)
	}
}
