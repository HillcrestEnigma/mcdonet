package datatype

import (
	"github.com/google/uuid"
)

func ReadUUID(r Reader) (value uuid.UUID, err error) {
	buf := make([]byte, 16)

	_, err = r.Read(buf)
	if err != nil {
		return
	}

	value, err = uuid.FromBytes(buf)
	return
}