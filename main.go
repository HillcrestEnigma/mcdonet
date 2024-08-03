package main

import (
	"github.com/HillcrestEnigma/mcbuild/connection"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		netConn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		c := connection.NewConnection(netConn)

		go c.HandleConnection()
	}
}
