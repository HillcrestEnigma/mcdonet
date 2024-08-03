package connection

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type handshake struct {
	protoVer   int
	serverAddr string
	serverPort uint16
	nextState  int
}

func HandleConnection(netConn net.Conn) {
	c := &connection{
		player: nil,
		net:    netConn,
		buf:    bufio.NewReadWriter(bufio.NewReader(netConn), bufio.NewWriter(netConn)),
	}

	defer c.close()

	isLegacy, err := c.handleLegacyServerListPing()
	if err != nil {
		return
	}
	if isLegacy {
		return
	}

	handshake, err := c.readHandshake()
	if err != nil {
		return
	}

	switch handshake.nextState {
	case 1:
		err = c.handleServerListPing()
	case 2:
		err = c.handleLogin()
	}

	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed by client")
		} else {
			log.Println(err)
		}
	}
}

func (c *connection) readHandshake() (h *handshake, err error) {
	p, err := c.readPacket(0x00)
	if err != nil {
		return
	}

	protoVer, err := p.ReadVarInt()
	if err != nil {
		return
	}

	serverAddr, err := p.ReadString()
	if err != nil {
		return
	}

	serverPort, err := p.ReadUInt16()
	if err != nil {
		return
	}

	nextState, err := p.ReadVarInt()
	if err != nil {
		return
	}

	if nextState < 1 || nextState > 3 {
		err = fmt.Errorf("invalid next state %d", nextState)
		return
	}

	h = &handshake{
		protoVer:   protoVer,
		serverAddr: serverAddr,
		serverPort: serverPort,
		nextState:  nextState,
	}

	return
}
