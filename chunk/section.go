package chunk

import "github.com/HillcrestEnigma/mcbuild/datatype"

type chunkSection struct {
	blockCount  int16
	blockStates palettedContainer
	biomes      palettedContainer
}

type chunkSectionBlock struct {
	block
	ChunkSection *chunkSection
	SectionX     uint8
	SectionY     uint8
	SectionZ     uint8
}

func newChunkSection() (s *chunkSection) {
	return &chunkSection{
		blockCount:  0,
		blockStates: *newPalettedContainer(16, 0),
		biomes:      *newPalettedContainer(4, 0),
		// Change default biome type from 0 to something else
	}
}

func (s *chunkSection) block(sectionX, sectionY, sectionZ uint8) (*chunkSectionBlock, error) {
	id, err := s.blockStates.get(sectionX, sectionY, sectionZ)
	if err != nil {
		return nil, err
	}

	block, err := NewBlockByID(id)
	if err != nil {
		panic(err)
	}

	return &chunkSectionBlock{
		block:        *block,
		ChunkSection: s,
		SectionX:     sectionX,
		SectionY:     sectionY,
		SectionZ:     sectionZ,
	}, nil
}

func (s *chunkSection) setBlock(sectionX, sectionY, sectionZ uint8, newBlock *block) error {
	oldBlock, err := s.block(sectionX, sectionY, sectionZ)
	if err != nil {
		return err
	}

	if oldBlock.identifier == newBlock.identifier {
		return nil
	}

	if !oldBlock.IsAir() {
		s.blockCount--
	}
	if !newBlock.IsAir() {
		s.blockCount++
	}

	s.blockStates.set(sectionX, sectionY, sectionZ, newBlock.id)
	
	return nil
}

func (b *chunkSectionBlock) set(newBlock *block) {
	b.ChunkSection.setBlock(b.SectionX, b.SectionY, b.SectionZ, newBlock)
	b.block = *newBlock
}

func WriteChunkSection(w datatype.Writer, s *chunkSection) (err error) {
	err = datatype.WriteNumber(w, s.blockCount)
	if err != nil {
		return
	}

	err = WritePalettedContainer(w, &s.blockStates, 15)
	if err != nil {
		return
	}

	err = WritePalettedContainer(w, &s.biomes, 6)
	if err != nil {
		return
	}

	return
}
