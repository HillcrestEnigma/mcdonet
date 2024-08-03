package connection

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/HillcrestEnigma/mcbuild/packet"
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

func HandleConnection(netConn net.Conn) {
	c := &Connection{
		state: 0,
		net:   netConn,
		buf:   bufio.NewReadWriter(bufio.NewReader(netConn), bufio.NewWriter(netConn)),
	}

	defer c.Close()

	isLegacy, err := c.HandleLegacyServerListPing()
	if err != nil {
		return
	}
	if isLegacy {
		return
	}

	handshake, err := c.ReadHandshake()
	if err != nil {
		return
	}

	switch handshake.nextState {
	case 1:
		err = c.HandleServerListPing()
	case 2:
		err = c.HandleLogin()
	}

	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed by client")
		} else {
			log.Println(err)
		}
	}
}

func (c *Connection) Close() {
	c.net.Close()
}

func (c *Connection) ReadPacket() (p *packet.Packet, err error) {
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

	p = packet.NewPacket(id)
	p.Write(byteSlice)

	return
}

func (c *Connection) WritePacket(p *packet.Packet) (err error) {
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


