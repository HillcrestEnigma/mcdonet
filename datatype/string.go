package datatype

import "github.com/google/uuid"

func ReadString(r reader) (string, error) {
	length, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	return readRawString(r, length)
}

func WriteString(w writer, str string) error {
	err := WriteVarInt(w, len(str))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(str))
	return err
}

func ReadUUID(r reader) (val uuid.UUID, err error) {
	buf := make([]byte, 16)

	_, err = r.Read(buf)
	if err != nil {
		return
	}

	val, err = uuid.FromBytes(buf)
	return
}

func WriteUUID(w writer, val uuid.UUID) (err error) {
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

func readRawString(r reader, length int) (string, error) {
	buf := make([]byte, length)
	_, err := r.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func writeRawString(w writer, str string) error {
	_, err := w.Write([]byte(str))
	return err
}
