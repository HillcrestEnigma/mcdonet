package connection

import (
	"github.com/HillcrestEnigma/mcdonet/config"
	"github.com/HillcrestEnigma/mcdonet/packet"
)

func (c *Connection) handleConfiguration() (err error) {
	err = c.writeClientboundKnownPacks()
	if err != nil {
		return
	}

	err = c.readServerboundKnownPacks()
	if err != nil {
		return
	}

	err = c.writeRegistryData()
	if err != nil {
		return
	}

	err = c.writeFinishConfiguration()
	if err != nil {
		return
	}

	err = c.readAckFinishConfiguration()
	if err != nil {
		return
	}

	return c.handlePlay()
}

func (c *Connection) writeClientboundKnownPacks() (err error) {
	p := packet.NewPacket(0x0E)

	// Known Pack Count Array
	err = p.WriteVarInt(1)
	if err != nil {
		return
	}

	// Pack Namespace
	err = p.WriteString("minecraft")
	if err != nil {
		return
	}

	// Pack ID
	err = p.WriteString("core")
	if err != nil {
		return
	}

	// Pack Version
	err = p.WriteString("1.21.1")
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *Connection) readServerboundKnownPacks() (err error) {
	p, err := c.acceptPacket(0x07)
	if err != nil {
		return
	}

	packCount, err := p.ReadVarInt()
	if err != nil {
		return
	}

	for i := 0; i < int(packCount); i++ {
		_, err = p.ReadString()
		if err != nil {
			return err
		}

		_, err = p.ReadString()
		if err != nil {
			return err
		}

		_, err = p.ReadString()
		if err != nil {
			return err
		}
	}

	return
}

func (c *Connection) writeRegistryData() error {
	for registryIdentifier, entries := range config.RegistriesByIdentifier {
		p := packet.NewPacket(0x07)

		err := p.WriteString(registryIdentifier)
		if err != nil {
			return err
		}

		err = p.WriteVarInt(int32(len(entries)))
		if err != nil {
			return err
		}

		for entryIdentifier := range entries {
			err = p.WriteString(entryIdentifier)
			if err != nil {
				return err
			}

			err = p.WriteBool(false)
			if err != nil {
				return err
			}
		}

		err = c.writePacket(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Connection) writeFinishConfiguration() (err error) {
	p := packet.NewPacket(0x03)

	return c.writePacket(p)
}

func (c *Connection) readAckFinishConfiguration() (err error) {
	_, err = c.acceptPacket(0x03)
	if err != nil {
		return
	}

	return
}
