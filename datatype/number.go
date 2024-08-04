package datatype

import "encoding/binary"

type number interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~float32 | ~float64
}

func ReadNumber[V number](r Reader) (val V, err error) {
	err = binary.Read(r, binary.BigEndian, &val)
	return
}

func WriteNumber[V number](w Writer, val V) (err error) {
	err = binary.Write(w, binary.BigEndian, val)
	return
}
