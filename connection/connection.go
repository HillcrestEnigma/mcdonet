package connection

import (
	"bufio"
	"bytes"
	"net"

	"github.com/HillcrestEnigma/mcbuild/datatype"
)

type ConnectionState int

const (
	ConnectionStateHandshake = iota
	ConnectionStateStatus
	ConnectionStateLogin
	ConnectionStatePlay
)

type Connection struct {
	state ConnectionState
	net   net.Conn
	buf   *bufio.ReadWriter
}

func NewConnection(c net.Conn) *Connection {
	return &Connection{
		state: 0,
		net:   c,
		buf:   bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)),
	}
}

func (c *Connection) Close() {
	c.net.Close()
}

func (c *Connection) ReadPacket() (packet *Packet, err error) {
	length, err := datatype.ReadVarInt(c.buf)
	if err != nil {
		return nil, err
	}

	id, err := datatype.ReadVarInt(c.buf)
	if err != nil {
		return nil, err
	}

	byteSlice := make([]byte, length-1)
	_, err = c.buf.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	data := bytes.NewBuffer(byteSlice)
	data.Write(byteSlice)

	packet = &Packet{
		id:   id,
		data: data,
	}

	return
}

func (c *Connection) WritePacket(p *Packet) (err error) {
	var buf bytes.Buffer

	err = datatype.WriteVarInt(&buf, p.id)
	if err != nil {
		return err
	}

	_, err = buf.Write(p.data.Bytes())
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
