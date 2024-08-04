package datatype

func ReadBool(r reader) (val bool, err error) {
	b, err := r.ReadByte()

	if err != nil {
		return
	}

	val = b != 0
	return
}

func WriteBool(w writer, val bool) (err error) {
	if val {
		err = w.WriteByte(1)
	} else {
		err = w.WriteByte(0)
	}

	return
}
