package connection

import "github.com/HillcrestEnigma/mcbuild/packet"

func (c *Connection) HandleLogin() (err error) {
	err = c.ReadLoginStart()
	if err != nil {
		return
	}

	err = c.WriteLoginSuccess()
	if err != nil {
		return
	}

	err = c.ReadLoginAck()
	if err != nil {
		return
	}

	return c.HandleConfiguration()
}

func (c *Connection) ReadLoginStart() (err error) {
	p, err := c.ReadPacket(0x00)
	if err != nil {
		return
	}

	username, err := p.ReadString()
	if err != nil {
		return
	}

	id, err := p.ReadUUID()
	if err != nil {
		return
	}

	c.player = &ConnectionPlayer{
		Username: username,
		UUID:     id,
	}
	return
}

func (c *Connection) WriteLoginSuccess() (err error) {
	p := packet.NewPacket(0x02)

	err = p.WriteUUID(c.player.UUID)
	if err != nil {
		return
	}

	err = p.WriteString(c.player.Username)
	if err != nil {
		return
	}

	err = p.WriteVarInt(0) // TODO: implement properties
	if err != nil {
		return
	}

	err = p.WriteBool(true)
	if err != nil {
		return
	}

	return c.WritePacket(p)
}

func (c *Connection) ReadLoginAck() (err error) {
	_, err = c.ReadPacket(0x03)
	if err != nil {
		return
	}

	return
}