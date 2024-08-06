package chunk

import (
	"bytes"
	"math"

	"github.com/HillcrestEnigma/mcbuild/datatype"
)

type Chunk struct {
	X             int32
	Z             int32
	min_y         int32
	height        int32
	heightmaps    heightmaps
	chunkSections []chunkSection
}

type ChunkBlock struct {
	chunkSectionBlock
	Chunk *Chunk
	Y     int32
}

// Consider accepting a Biome object instead of a uint32
func NewChunk(x, z, min_y, height int32) (chunk *Chunk) {
	chunkSections := make([]chunkSection, height/16)
	for i := range chunkSections {
		chunkSections[i] = newChunkSection()
	}

	chunk = &Chunk{
		X:             x,
		Z:             z,
		min_y:         min_y,
		height:        height,
		heightmaps:    newHeightmaps(),
		chunkSections: chunkSections,
	}
	return
}

func (c *Chunk) GetBlock(sectionX uint8, y int32, sectionZ uint8) *ChunkBlock {
	sectionIndex := (y - c.min_y) / 16
	sectionY := uint8(y % 16)

	block := c.chunkSections[sectionIndex].getBlock(sectionX, sectionY, sectionZ)

	return &ChunkBlock{
		chunkSectionBlock: *block,
		Chunk:             c,
		Y:                 y,
	}
}

func (c *Chunk) SetBlock(sectionX uint8, y int32, sectionZ uint8, block *block) {
	chunkBlock := c.GetBlock(sectionX, y, sectionZ)
	chunkBlock.set(block)
	c.recomputeHeightAtSectionXZ(sectionX, y, sectionZ)
}

func WriteNetworkChunk(w datatype.Writer, c *Chunk) (err error) {
	err = datatype.WriteNumber(w, c.X)
	if err != nil {
		return
	}

	datatype.WriteNumber(w, c.Z)
	if err != nil {
		return
	}

	heightmapsBPE := uint8(math.Ceil(math.Log2(float64(c.height + 1))))
	heightmapsNBT := datatype.NBT{
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
		err = WriteChunkSection(&buf, &section)
		if err != nil {
			return
		}
	}

	datatype.WriteVarInt(w, int32(buf.Len()))
	w.Write(buf.Bytes())

	return
}
