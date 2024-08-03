package datatype

func ReadBool(r Reader) (value bool, err error) {
	b, err := r.ReadByte()

	if err != nil {
		return
	}

	value = b != 0
	return
}

func WriteBool(w Writer, value bool) (err error) {
	if value {
		err = w.WriteByte(1)
	} else {
		err = w.WriteByte(0)
	}

	return
}