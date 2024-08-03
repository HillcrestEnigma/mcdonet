package datatype

import "encoding/binary"

type number interface {
	uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64
}

func ReadInt[V number](r reader) (val V, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteInt[V number](w writer, val V) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}