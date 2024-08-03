package connection

import (
	"log"

	"github.com/HillcrestEnigma/mcbuild/packet"
)

func (c *connection) handleConfig() (err error) {
	log.Println("Write clientbound known packs")
	err = c.writeClientboundKnownPacks()
	if err != nil {
		return
	}

	// log.Println("Write registry data")
	// err = c.WriteRegistryData()
	// if err != nil {
	// 	return
	// }

	log.Println("Write finish configuration")
	err = c.writeFinishConfig()
	if err != nil {
		return
	}

	err = c.readAckFinishConfig()
	if err != nil {
		return
	}

	return c.handlePlay()
}

func (c *connection) writeClientboundKnownPacks() (err error) {
	p := packet.NewPacket(0x0E)

	// Known Pack Count Array
	err = p.WriteVarInt(1)
	if err != nil {
		return
	}

	// Pack Namespace
	err = p.WriteString("minecraft:core")
	if err != nil {
		return
	}

	// TODO: fix
	// Pack ID
	err = p.WriteString("minecraft:core")
	if err != nil {
		return
	}

	// Pack Version
	err = p.WriteString("1.21")
	if err != nil {
		return
	}

	return c.writePacket(p)
}

// func (c *Connection) writeRegistryData() (err error) {
// 	p := packet.NewPacket(0x07)

// 	// TODO: This is probably incorrect, fix
// 	err = p.WriteString("minecraft:core")
// 	if err != nil {
// 		return
// 	}

// 	err = p.WriteVarInt(0)
// 	if err != nil {
// 		return
// 	}

// 	return c.WritePacket(p)
// }

func (c *connection) writeFinishConfig() (err error) {
	p := packet.NewPacket(0x03)

	return c.writePacket(p)
}

func (c *connection) readAckFinishConfig() (err error) {
	_, err = c.acceptPacket(0x03)
	if err != nil {
		return
	}

	return
}
