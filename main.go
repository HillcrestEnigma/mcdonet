package main

import (
	"log"
	"net"

	"github.com/HillcrestEnigma/mcbuild/connection"
)

func main() {
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
