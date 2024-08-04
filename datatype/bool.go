package datatype

func ReadBool(r Reader) (val bool, err error) {
	b, err := r.ReadByte()

	if err != nil {
		return
	}

	val = b != 0
	return
}

func WriteBool(w Writer, val bool) (err error) {
	if val {
		err = w.WriteByte(1)
	} else {
		err = w.WriteByte(0)
	}

	return
}
