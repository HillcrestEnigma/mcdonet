package connection

import (
	"errors"

	"github.com/HillcrestEnigma/mcbuild/packet"
)

type StatusResponseVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type StatusResponse struct {
	Version StatusResponseVersion `json:"version"`
}

func (c *Connection) HandleServerListPing() error {
	for {
		p, err := c.ReadPacket(0x00, 0x01)
		if err != nil {
			return err
		}

		switch p.Id {
		case 0x00:
			err = c.HandleStatusRequest()
			if err != nil {
				return err
			}
		case 0x01:
			err = c.HandleStatusPing(p)
			if err != nil {
				return err
			}
			return nil
		default:
			return errors.New("invalid status packet")
		}
	}
}

func (c *Connection) HandleStatusRequest() (err error) {
	p := packet.NewPacket(0x00)

	response := StatusResponse{
		Version: StatusResponseVersion{
			Name:     "1.21",
			Protocol: 767,
		},
	}

	err = p.WriteJSON(response)
	if err != nil {
		return
	}

	return c.WritePacket(p)
}

func (c *Connection) HandleStatusPing(request *packet.Packet) error {
	payload, err := request.ReadInt64()
	if err != nil {
		return err
	}

	response := packet.NewPacket(0x01)
	err = response.WriteInt64(payload)
	if err != nil {
		return err
	}

	return c.WritePacket(response)
}
