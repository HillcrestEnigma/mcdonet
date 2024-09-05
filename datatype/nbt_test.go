package datatype_test

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/HillcrestEnigma/mcdonet/datatype"
)

func openFile(t testing.TB, filename string) (r datatype.Reader, file *os.File) {
	t.Helper()

	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}

	r = bufio.NewReader(file)
	return
}

func openNBT(t testing.TB, filename string) (nbt *datatype.NBT) {
	t.Helper()

	r, f := openFile(t, filename)
	defer f.Close()

	nbt, err := datatype.ReadNBT(r)
	if err != nil {
		t.Fatalf("Error reading NBT: %v", err)
	}

	return
}

func assertReadNBTMatchesExpected(t testing.TB, filename string, expected *datatype.NBT) {
	t.Helper()

	got := openNBT(t, filename)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got NBT %+v, expected %+v", got, expected)
	}
}

func assertWriteNBTMatchesReadNBT(t testing.TB, filename string) {
	t.Helper()

	original := openNBT(t, filename)

	var written bytes.Buffer
	err := datatype.WriteNBT(&written, original)
	if err != nil {
		t.Fatalf("Error writing NBT: %v", err)
	}

	reread, err := datatype.ReadNBT(&written)
	if err != nil {
		t.Fatalf("Error reading NBT: %v", err)
	}

	if !reflect.DeepEqual(reread, original) {
		t.Errorf("Got NBT %+v, expected %+v", reread, original)
	}
}

func assertWriteNetworkNBTMatchesReadNetworkNBT(t testing.TB, filename string) {
	t.Helper()

	original := openNBT(t, filename)

	var written bytes.Buffer
	err := datatype.WriteNetworkNBT(&written, original)
	if err != nil {
		t.Fatalf("Error writing Network NBT: %v", err)
	}

	reread, err := datatype.ReadNetworkNBT(&written)
	if err != nil {
		t.Fatalf("Error reading Network NBT: %v", err)
	}

	if !reflect.DeepEqual(reread.Compound, original.Compound) {
		t.Errorf("Got Network NBT %+v, expected %+v", reread.Compound, original.Compound)
	}
}

func TestReadNBT(t *testing.T) {
	t.Run("helloworld.nbt", func(t *testing.T) {
		expected := &datatype.NBT{
			Name: "hello world",
			Compound: datatype.NBTCompound{
				"name": "Bananrama",
			},
		}

		assertReadNBTMatchesExpected(t, "helloworld.nbt", expected)
	})

	t.Run("bigtest.nbt", func(t *testing.T) {
		expected := &datatype.NBT{
			Name: "Level",
			Compound: datatype.NBTCompound{
				"nested compound test": datatype.NBTCompound{
					"egg": datatype.NBTCompound{
						"name":  "Eggbert",
						"value": float32(0.5),
					},
					"ham": datatype.NBTCompound{
						"name":  "Hampus",
						"value": float32(0.75),
					},
				},
				"intTest":    int32(2147483647),
				"byteTest":   int8(127),
				"stringTest": "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!",
				"listTest (long)": datatype.NBTList{
					int64(11),
					int64(12),
					int64(13),
					int64(14),
					int64(15),
				},
				"doubleTest": float64(0.49312871321823148),
				"floatTest":  float32(0.49823147058486938),
				"longTest":   int64(9223372036854775807),
				"listTest (compound)": datatype.NBTList{
					datatype.NBTCompound{
						"created-on": int64(1264099775885),
						"name":       "Compound tag #0",
					},
					datatype.NBTCompound{
						"created-on": int64(1264099775885),
						"name":       "Compound tag #1",
					},
				},
				"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))": []byte{
					0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48, 0, 62, 34, 16, 8, 10, 22, 44, 76, 18, 70, 32, 4, 86, 78, 80, 92, 14, 46, 88, 40, 2, 74, 56, 48, 50, 62, 84, 16, 58, 10, 72, 44, 26, 18, 20, 32, 54, 86, 28, 80, 42, 14, 96, 88, 90, 2, 24, 56, 98, 50, 12, 84, 66, 58, 60, 72, 94, 26, 68, 20, 82, 54, 36, 28, 30, 42, 64, 96, 38, 90, 52, 24, 6, 98, 0, 12, 34, 66, 8, 60, 22, 94, 76, 68, 70, 82, 4, 36, 78, 30, 92, 64, 46, 38, 40, 52, 74, 6, 48,
				},
				"shortTest": int16(32767),
			},
		}

		assertReadNBTMatchesExpected(t, "bigtest.nbt", expected)
	})
}

func TestWriteNBT(t *testing.T) {
	t.Run("helloworld.nbt", func(t *testing.T) {
		assertWriteNBTMatchesReadNBT(t, "helloworld.nbt")
	})

	t.Run("bigtest.nbt", func(t *testing.T) {
		assertWriteNBTMatchesReadNBT(t, "bigtest.nbt")
	})
}

func TestNetworkNBT(t *testing.T) {
	t.Run("helloworld.nbt", func(t *testing.T) {
		assertWriteNetworkNBTMatchesReadNetworkNBT(t, "helloworld.nbt")
	})

	t.Run("bigtest.nbt", func(t *testing.T) {
		assertWriteNetworkNBTMatchesReadNetworkNBT(t, "bigtest.nbt")
	})
}
