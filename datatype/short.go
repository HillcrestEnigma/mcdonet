package datatype

func ReadUnsignedShort(b Reader) (uint16, error) {
	buf := make([]byte, 2)
	_, err := b.Read(buf)
	if err != nil {
		return 0, err
	}

	return (uint16(buf[0]) << 8) | uint16(buf[1]), nil
}
