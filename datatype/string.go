package datatype

import (
	"io"

	"github.com/google/uuid"
)

func ReadString(r Reader) (string, error) {
	_, length, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	return readRawString(r, length)
}

func WriteString(w Writer, str string) error {
	err := WriteVarInt(w, int32(len(str)))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(str))
	return err
}

func ReadUUID(r Reader) (val uuid.UUID, err error) {
	buf := make([]byte, 16)

	_, err = r.Read(buf)
	if err != nil {
		return
	}

	val, err = uuid.FromBytes(buf)
	return
}

func WriteUUID(w Writer, val uuid.UUID) (err error) {
	bin, err := val.MarshalBinary()
	if err != nil {
		return
	}

	_, err = w.Write(bin)
	if err != nil {
		return
	}

	return
}

func readRawString(r Reader, length int32) (string, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func writeRawString(w Writer, str string) error {
	_, err := w.Write([]byte(str))
	return err
}
