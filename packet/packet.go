package packet

import (
	"bytes"
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