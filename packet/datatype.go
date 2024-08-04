package packet

import (
	"encoding/json"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/google/uuid"
)

func (p *Packet) ReadVarInt() (val int32, err error) {
	_, val, err = datatype.ReadVarInt(p)
	return
}

func (p *Packet) WriteVarInt(val int32) error {
	return datatype.WriteVarInt(p, val)
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
	return datatype.ReadNumber[uint8](p)
}

func (p *Packet) WriteUInt8(val uint8) error {
	return datatype.WriteNumber(p, val)
}

func (p *Packet) ReadInt8() (int8, error) {
	return datatype.ReadNumber[int8](p)
}

func (p *Packet) WriteInt8(val int8) error {
	return datatype.WriteNumber(p, val)
}

func (p *Packet) ReadUInt16() (uint16, error) {
	return datatype.ReadNumber[uint16](p)
}

func (p *Packet) WriteUInt16(val uint16) error {
	return datatype.WriteNumber(p, val)
}

func (p *Packet) ReadInt32() (int32, error) {
	return datatype.ReadNumber[int32](p)
}

func (p *Packet) WriteInt32(val int32) error {
	return datatype.WriteNumber(p, val)
}

func (p *Packet) ReadInt64() (int64, error) {
	return datatype.ReadNumber[int64](p)
}

func (p *Packet) WriteInt64(val int64) error {
	return datatype.WriteNumber(p, val)
}

func (p *Packet) ReadUUID() (uuid.UUID, error) {
	return datatype.ReadUUID(p)
}

func (p *Packet) WriteUUID(val uuid.UUID) error {
	return datatype.WriteUUID(p, val)
}

func (p *Packet) ReadBool() (bool, error) {
	return datatype.ReadBool(p)
}

func (p *Packet) WriteBool(val bool) error {
	return datatype.WriteBool(p, val)
}

func (p *Packet) ReadNBT() (*datatype.NBT, error) {
	return datatype.ReadNetworkNBT(p)
}

func (p *Packet) WriteNBT(val *datatype.NBT) error {
	return datatype.WriteNetworkNBT(p, val)
}