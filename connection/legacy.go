package connection

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

func (c *connection) handleLegacyServerListPing() (bool, error) {
	sig, err := c.buf.Peek(3)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(sig, []byte{0xFE, 0x01, 0xFA}) {
		return false, nil
	}

	info := []string{
		"§1",
		"47",
		"Greater than 1.7 pls",
		"MCBuild Server",
		"0",
		"20",
	}

	response := make([]byte, 0)
	for idx, str := range info {
		encoded := utf16.Encode([]rune(str))
		for _, v := range encoded {
			response = binary.BigEndian.AppendUint16(response, v)
		}

		if idx != len(info)-1 {
			response = append(response, 0x00, 0x00)
		}
	}

	c.buf.WriteByte(0xFF)
	binary.Write(c.buf, binary.BigEndian, int16(len(response)/2))
	c.buf.Write(response)

	return true, c.buf.Flush()
}
