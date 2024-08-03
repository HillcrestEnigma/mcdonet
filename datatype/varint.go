package datatype

func ReadVarInt(b Reader) (value int, err error) {
	for shift := 0; true; shift += 7 {
		cur, err := b.ReadByte()

		if err != nil {
			return 0, err
		}

		value |= (int(cur) & 0x7F) << shift

		if (cur & 0x80) == 0 {
			break
		}
	}

	return
}

func WriteVarInt(b Writer, value int) (err error) {
	for {
		if ((value & 0x80) == 0) {
			b.WriteByte(byte(value))
			return;
		}

		b.WriteByte(byte((value & 0x7F) | 0x80))

		value >>= 7
	}
}