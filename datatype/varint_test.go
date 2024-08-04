package datatype_test

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"

	"github.com/HillcrestEnigma/mcbuild/datatype"
)

var varInts = map[int32][]byte{
	0:           {0x00},
	1:           {0x01},
	2:           {0x02},
	127:         {0x7f},
	128:         {0x80, 0x01},
	255:         {0xff, 0x01},
	25565:       {0xdd, 0xc7, 0x01},
	2097151:     {0xff, 0xff, 0x7f},
	2147483647:  {0xff, 0xff, 0xff, 0xff, 0x07},
	-1:          {0xff, 0xff, 0xff, 0xff, 0x0f},
	-2147483648: {0x80, 0x80, 0x80, 0x80, 0x08},
}

func testReadVarIntMatchesExpected(t *testing.T, expected int32, val []byte) {
	t.Helper()

	t.Run(strconv.Itoa(int(expected)), func(t *testing.T) {
		buf := bytes.NewBuffer(val)

		n, got, err := datatype.ReadVarInt(buf)
		if err != nil {
			t.Fatalf("Error reading varint: %v", err)
		}

		if n != len(val) {
			t.Errorf("ReadVarInt consumed %d bytes, expected %d", n, len(val))
		}
		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}
	})
}

func testWriteVarIntMatchesExpected(t *testing.T, val int32, expected []byte) {
	t.Helper()

	t.Run(strconv.Itoa(int(val)), func(t *testing.T) {
		var buf bytes.Buffer

		err := datatype.WriteVarInt(&buf, val)
		if err != nil {
			t.Errorf("Error writing varint: %v", err)
		}

		got := buf.Bytes()

		if !bytes.Equal(got, expected) {
			t.Errorf("got %d, expected %d", got, expected)
		}
	})
}

func TestReadVarInt(t *testing.T) {
	for val, expected := range varInts {
		testReadVarIntMatchesExpected(t, val, expected)
	}
	testReadVarIntMatchesExpected(t, 1, []byte{0x81, 0x00})
}

func TestWriteVarInt(t *testing.T) {
	for val, expected := range varInts {
		testWriteVarIntMatchesExpected(t, val, expected)
	}
}

func FuzzReadWriteVarInt(f *testing.F) {
	for i := 0; i < 100; i++ {
		f.Add(rand.Int31())
		f.Add(-rand.Int31())
	}

	var buf bytes.Buffer

	f.Fuzz(func(t *testing.T, val int32) {
		err := datatype.WriteVarInt(&buf, val)
		if err != nil {
			t.Fatalf("Error writing varint: %v", err)
		}

		length := buf.Len()

		n, got, err := datatype.ReadVarInt(&buf)
		if err != nil {
			t.Fatalf("Error reading varint: %v", err)
		}

		if n != length {
			t.Fatalf("ReadVarInt consumed %d bytes, expected %d", n, length)
		}
		if got != val {
			t.Errorf("got %d, expected %d", got, val)
		}
	})
}
