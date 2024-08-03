package packet

import (
	"bytes"
	"encoding/json"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/google/uuid"
)

type Packet struct {
	Id   int
	Data *bytes.Buffer
}

func NewPacket(id int) *Packet {
	return &Packet{
		Id:   id,
		Data: bytes.NewBuffer([]byte{}),
	}
}

func (p *Packet) Read(b []byte) (n int, err error) {
	return p.Data.Read(b)
}

func (p *Packet) ReadByte() (byte, error) {
	return p.Data.ReadByte()
}

func (p *Packet) Write(b []byte) (n int, err error) {
	return p.Data.Write(b)
}

func (p *Packet) WriteByte(b byte) error {
	return p.Data.WriteByte(b)
}

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

func (p *Packet) ReadUInt16() (uint16, error) {
	return datatype.ReadUInt16(p)
}

func (p *Packet) ReadInt64() (int64, error) {
	return datatype.ReadInt64(p)
}

func (p *Packet) WriteInt64(value int64) error {
	return datatype.WriteInt64(p, value)
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