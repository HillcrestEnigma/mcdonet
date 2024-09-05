package main

import (
	"net"

	"github.com/HillcrestEnigma/mcdonet/connection"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		netConn, err := l.Accept()
		if err != nil {
			return err
		}

		c := connection.NewConnection(netConn)
		go c.HandleConnection()
	}
}
