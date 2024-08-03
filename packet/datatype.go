package packet

import (
	"encoding/json"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/google/uuid"
)

func (p *Packet) ReadVarInt() (int, error) {
	return datatype.ReadVarInt(p)
}

func (p *Packet) WriteVarInt(value int) error {
	return datatype.WriteVarInt(p, value)
}


func (p *Packet) ReadString() (string, error) {
	return datatype.ReadString(p)
}

func (p *Packet) WriteString(str string) error {
	return datatype.WriteString(p, str)
}


func (p *Packet) WriteJSON(obj any) (err error) {
	res, err := json.Marshal(obj)
	if err != nil {
		return
	}

	return p.WriteString(string(res))
}


func (p *Packet) ReadUInt8() (uint8, error) {
	return datatype.ReadInt[uint8](p)
}

func (p *Packet) WriteUInt8(value uint8) error {
	return datatype.WriteInt(p, value)
}


func (p *Packet) ReadInt8() (int8, error) {
	return datatype.ReadInt[int8](p)
}

func (p *Packet) WriteInt8(value int8) error {
	return datatype.WriteInt(p, value)
}


func (p *Packet) ReadUInt16() (uint16, error) {
	return datatype.ReadInt[uint16](p)
}

func (p *Packet) WriteUInt16(value uint16) error {
	return datatype.WriteInt(p, value)
}


func (p *Packet) ReadInt32() (int32, error) {
	return datatype.ReadInt[int32](p)
}

func (p *Packet) WriteInt32(value int32) error {
	return datatype.WriteInt(p, value)
}


func (p *Packet) ReadInt64() (int64, error) {
	return datatype.ReadInt[int64](p)
}

func (p *Packet) WriteInt64(value int64) error {
	return datatype.WriteInt(p, value)
}


func (p *Packet) ReadUUID() (uuid.UUID, error) {
	return datatype.ReadUUID(p)
}

func (p *Packet) WriteUUID(value uuid.UUID) error {
	return datatype.WriteUUID(p, value)
}


func (p *Packet) ReadBool() (bool, error) {
	return datatype.ReadBool(p)
}

func (p *Packet) WriteBool(value bool) error {
	return datatype.WriteBool(p, value)
}