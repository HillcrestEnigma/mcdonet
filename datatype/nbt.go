package datatype

import (
	"fmt"
	"log"
)

type nbtList []any
type nbtCompound map[string]any

type NBT struct {
	Name     string
	Compound *nbtCompound
}

func ReadNBT(r reader) (val *NBT, err error) {
	typeID, name, err := readNBTHeader(r)
	if err != nil {
		return
	}
	if typeID != 10 {
		err = fmt.Errorf("expected NBT compound type ID 10, got %d", typeID)
		return
	}

	compound, err := readNBTPayload(r, typeID)
	if err != nil {
		return
	}

	val = &NBT{Name: name, Compound: compound.(*nbtCompound)}

	return
}

func WriteNBT(w writer, val *NBT) (err error) {
	err = writeNBTHeader(w, 10, val.Name)
	if err != nil {
		return
	}

	err = writeNBTPayload(w, val.Compound)
	return
}

func readNBTHeader(r reader) (typeID int8, name string, err error) {
	typeID, err = ReadNumber[int8](r)
	if err != nil {
		return
	}
	if typeID == 0 {
		return
	}

	name, err = readNBTString(r)
	if err != nil {
		return
	}
	log.Println("readNBTHeader", typeID, name)

	return
}

func writeNBTHeader(w writer, typeID int8, name string) (err error) {
	log.Println("writeNBTHeader", typeID, name)
	err = WriteNumber(w, typeID)
	if err != nil {
		return
	}
	if typeID == 0 {
		return
	}

	err = writeNBTString(w, name)
	if err != nil {
		return
	}

	return
}

func readNBTPayload(r reader, typeID int8) (val any, err error) {
	log.Println("readNBTPayload", typeID)
	switch typeID {
	case 1:
		val, err = ReadNumber[int8](r)
	case 2:
		val, err = ReadNumber[int16](r)
	case 3:
		val, err = ReadNumber[int32](r)
	case 4:
		val, err = ReadNumber[int64](r)
	case 5:
		val, err = ReadNumber[float32](r)
	case 6:
		val, err = ReadNumber[float64](r)
	case 7:
		length, err := ReadNumber[int32](r)
		if err != nil {
			return nil, err
		}

		array := make([]byte, length)
		_, err = r.Read(array)
		if err != nil {
			return nil, err
		}

		val = &array
	case 8:
		val, err = readNBTString(r)
	case 9:
		listTypeID, err := ReadNumber[int8](r)
		if err != nil {
			return nil, err
		}

		length, err := ReadNumber[int32](r)
		if err != nil {
			return nil, err
		}

		list := make(nbtList, length)
		for i := int32(0); i < length; i++ {
			item, err := readNBTPayload(r, listTypeID)
			if err != nil {
				return nil, err
			}

			list[i] = item
		}

		val = &list
	case 10:
		compound := make(nbtCompound)
		for {
			typeID, name, err := readNBTHeader(r)
			if err != nil {
				return nil, err
			}
			if typeID == 0 {
				break
			}

			payload, err := readNBTPayload(r, typeID)
			if err != nil {
				return nil, err
			}

			compound[name] = payload
		}

		val = &compound
	case 11:
		length, err := ReadNumber[int32](r)
		if err != nil {
			return nil, err
		}

		array := make([]int32, length)
		for i := int32(0); i < length; i++ {
			item, err := ReadNumber[int32](r)
			if err != nil {
				return nil, err
			}

			array[i] = item
		}

		val = &array
	case 12:
		length, err := ReadNumber[int32](r)
		if err != nil {
			return nil, err
		}

		array := make([]int64, length)
		for i := int32(0); i < length; i++ {
			item, err := ReadNumber[int64](r)
			if err != nil {
				return nil, err
			}

			array[i] = item
		}

		val = &array
	default:
		err = fmt.Errorf("unknown NBT type ID: %d", typeID)
	}

	return
}

func writeNBTPayload(w writer, val any) (err error) {
	log.Println("writeNBTPayload")
	typeID, err := getTypeIDFromType(val)
	if err != nil {
		return
	}

	switch typeID {
	case 1:
		err = WriteNumber(w, val.(int8))
	case 2:
		err = WriteNumber(w, val.(int16))
	case 3:
		err = WriteNumber(w, val.(int32))
	case 4:
		err = WriteNumber(w, val.(int64))
	case 5:
		err = WriteNumber(w, val.(float32))
	case 6:
		err = WriteNumber(w, val.(float64))
	case 7:
		array := *val.(*[]byte)

		err = WriteNumber(w, int32(len(array)))
		if err != nil {
			return
		}

		_, err = w.Write(array)
	case 8:
		err = writeNBTString(w, val.(string))
	case 9:
		list := val.(*nbtList)

		var listTypeID int8 = 0
		if len(*list) > 0 {
			listTypeID, err = getTypeIDFromType((*list)[0])
			if err != nil {
				return
			}
		}

		err = WriteNumber(w, listTypeID)
		if err != nil {
			return
		}

		err = WriteNumber(w, int32(len(*list)))
		if err != nil {
			return
		}

		for _, item := range *list {
			err = writeNBTPayload(w, item)
			if err != nil {
				return
			}
		}
	case 10:
		compound := val.(*nbtCompound)

		for name, item := range *compound {
			log.Println("compound item")
			itemTypeID, err := getTypeIDFromType(item)
			if err != nil {
				return err
			}

			err = writeNBTHeader(w, itemTypeID, name)
			if err != nil {
				return err
			}

			err = writeNBTPayload(w, item)
			if err != nil {
				return err
			}
		}

		err = w.WriteByte(0x00)
		log.Println("write compound end")
	case 11:
		array := val.([]int32)

		err = WriteNumber(w, int32(len(array)))
		if err != nil {
			return
		}

		for _, item := range array {
			err = WriteNumber(w, item)
			if err != nil {
				return
			}
		}
	case 12:
		array := val.([]int64)

		err = WriteNumber(w, int32(len(array)))
		if err != nil {
			return
		}

		for _, item := range array {
			err = WriteNumber(w, item)
			if err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unknown NBT type ID: %d", typeID)
	}

	return
}

func readNBTString(r reader) (val string, err error) {
	length, err := ReadNumber[int16](r)
	if err != nil {
		return
	}

	return readRawString(r, int(length))
}

func writeNBTString(w writer, val string) (err error) {
	log.Println("writeNBTString", val)
	err = WriteNumber(w, int16(len(val)))
	if err != nil {
		return
	}

	return writeRawString(w, val)
}

func getTypeIDFromType(val any) (typeID int8, err error) {
	switch val.(type) {
	case int8:
		typeID = 1
	case int16:
		typeID = 2
	case int32:
		typeID = 3
	case int64:
		typeID = 4
	case float32:
		typeID = 5
	case float64:
		typeID = 6
	case *[]byte:
		typeID = 7
	case string:
		typeID = 8
	case *nbtList:
		typeID = 9
	case *nbtCompound:
		typeID = 10
	case *[]int32:
		typeID = 11
	case *[]int64:
		typeID = 12
	default:
		typeID = -1
		err = fmt.Errorf("unknown NBT type: %T", val)
	}

	log.Println("getTypeIDFromType", val, typeID)

	return
}
