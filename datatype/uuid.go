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

func WriteUUID(w Writer, value uuid.UUID) (err error) {
	bin, err := value.MarshalBinary()
	if err != nil {
		return
	}

	_, err = w.Write(bin)
	if err != nil {
		return
	}

	return
}