package chunk

import (
	"bytes"
	"math"

	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/HillcrestEnigma/mcbuild/world/block"
)

type ChunkColumn struct {
	X             int32
	Z             int32
	min_y         int32
	height        int32
	heightmaps    *heightmaps
	chunkSections []*chunkSection
}

type ChunkColumnBlock struct {
	chunkSectionBlock
	ChunkColumn *ChunkColumn
	Y     int32
}

// Consider accepting a Biome object instead of a uint32
func NewChunkColumn(x, z, min_y, height int32) (chunk *ChunkColumn) {
	chunkSections := make([]*chunkSection, height/16)
	chunk = &ChunkColumn{
		X:             x,
		Z:             z,
		min_y:         min_y,
		height:        height,
		heightmaps:    newHeightmaps(),
		chunkSections: chunkSections,
	}

	for i := range chunkSections {
		chunkSections[i] = newChunkSection(chunk)
	}
	return
}

func (c *ChunkColumn) Block(sectionX uint8, y int32, sectionZ uint8) (*ChunkColumnBlock, error) {
	sectionIndex := (y - c.min_y) / 16
	sectionY := uint8(y % 16)

	block, err := c.chunkSections[sectionIndex].block(sectionX, sectionY, sectionZ)
	if err != nil {
		return nil, err
	}

	return &ChunkColumnBlock{
		chunkSectionBlock: *block,
		ChunkColumn:             c,
		Y:                 y,
	}, nil
}

func (c *ChunkColumn) SetBlock(sectionX uint8, y int32, sectionZ uint8, newBlock *block.Block) error {
	chunkBlock, err := c.Block(sectionX, y, sectionZ)
	if err != nil {
		return err
	}

	chunkBlock.chunkSectionBlock.set(newBlock)

	err = c.recomputeHeightAtSectionXZ(sectionX, y, sectionZ)
	if err != nil {
		return err
	}
	return nil
}

func (b *ChunkColumnBlock) set(newBlock *block.Block) {
	b.ChunkColumn.SetBlock(b.SectionX, b.Y, b.SectionZ, newBlock)
}

func WriteNetworkChunkColumn(w datatype.Writer, c *ChunkColumn) (err error) {
	err = datatype.WriteNumber(w, c.X)
	if err != nil {
		return
	}

	err = datatype.WriteNumber(w, c.Z)
	if err != nil {
		return
	}

	heightmapsBPE := uint8(math.Ceil(math.Log2(float64(c.height + 1))))
	heightmapsNBT := &datatype.NBT{
		Name: "Heightmaps",
		Compound: datatype.NBTCompound{
			"WORLD_SURFACE":             datatype.PackIntoLongArray(heightmapsBPE, c.heightmaps[0]),
			"OCEAN_FLOOR":               datatype.PackIntoLongArray(heightmapsBPE, c.heightmaps[1]),
			"MOTION_BLOCKING":           datatype.PackIntoLongArray(heightmapsBPE, c.heightmaps[2]),
			"MOTION_BLOCKING_NO_LEAVES": datatype.PackIntoLongArray(heightmapsBPE, c.heightmaps[3]),
		},
	}
	err = datatype.WriteNetworkNBT(w, heightmapsNBT)

	var buf bytes.Buffer
	for _, section := range c.chunkSections {
		err = WriteChunkSection(&buf, section)
		if err != nil {
			return
		}
	}

	datatype.WriteVarInt(w, int32(buf.Len()))
	w.Write(buf.Bytes())

	return
}
