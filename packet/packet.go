package packet

import (
	"bytes"
	"fmt"
	"io"
	"slices"

	"github.com/HillcrestEnigma/mcdonet/datatype"
)

type Packet struct {
	bytes.Buffer
	Id int32
}

func NewPacket(id int32) *Packet {
	return &Packet{
		Id: id,
	}
}

func ReadPacket(r datatype.Reader, allowedIDs ...int32) (p *Packet, err error) {
	_, length, err := datatype.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	n, id, err := datatype.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	if len(allowedIDs) > 0 && !slices.Contains(allowedIDs, id) {
		return nil, fmt.Errorf("unexpected packet id %d", id)
	}

	byteSlice := make([]byte, length-int32(n))
	_, err = io.ReadFull(r, byteSlice)
	if err != nil {
		return nil, err
	}

	p = NewPacket(id)
	p.Write(byteSlice)

	return
}

func AcceptPacket(r datatype.Reader, acceptableIDs ...int32) (p *Packet, err error) {
	for {
		p, err = ReadPacket(r)
		if err != nil {
			return
		}

		if slices.Contains(acceptableIDs, p.Id) {
			return
		}
	}
}

func WritePacket(w datatype.Writer, p *Packet) (err error) {
	var buf bytes.Buffer

	err = datatype.WriteVarInt(&buf, p.Id)
	if err != nil {
		return err
	}

	_, err = buf.Write(p.Bytes())
	if err != nil {
		return err
	}

	err = datatype.WriteVarInt(w, int32(buf.Len()))
	if err != nil {
		return err
	}

	_, err = w.Write(buf.Bytes())

	return
}
