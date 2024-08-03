package datatype

import "encoding/binary"

func ReadInt64(r reader) (value int64, err error) {
	err = binary.Read(r, binary.LittleEndian, &value)
	return
}

func WriteInt64(w writer, value int64) (err error) {
	err = binary.Write(w, binary.LittleEndian, value)
	return
}
