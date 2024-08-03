package connection

import (
	"log"

	"github.com/HillcrestEnigma/mcbuild/packet"
)

func (c *Connection) HandleConfiguration() (err error) {
	log.Println("Write clientbound known packs")
	err = c.WriteClientboundKnownPacks()
	if err != nil {
		return
	}

	log.Println("Read serverbound known packs")
	err = c.ReadServerboundKnownPacks()
	if err != nil {
		return
	}

	// log.Println("Write registry data")
	// err = c.WriteRegistryData()
	// if err != nil {
	// 	return
	// }

	log.Println("Write finish configuration")
	err = c.WriteFinishConfiguration()
	if err != nil {
		return
	}

	err = c.ReadAckFinishConfiguration()
	if err != nil {
		return
	}

	return c.HandlePlay()
}

func (c *Connection) WriteClientboundKnownPacks() (err error) {
	p := packet.NewPacket(0x0E)

	err = p.WriteVarInt(1)
	if err != nil {
		return
	}

	err = p.WriteString("minecraft:core")
	if err != nil {
		return
	}

	// TODO: fix
	err = p.WriteString("whatever???")
	if err != nil {
		return
	}

	err = p.WriteString("1.21")
	if err != nil {
		return
	}

	return c.WritePacket(p)
}

// TODO: Technically we can just omit this function and choose to drop this packet
func (c *Connection) ReadServerboundKnownPacks() (err error) {
	p, err := c.AcceptPacket(0x07)
	if err != nil {
		return
	}

	packCount, err := p.ReadVarInt()
	if err != nil {
		return
	}

	log.Println("Pack count:", packCount)

	var stuff string
	for i := 0; i < packCount; i++ {
		stuff, err = p.ReadString()
		if err != nil {
			return
		}
		log.Println("Stuff:", stuff)
	}

	return
}

func (c *Connection) WriteRegistryData() (err error) {
	p := packet.NewPacket(0x07)

	// TODO: This is probably incorrect, fix
	err = p.WriteString("minecraft:core")
	if err != nil {
		return
	}

	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	return c.WritePacket(p)
}

func (c *Connection) WriteFinishConfiguration() (err error) {
	p := packet.NewPacket(0x03)

	return c.WritePacket(p)
}

func (c *Connection) ReadAckFinishConfiguration() (err error) {
	_, err = c.AcceptPacket(0x03)
	if err != nil {
		return
	}

	return
}
