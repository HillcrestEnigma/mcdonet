package datatype

import "github.com/google/uuid"

func ReadString(r reader) (string, error) {
	length, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	buf := make([]byte, length)
	_, err = r.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func WriteString(w writer, str string) error {
	err := WriteVarInt(w, len(str))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(str))
	return err
}

func ReadUUID(r reader) (value uuid.UUID, err error) {
	buf := make([]byte, 16)

	_, err = r.Read(buf)
	if err != nil {
		return
	}

	value, err = uuid.FromBytes(buf)
	return
}

func WriteUUID(w writer, value uuid.UUID) (err error) {
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