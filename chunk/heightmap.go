package chunk

// heightmap is indexed [z][x]
type heightmap []int32

// index 0: WORLD_SURFACE
// index 1: OCEAN_FLOOR
// index 2: MOTION_BLOCKING
// index 3: MOTION_BLOCKING_NO_LEAVES
type heightmaps [4]heightmap

func newHeightmaps() *heightmaps {
	var h heightmaps
	for i := range h {
		h[i] = make(heightmap, 16*16)
	}

	return &h
}

func (h *heightmap) get(x, z uint8) int32 {
	return (*h)[z*16+x]
}

func (h *heightmap) set(x, z uint8, value int32) {
	(*h)[z*16+x] = value
}

func (c *Chunk) recomputeHeightAtSectionXZ(sectionX uint8, modifiedY int32, sectionZ uint8) error {
	startY := modifiedY
	for i := range c.heightmaps {
		startY = max(startY, c.heightmaps[i].get(sectionX, sectionZ))
	}

	var heights [4]int32
	var foundHeights byte // Stores if the correct height for each heightmap type has been found
	// 0, 0x01: WORLD_SURFACE
	// 1, 0x02: OCEAN_FLOOR
	// 2, 0x04: MOTION_BLOCKING
	// 3, 0x08: MOTION_BLOCKING_NO_LEAVES

	for y := startY; y >= 0; y-- {
		block, err := c.Block(sectionX, y, sectionZ)
		if err != nil {
			return err
		}

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
		if foundHeights&0x04 == 0 && block.IsMotionBlocking() && !block.IsLeaves() {
			heights[3] = y + 1
			foundHeights |= 0x08
		}
		if foundHeights == 0x07 {
			break
		}
	}

	for i := range heights {
		c.heightmaps[i].set(sectionX, sectionZ, heights[i])
	}
	return nil
}
