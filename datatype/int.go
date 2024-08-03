package datatype

import "encoding/binary"

func ReadUInt8(r reader) (val uint8, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteUInt8(w writer, val uint8) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}

func ReadInt8(r reader) (val int8, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteInt8(w writer, val int8) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}

func ReadUInt16(r reader) (val uint16, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteUInt16(w writer, val uint16) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}

func ReadInt32(r reader) (val int32, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteInt32(w writer, val int32) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}

func ReadInt64(r reader) (val int64, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func WriteInt64(w writer, val int64) (err error) {
	err = binary.Write(w, binary.LittleEndian, val)
	return
}