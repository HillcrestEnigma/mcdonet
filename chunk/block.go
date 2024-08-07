package chunk

import (
	"github.com/HillcrestEnigma/mcbuild/config"
)

type block struct {
	identifier string // block identifier, e.g. "minecraft:stone"
	id         int32  // block state registry ID, e.g. 0 -> "minecraft:air"
}

func NewBlockByIdentifier(identifier string, properties ...config.BlockStateProperties) (*block, error) {
	blockState, err := config.BlockStateByIdentifier(identifier, properties...)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func NewBlockByID(id int32) (*block, error) {
	blockState, err := config.BlockStateByID(id)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func (b *block) IsAir() bool {
	switch b.identifier {
	case "minecraft:air", "minecraft:cave_air", "minecraft:void_air":
		return true
	default:
		return false
	}
}

func (b *block) IsFluid() bool {
	switch b.identifier {
	case "minecraft:water", "minecraft:bubble_column", "minecraft:lava":
		return true
	default:
		return false
	}
}

// TODO: fix
func (b *block) IsMotionBlocking() bool {
	return b.IsAir()
}

func (b *block) IsLeaves() bool {
	blockInfo, err := config.BlockDataByIdentifier(b.identifier)
	if err != nil {
		panic(err)
	}

	switch blockInfo.BlockType {
	case "minecraft:leaves", "minecraft:cherry_leaves", "minecraft:mangrove_leaves":
		return true
	default:
		return false
	}
}

func newBlockFromBlockState(blockState *config.BlockState) *block {
	return &block{
		identifier: blockState.Block.Identifier,
		id:         blockState.ID,
	}
}
