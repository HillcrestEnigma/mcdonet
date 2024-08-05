package chunk

type Chunk struct {
	X             int32
	Z             int32
	min_y         int32
	height        int32
	heightmaps    heightmaps
	chunkSections []*chunkSection
}

type ChunkBlock struct {
	chunkSectionBlock
	Chunk *Chunk
	Y     int32
}

func NewChunk(x, z, min_y, height int32) (chunk *Chunk) {
	chunkSections := make([]*chunkSection, height/16)
	for i := range chunkSections {
		chunkSections[i] = newChunkSection()
	}

	chunk = &Chunk{
		X:             x,
		Z:             z,
		min_y:         min_y,
		height:        height,
		heightmaps:    *newHeightmaps(),
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

func (c *Chunk) recomputeHeightAtSectionXZ(sectionX uint8, modifiedY int32, sectionZ uint8) {
	startY := modifiedY
	for i := 0; i < 3; i++ {
		startY = max(startY, c.heightmaps[i].get(sectionX, sectionZ))
	}

	heights := make([]int32, 3)
	var foundHeights byte // Stores if the correct height for each heightmap type has been found
	// 0, 0x01: WORLD_SURFACE
	// 1, 0x02: OCEAN_FLOOR
	// 2, 0x04: MOTION_BLOCKING

	for y := startY; y >= 0; y-- {
		block := c.GetBlock(sectionX, y, sectionZ)

		if foundHeights&0x01 == 0 && !block.IsAir() {
			heights[0] = y + 1
			foundHeights |= 0x01
		}
		if foundHeights&0x02 == 0 && !block.IsAir() && !block.IsFluid() {
			heights[1] = y + 1
			foundHeights |= 0x02
		}
		if foundHeights&0x04 == 0 && block.IsMotionBlocking() {
			heights[2] = y + 1
			foundHeights |= 0x04
		}
		if foundHeights == 0x07 {
			break
		}
	}

	for i := range heights {
		c.heightmaps[i].set(sectionX, sectionZ, heights[i])
	}
}
