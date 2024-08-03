package connection

import "github.com/HillcrestEnigma/mcbuild/packet"

func (c *connection) handlePlay() (err error) {
	err = c.writeLoginPlay()
	if err != nil {
		return
	}

	return
}

func (c *connection) writeLoginPlay() (err error) {
	p := packet.NewPacket(0x2B)

	// TODO: implement

	return c.writePacket(p)
}
