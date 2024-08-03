package datatype

func ReadString(b Reader) (string, error) {
	length, err := ReadVarInt(b)

	if err != nil {
		return "", err
	}

	buf := make([]byte, length)
	_, err = b.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func WriteString(b Writer, str string) error {
	err := WriteVarInt(b, len(str))
	if err != nil {
		return err
	}

	_, err = b.Write([]byte(str))
	return err
}