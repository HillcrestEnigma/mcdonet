package connection

import (
	_ "embed"
	"encoding/json"
	"log"
	"sync"

	"github.com/HillcrestEnigma/mcbuild/packet"
)

var (
	//go:embed registryData.json
	registryJSONData     []byte
	registryData         map[string][]string
	readRegistryDataOnce sync.Once
)

func (c *connection) handleConfig() (err error) {
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
	err = p.WriteString("1.21")
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *connection) readServerboundKnownPacks() (err error) {
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

func (c *connection) writeRegistryData() error {
	readRegistryDataOnce.Do(func() {
		err := json.Unmarshal(registryJSONData, &registryData)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	})

	for registryID, entries := range registryData {
		p := packet.NewPacket(0x07)

		err := p.WriteString(registryID)
		if err != nil {
			return err
		}

		err = p.WriteVarInt(int32(len(entries)))
		if err != nil {
			return err
		}

		for _, entryID := range entries {
			err = p.WriteString(entryID)
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
