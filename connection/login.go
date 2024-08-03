package connection

import (
	"log"

	"github.com/google/uuid"
)

type LoginStart struct {
	Username string
	UUID     uuid.UUID
}

func (c *Connection) HandleLogin() error {
	loginStart, err := c.ReadLoginStart()
	if err != nil {
		return err
	}

	log.Println("LoginStart", loginStart)

	return nil
}

func (c *Connection) ReadLoginStart() (loginStart *LoginStart, err error) {
	packet, err := c.ReadPacket()
	if err != nil {
		return
	}

	username, err := packet.ReadString()
	if err != nil {
		return
	}

	id, err := packet.ReadUUID()
	if err != nil {
		return
	}

	loginStart = &LoginStart{
		Username: username,
		UUID:     id,
	}
	return
}
