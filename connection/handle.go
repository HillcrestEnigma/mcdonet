package connection

import (
	"io"
	"log"
)

func (c *Connection) HandleConnection() {
	defer c.Close()

	isLegacy, err := c.HandleLegacyServerListPing()
	if err != nil {
		return
	}
	if isLegacy {
		return
	}

	handshake, err := c.ReadHandshake()
	if err != nil {
		return
	}

	switch handshake.nextState {
	case 1:
		err = c.HandleServerListPing()
		// case 2:
		// 	err = c.HandleLogin()
	}

	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed by client")
		} else {
			log.Println(err)
		}
	}
}
