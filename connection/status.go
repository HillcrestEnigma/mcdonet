package connection

import (
	"errors"

	"github.com/HillcrestEnigma/mcbuild/packet"
)

type statusResponseVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type statusResponse struct {
	Version statusResponseVersion `json:"version"`
}

func (c *connection) handleServerListPing() error {
	for {
		p, err := c.readPacket(0x00, 0x01)
		if err != nil {
			return err
		}

		switch p.Id {
		case 0x00:
			err = c.handleStatusRequest()
			if err != nil {
				return err
			}
		case 0x01:
			err = c.handleStatusPing(p)
			if err != nil {
				return err
			}
			return nil
		default:
			return errors.New("invalid status packet")
		}
	}
}

func (c *connection) handleStatusRequest() (err error) {
	p := packet.NewPacket(0x00)

	response := statusResponse{
		Version: statusResponseVersion{
			Name:     "1.21",
			Protocol: 767,
		},
	}

	err = p.WriteJSON(response)
	if err != nil {
		return
	}

	return c.writePacket(p)
}

func (c *connection) handleStatusPing(request *packet.Packet) error {
	payload, err := request.ReadInt64()
	if err != nil {
		return err
	}

	response := packet.NewPacket(0x01)
	err = response.WriteInt64(payload)
	if err != nil {
		return err
	}

	return c.writePacket(response)
}
