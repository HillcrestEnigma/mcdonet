package connection

import (
	"errors"
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
		packet, err := c.ReadPacket()
		if err != nil {
			return err
		}

		switch packet.id {
		case 0x00:
			err = c.HandleStatusRequest()
			if err != nil {
				return err
			}
		case 0x01:
			err = c.HandleStatusPing(packet)
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
	packet := NewPacket(0x00)

	response := StatusResponse{
		Version: StatusResponseVersion{
			Name:     "1.21",
			Protocol: 767,
		},
	}

	err = packet.WriteJSON(response)
	if err != nil {
		return
	}

	return c.WritePacket(packet)
}

func (c *Connection) HandleStatusPing(request *Packet) error {
	payload, err := request.ReadLong()
	if err != nil {
		return err
	}

	response := NewPacket(0x01)
	err = response.WriteLong(payload)
	if err != nil {
		return err
	}

	return c.WritePacket(response)
}
