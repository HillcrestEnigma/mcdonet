package connection

import (
	"bufio"
	"net"

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

func (c *connection) readPacket(allowedIDs ...int32) (p *packet.Packet, err error) {
	return packet.ReadPacket(c.buf, allowedIDs...)
}

func (c *connection) acceptPacket(acceptableIDs ...int32) (p *packet.Packet, err error) {
	return packet.AcceptPacket(c.buf, acceptableIDs...)
}

func (c *connection) writePacket(p *packet.Packet) (err error) {
	err = packet.WritePacket(c.buf, p)
	if err != nil {
		return
	}

	return c.buf.Flush()
}
