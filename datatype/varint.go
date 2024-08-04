package datatype

func ReadVarInt(r reader) (val int, err error) {
	for shift := 0; true; shift += 7 {
		cur, err := r.ReadByte()

		if err != nil {
			return 0, err
		}

		val |= (int(cur) & 0x7F) << shift

		if (cur & 0x80) == 0 {
			break
		}
	}

	return
}

func WriteVarInt(w writer, val int) (err error) {
	for {
		if (val & 0x80) == 0 {
			w.WriteByte(byte(val))
			return
		}

		w.WriteByte(byte((val & 0x7F) | 0x80))

		val >>= 7
	}
}
