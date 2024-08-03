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
	packet, err := c.ReadPacket()
	if err != nil {
		return
	}

	protoVer, err := packet.ReadVarInt()
	if err != nil {
		return
	}

	serverAddr, err := packet.ReadString()
	if err != nil {
		return
	}

	serverPort, err := packet.ReadUnsignedShort()
	if err != nil {
		return
	}

	nextState, err := packet.ReadVarInt()
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
