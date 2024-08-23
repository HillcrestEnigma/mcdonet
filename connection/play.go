package connection

import (
	"time"

	"github.com/HillcrestEnigma/mcbuild/packet"
	"github.com/HillcrestEnigma/mcbuild/world/block"
	"github.com/HillcrestEnigma/mcbuild/world/chunk"
)

func (c *connection) handlePlay() (err error) {
	err = c.writeLoginPlay()
	if err != nil {
		return
	}

	// 13 = Start waiting for level chunks
	// 0 = No value for this event
	err = c.writeGameEvent(13, 0)
	if err != nil {
		return
	}

	// err = c.writeChunkDataAndUpdateLight()
	// if err != nil {
	// 	return
	// }

	dirtBlock, err := block.NewBlockByIdentifier("minecraft:dirt")
	if err != nil {
		return err
	}

	spawn := chunk.NewChunkColumn(0, 0, -64, 384)
	for x := uint8(0); x < 16; x++ {
		for z := uint8(0); z < 16; z++ {
			err = spawn.SetBlock(uint8(x), 0, uint8(z), dirtBlock)
			if err != nil {
				return
			}
		}
	}
	c.writeChunkDataAndUpdateLight(spawn)

	time.Sleep(100 * time.Second)

	return
}

func (c *connection) writeLoginPlay() (err error) {
	p := packet.NewPacket(0x2b)

	// Player Entity ID
	err = p.WriteInt32(0)
	if err != nil {
		return
	}

	// Gamemode is Hardcore
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	// Dimension Count Array
	err = p.WriteVarInt(1)
	if err != nil {
		return
	}

	// Dimension Identifier
	err = p.WriteString("minecraft:overworld")
	if err != nil {
		return
	}

	// Max Players
	err = p.WriteVarInt(20)
	if err != nil {
		return
	}

	// Render Distance
	err = p.WriteVarInt(2)
	if err != nil {
		return
	}

	// Simulation Distance
	err = p.WriteVarInt(2)
	if err != nil {
		return
	}

	// Reduced Debug Info
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	// Enable Respawn Screen
	err = p.WriteBool(true)
	if err != nil {
		return
	}

	// Do Limited Crafting
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	// Dimension Type
	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	// Dimension Name
	err = p.WriteString("minecraft:overworld")
	if err != nil {
		return
	}

	// Hashed Seed
	err = p.WriteInt64(0)
	if err != nil {
		return
	}

	// Gamemode
	err = p.WriteUInt8(0)
	if err != nil {
		return
	}

	// Previous Gamemode
	err = p.WriteInt8(-1)
	if err != nil {
		return
	}

	// Is Debug World
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	// Is Flat World
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	// Has Death Location
	err = p.WriteBool(false)
	if err != nil {
		return
	}
	// Omit 2 optional fields related to Has Death Location

	// Portal Cooldown
	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	// Enforces Secure Chat
	err = p.WriteBool(false)
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *connection) writeGameEvent(event uint8, value float32) (err error) {
	p := packet.NewPacket(0x22)

	// Event
	err = p.WriteUInt8(event)
	if err != nil {
		return
	}

	// Value
	err = p.WriteFloat32(value)
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *connection) writeChunkDataAndUpdateLight(chunkData *chunk.ChunkColumn) (err error) {
	p := packet.NewPacket(0x27)

	p.WriteChunk(chunkData)

	// TODO: implement block entity data
	p.WriteVarInt(0)

	// TODO: implement light data
	p.WriteVarInt(0)
	p.WriteVarInt(0)
	p.WriteVarInt(0)
	p.WriteVarInt(0)
	p.WriteVarInt(0)
	p.WriteVarInt(0)

	return c.writePacket(p)
}
