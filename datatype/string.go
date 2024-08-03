package datatype

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
