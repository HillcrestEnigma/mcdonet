package datatype

func ReadVarInt(r reader) (value int, err error) {
	for shift := 0; true; shift += 7 {
		cur, err := r.ReadByte()

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

func WriteVarInt(w writer, value int) (err error) {
	for {
		if (value & 0x80) == 0 {
			w.WriteByte(byte(value))
			return
		}

		w.WriteByte(byte((value & 0x7F) | 0x80))

		value >>= 7
	}
}
