package datatype

func ReadVarInt(r Reader) (n int, val int32, err error) {
	for n = 1; true; n++ {
		cur, err := r.ReadByte()
		if err != nil {
			return 0, 0, err
		}

		val |= (int32(cur) & 0x7f) << ((n - 1) * 7)

		if (cur & 0x80) == 0 {
			break
		}
	}

	return
}

func WriteVarInt(w Writer, val int32) (err error) {
	uval := uint32(val)

	for {
		if (uval & ^uint32(0x7f)) == 0 {
			err = w.WriteByte(byte(uval))
			return
		}

		err = w.WriteByte(byte((uval & 0x7f) | 0x80))
		if err != nil {
			return
		}

		uval >>= 7
	}
}
