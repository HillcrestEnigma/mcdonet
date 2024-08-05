package connection

import (
	"log"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/HillcrestEnigma/mcbuild/packet"
)

func (c *connection) handlePlay() (err error) {
	err = c.writeLoginPlay()
	if err != nil {
		return
	}

	// 13 = Start waiting for level chunks
	err = c.writeGameEvent(13, 0)

	return
}

func (c *connection) writeLoginPlay() (err error) {
	log.Println("Write login play")
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

func (c *connection) writeChunkDataAndUpdateLight() (err error) {
	p := packet.NewPacket(0x27)

	// Chunk X
	err = p.WriteInt32(0)
	if err != nil {
		return
	}

	// Chunk Z
	err = p.WriteInt32(0)
	if err != nil {
		return
	}

	// Heightmaps
	heightmap := make([]int64, 16*16)
	for i := range heightmap {
		heightmap[i] = 65
	}

	heightmaps := &datatype.NBT{
		Name: "",
		Compound: &datatype.NBTCompound{
			"MOTION_BLOCKING": &heightmap,
			"WORLD_SURFACE":   &heightmap,
		},
	}

	err = p.WriteNBT(heightmaps)
	if err != nil {
		return
	}

	// Full Chunk
	err = p.WriteBool(true)
	if err != nil {
		return
	}

	// Primary Bit Mask
	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	// Heightmaps
	err = p.WriteNBT(nil)
	if err != nil {
		return
	}

	// Biomes
	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	// Data
	err = p.WriteVarInt(0)
	if err != nil {
		return
	}

	// Block Entities
	err = p.WriteNBT(nil)
	if err != nil {
		return
	}

	return c.writePacket(p)
}
