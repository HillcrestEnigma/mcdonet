package connection

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"slices"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/HillcrestEnigma/mcbuild/packet"
	"github.com/google/uuid"
)

type connectionPlayer struct {
	UUID     uuid.UUID
	Username string
}

type connection struct {
	player *connectionPlayer
	net    net.Conn
	buf    *bufio.ReadWriter
}

func (c *connection) close() {
	c.net.Close()
}

func (c *connection) readPacket(allowedIDs ...int) (p *packet.Packet, err error) {
	length, err := datatype.ReadVarInt(c.buf)
	if err != nil {
		return nil, err
	}

	id, err := datatype.ReadVarInt(c.buf)
	if err != nil {
		return nil, err
	}
	if len(allowedIDs) > 0 && !slices.Contains(allowedIDs, id) {
		return nil, fmt.Errorf("unexpected packet id %d", id)
	}

	byteSlice := make([]byte, length-1)
	_, err = c.buf.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	p = packet.NewPacket(id)
	p.Write(byteSlice)

	return
}

func (c *connection) acceptPacket(acceptableIDs ...int) (p *packet.Packet, err error) {
	for {
		p, err = c.readPacket()
		if err != nil {
			return
		}

		if slices.Contains(acceptableIDs, p.Id) {
			return
		}
	}
}

func (c *connection) writePacket(p *packet.Packet) (err error) {
	var buf bytes.Buffer

	err = datatype.WriteVarInt(&buf, p.Id)
	if err != nil {
		return err
	}

	_, err = buf.Write(p.Data.Bytes())
	if err != nil {
		return err
	}

	err = datatype.WriteVarInt(c.buf, buf.Len())
	if err != nil {
		return err
	}

	_, err = c.buf.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return c.buf.Flush()
}
