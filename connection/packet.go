package connection

import (
	"bytes"
	"encoding/json"

	"github.com/HillcrestEnigma/mcbuild/datatype"
)

type Packet struct {
	id   int
	data *bytes.Buffer
}

func NewPacket(id int) *Packet {
	return &Packet{
		id:   id,
		data: bytes.NewBuffer([]byte{}),
	}
}

func (p *Packet) Read(b []byte) (n int, err error) {
	return p.data.Read(b)
}

func (p *Packet) ReadByte() (byte, error) {
	return p.data.ReadByte()
}

func (p *Packet) Write(b []byte) (n int, err error) {
	return p.data.Write(b)
}

func (p *Packet) WriteByte(b byte) error {
	return p.data.WriteByte(b)
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

func (p *Packet) ReadUnsignedShort() (uint16, error) {
	return datatype.ReadUnsignedShort(p)
}

func (p *Packet) ReadLong() (int64, error) {
	return datatype.ReadLong(p)
}

func (p *Packet) WriteLong(value int64) error {
	return datatype.WriteLong(p, value)
}