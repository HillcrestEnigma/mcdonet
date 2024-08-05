package connection

import "github.com/HillcrestEnigma/mcbuild/packet"

func (c *connection) handleLogin() (err error) {
	err = c.readLoginStart()
	if err != nil {
		return
	}

	err = c.writeLoginSuccess()
	if err != nil {
		return
	}

	err = c.readLoginAck()
	if err != nil {
		return
	}

	return c.handleConfiguration()
}

func (c *connection) readLoginStart() (err error) {
	p, err := c.readPacket(0x00)
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

	c.player = &connectionPlayer{
		Username: username,
		UUID:     id,
	}
	return
}

func (c *connection) writeLoginSuccess() (err error) {
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

	// Strict Error Handling
	err = p.WriteBool(true)
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *connection) readLoginAck() (err error) {
	_, err = c.readPacket(0x03)
	if err != nil {
		return
	}

	return
}
