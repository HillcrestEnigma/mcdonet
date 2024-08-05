package chunk

// heightmap is indexed [z][x]
type heightmap []int32

// index 0: WORLD_SURFACE
// index 1: OCEAN_FLOOR
// index 2: MOTION_BLOCKING
type heightmaps []heightmap

func newHeightmaps() *heightmaps {
	h := make(heightmaps, 3)
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