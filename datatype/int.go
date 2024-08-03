package datatype

import "encoding/binary"

func ReadUInt16(r reader) (uint16, error) {
	buf := make([]byte, 2)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}

	return (uint16(buf[0]) << 8) | uint16(buf[1]), nil
}

func ReadInt64(r reader) (value int64, err error) {
	err = binary.Read(r, binary.LittleEndian, &value)
	return
}

func WriteInt64(w writer, value int64) (err error) {
	err = binary.Write(w, binary.LittleEndian, value)
	return
}