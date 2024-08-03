package datatype

import "encoding/binary"

func ReadInt64(r Reader) (value int64, err error) {
	err = binary.Read(r, binary.LittleEndian, &value)
	return
}

func WriteInt64(w Writer, value int64) (err error) {
	err = binary.Write(w, binary.LittleEndian, value)
	return
}
