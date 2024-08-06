package chunk

import (
	"github.com/HillcrestEnigma/mcbuild/datatype"
)

// data is indexed [y][z][x]
type palettedContainer struct {
	size    uint8 // size in one dimension
	palette map[int32]uint16
	data    []int32
}

func newPalettedContainer(size uint8, init int32) (p *palettedContainer) {
	dataArraySize := uint16(size) * uint16(size) * uint16(size)
	p = &palettedContainer{
		size:    size,
		palette: make(map[int32]uint16),
		data:    make([]int32, dataArraySize),
	}
	for i := range p.data {
		p.data[i] = init
	}
	p.palette[init] = uint16(dataArraySize)
	return
}

func (p *palettedContainer) get(x, y, z uint8) int32 {
	return p.data[p.getDataIndex(x, y, z)]
}

func (p *palettedContainer) set(x, y, z uint8, value int32) {
	index := p.getDataIndex(x, y, z)

	oldValue := p.data[index]
	if oldValue == value {
		return
	}

	p.palette[oldValue]--
	if p.palette[oldValue] == 0 {
		delete(p.palette, oldValue)
	}

	if _, ok := p.palette[value]; !ok {
		p.palette[value]++
	} else {
		p.palette[value] = 1
	}

	p.data[index] = value
}

func (p *palettedContainer) getDataIndex(x, y, z uint8) (index uint32) {
	size := uint32(p.size)
	return uint32(y)*size*size + uint32(z)*size + uint32(x)
}

func WritePalettedContainer(w datatype.Writer, p *palettedContainer, bpe uint8) (err error) {
	err = w.WriteByte(bpe)
	if err != nil {
		return
	}

	// TODO: implement sending the palette
	// bpe can only be either 15 for blocks or 6 for biomes

	dataArray := datatype.PackIntoLongArray(bpe, p.data)
	err = datatype.WriteVarInt(w, int32(len(dataArray)))
	if err != nil {
		return
	}

	for _, val := range dataArray {
		err = datatype.WriteNumber(w, val)
		if err != nil {
			return
		}
	}

	return
}
