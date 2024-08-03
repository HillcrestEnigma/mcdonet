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
