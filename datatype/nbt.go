package datatype

import (
	"fmt"
	"strings"
)

type NBTList []any
type NBTCompound map[string]any

type NBT struct {
	Name     string
	Compound NBTCompound
}

// TODO: consider changing function signatures to accept a pointer to NBT
func ReadNBT(r Reader) (val *NBT, err error) {
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

	val = &NBT{Name: name, Compound: compound.(NBTCompound)}
	return
}

func ReadNetworkNBT(r Reader) (val *NBT, err error) {
	typeID, err := ReadNumber[int8](r)
	if err != nil {
		return
	}

	compound, err := readNBTPayload(r, typeID)
	if err != nil {
		return
	}

	val = &NBT{Name: "", Compound: compound.(NBTCompound)}
	return
}

func WriteNBT(w Writer, val *NBT) (err error) {
	err = writeNBTHeader(w, 10, val.Name)
	if err != nil {
		return
	}

	err = writeNBTPayload(w, val.Compound)
	return
}

func WriteNetworkNBT(w Writer, val *NBT) (err error) {
	err = WriteNumber(w, int8(10))
	if err != nil {
		return
	}

	err = writeNBTPayload(w, val.Compound)
	return
}

func readNBTHeader(r Reader) (typeID int8, name string, err error) {
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

	return
}

func writeNBTHeader(w Writer, typeID int8, name string) (err error) {
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

func readNBTPayload(r Reader, typeID int8) (val any, err error) {
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

		val = array
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

		list := make(NBTList, length)
		for i := int32(0); i < length; i++ {
			item, err := readNBTPayload(r, listTypeID)
			if err != nil {
				return nil, err
			}

			list[i] = item
		}

		val = list
	case 10:
		compound := make(NBTCompound)
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

		val = compound
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

		val = array
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

		val = array
	default:
		err = fmt.Errorf("unknown NBT type ID: %d", typeID)
	}

	return
}

func writeNBTPayload(w Writer, val any) (err error) {
	typeID, err := typeIDFromType(val)
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
		array := val.([]byte)

		err = WriteNumber(w, int32(len(array)))
		if err != nil {
			return
		}

		_, err = w.Write(array)
	case 8:
		err = writeNBTString(w, val.(string))
	case 9:
		list := val.(NBTList)

		var listTypeID int8 = 0
		if len(list) > 0 {
			listTypeID, err = typeIDFromType((list)[0])
			if err != nil {
				return
			}
		}

		err = WriteNumber(w, listTypeID)
		if err != nil {
			return
		}

		err = WriteNumber(w, int32(len(list)))
		if err != nil {
			return
		}

		for _, item := range list {
			err = writeNBTPayload(w, item)
			if err != nil {
				return
			}
		}
	case 10:
		compound := val.(NBTCompound)
		for name, item := range compound {
			itemTypeID, err := typeIDFromType(item)
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

func readNBTString(r Reader) (val string, err error) {
	length, err := ReadNumber[int16](r)
	if err != nil {
		return
	}

	return readRawString(r, int32(length))
}

func writeNBTString(w Writer, val string) (err error) {
	err = WriteNumber(w, int16(len(val)))
	if err != nil {
		return
	}

	return writeRawString(w, val)
}

func typeIDFromType(val any) (typeID int8, err error) {
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
	case []byte:
		typeID = 7
	case string:
		typeID = 8
	case NBTList:
		typeID = 9
	case NBTCompound:
		typeID = 10
	case []int32:
		typeID = 11
	case []int64:
		typeID = 12
	default:
		typeID = -1
		err = fmt.Errorf("unknown NBT type: %T", val)
	}

	return
}

func (nbt *NBT) String() (str string) {
	str, _ = nbtStringWithIndent(nbt.Name, nbt.Compound, 0)

	return
}

func nbtStringWithIndent(name string, val any, indentLevel int) (str string, err error) {
	indent := func(level int) string {
		return strings.Repeat("  ", level)
	}

	pluralize := func(count int, singular, plural string) string {
		if count == 1 {
			return singular
		} else {
			return plural
		}
	}

	fmtTagHeader := func(typeName string, name string) string {
		if name == "" {
			name = "None"
		} else {
			name = fmt.Sprintf("'%s'", name)
		}

		return fmt.Sprintf("TAG_%s(%s): ", typeName, name)
	}

	typeID, err := typeIDFromType(val)

	str = strings.Repeat("  ", indentLevel)

	switch typeID {
	case 1:
		str += fmtTagHeader("Byte", name) + fmt.Sprintf("%d", val.(int8))
	case 2:
		str += fmtTagHeader("Short", name) + fmt.Sprintf("%d", val.(int16))
	case 3:
		str += fmtTagHeader("Int", name) + fmt.Sprintf("%d", val.(int32))
	case 4:
		str += fmtTagHeader("Long", name) + fmt.Sprintf("%dL", val.(int64))
	case 5:
		str += fmtTagHeader("Float", name) + fmt.Sprintf("%f", val.(float32))
	case 6:
		str += fmtTagHeader("Double", name) + fmt.Sprintf("%f", val.(float64))
	case 7:
		length := len(val.([]byte))
		str += fmtTagHeader("Byte_Array", name)
		str += fmt.Sprintf("[%d %s]", length, pluralize(length, "byte", "bytes"))
	case 8:
		str += fmtTagHeader("String", name) + fmt.Sprintf("'%s'", val.(string))
	case 9:
		list := val.(NBTList)

		str += fmtTagHeader("List", name) + fmt.Sprintf("%d ", len(list))
		str += pluralize(len(list), "entry", "entries") + "\n" + indent(indentLevel) + "{\n"
		for _, item := range list {
			itemStr, err := nbtStringWithIndent("", item, indentLevel+1)
			if err != nil {
				return "", err
			}

			str += itemStr + "\n"
		}

		str += indent(indentLevel) + "}"
	case 10:
		compound := val.(NBTCompound)

		str += fmtTagHeader("Compound", name) + fmt.Sprintf("%d ", len(compound))
		str += pluralize(len(compound), "entry", "entries") + "\n" + indent(indentLevel) + "{\n"
		for name, item := range compound {
			itemStr, err := nbtStringWithIndent(name, item, indentLevel+1)
			if err != nil {
				return "", err
			}

			str += itemStr + "\n"
		}

		str += indent(indentLevel) + "}"
	case 11:
		array := val.([]int32)
		str += fmtTagHeader("Int_Array", name) + fmt.Sprintf("%v", array)
	case 12:
		array := val.([]int64)
		str += fmtTagHeader("Long_Array", name) + fmt.Sprintf("%v", array)
	default:
		str += fmtTagHeader("Invalid", name) + fmt.Sprintf("%v", val)

		err = fmt.Errorf("unknown NBT type ID: %d", typeID)
	}

	return
}
