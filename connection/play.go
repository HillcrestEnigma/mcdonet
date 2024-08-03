package connection

import "github.com/HillcrestEnigma/mcbuild/packet"

func (c *Connection) HandlePlay() (err error) {
	err = c.WriteLoginPlay()
	if err != nil {
		return
	}

	return
}

func (c *Connection) WriteLoginPlay() (err error) {
	p := packet.NewPacket(0x2B)

	// TODO: implement

	return c.WritePacket(p)
}
