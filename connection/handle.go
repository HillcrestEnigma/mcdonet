package connection

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"unicode/utf16"
)

type handshake struct {
	protoVer   int32
	serverAddr string
	serverPort uint16
	nextState  int32
}

func (c *Connection) HandleConnection() {
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

func (c *Connection) readHandshake() (h *handshake, err error) {
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

func (c *Connection) handleLegacyServerListPing() (bool, error) {
	sig, err := c.buf.Peek(3)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(sig, []byte{0xFE, 0x01, 0xFA}) {
		return false, nil
	}

	info := []string{
		"§1",                   // Always Required
		"47",                   // Protocol Version
		"Greater than 1.7 pls", // Version
		"McDoNet Server",       // Server Name
		"0",                    // Online Players
		"20",                   // Max Players
	}

	response := make([]byte, 0)
	for idx, str := range info {
		encoded := utf16.Encode([]rune(str))
		for _, v := range encoded {
			response = binary.BigEndian.AppendUint16(response, v)
		}

		if idx != len(info)-1 {
			response = append(response, 0x00, 0x00)
		}
	}

	c.buf.WriteByte(0xFF)
	binary.Write(c.buf, binary.BigEndian, int16(len(response)/2))
	c.buf.Write(response)

	return true, c.buf.Flush()
}
