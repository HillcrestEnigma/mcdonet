package chunk

import (
	"github.com/HillcrestEnigma/mcbuild/datatype"
	"github.com/HillcrestEnigma/mcbuild/world/block"
)

type chunkSection struct {
	chunkColumn *ChunkColumn
	blockCount  int16
	blockStates palettedContainer
	biomes      palettedContainer
}

type chunkSectionBlock struct {
	block.Block
	ChunkSection *chunkSection
	SectionX     uint8
	SectionY     uint8
	SectionZ     uint8
}

func newChunkSection(chunkColumn *ChunkColumn) (s *chunkSection) {
	return &chunkSection{
		chunkColumn: chunkColumn,
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

	b, err := block.NewBlockByID(id)
	if err != nil {
		panic(err)
	}

	return &chunkSectionBlock{
		Block:        *b,
		ChunkSection: s,
		SectionX:     sectionX,
		SectionY:     sectionY,
		SectionZ:     sectionZ,
	}, nil
}

func (s *chunkSection) setBlock(sectionX, sectionY, sectionZ uint8, newBlock *block.Block) error {
	oldBlock, err := s.block(sectionX, sectionY, sectionZ)
	if err != nil {
		return err
	}

	if oldBlock.Identifier == newBlock.Identifier {
		return nil
	}

	if !oldBlock.IsAir() {
		s.blockCount--
	}
	if !newBlock.IsAir() {
		s.blockCount++
	}

	s.blockStates.set(sectionX, sectionY, sectionZ, newBlock.ID)

	return nil
}

func (b *chunkSectionBlock) set(newBlock *block.Block) {
	b.ChunkSection.setBlock(b.SectionX, b.SectionY, b.SectionZ, newBlock)
	b.Block = *newBlock
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
