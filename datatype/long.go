package datatype

import "encoding/binary"

func ReadLong(b Reader) (value int64, err error) {
	err = binary.Read(b, binary.LittleEndian, &value)
	return
}

func WriteLong(b Writer, value int64) (err error) {
	err = binary.Write(b, binary.LittleEndian, value)
	return
}