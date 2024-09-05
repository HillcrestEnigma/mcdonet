package connection

import (
	"bufio"
	"net"

	"github.com/HillcrestEnigma/mcdonet/packet"
	"github.com/google/uuid"
)

type connectionPlayer struct {
	UUID     uuid.UUID
	Username string
}

type Connection struct {
	player *connectionPlayer
	net    net.Conn
	buf    *bufio.ReadWriter
}

func NewConnection(c net.Conn) *Connection {
	return &Connection{
		player: nil,
		net:    c,
		buf:    bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)),
	}
}

func (c *Connection) close() {
	c.net.Close()
}

func (c *Connection) readPacket(allowedIDs ...int32) (p *packet.Packet, err error) {
	return packet.ReadPacket(c.buf, allowedIDs...)
}

func (c *Connection) acceptPacket(acceptableIDs ...int32) (p *packet.Packet, err error) {
	return packet.AcceptPacket(c.buf, acceptableIDs...)
}

func (c *Connection) writePacket(p *packet.Packet) (err error) {
	err = packet.WritePacket(c.buf, p)
	if err != nil {
		return
	}

	return c.buf.Flush()
}
