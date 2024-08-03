package connection

import (
	"fmt"
)

type Handshake struct {
	protoVer   int
	serverAddr string
	serverPort uint16
	nextState  int
}

func (c *Connection) ReadHandshake() (handshake *Handshake, err error) {
	p, err := c.ReadPacket(0x00)
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

	handshake = &Handshake{
		protoVer:   protoVer,
		serverAddr: serverAddr,
		serverPort: serverPort,
		nextState:  nextState,
	}

	return
}
