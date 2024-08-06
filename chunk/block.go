package chunk

import (
	"github.com/HillcrestEnigma/mcbuild/registry"
)

type block struct {
	Identifier string // block identifier, e.g. "minecraft:stone"
	RegistryID int32  // block state registry ID, e.g. 0 -> "minecraft:air"
}

func NewBlockByIdentifier(identifier string, properties ...registry.BlockStateProperties) (*block, error) {
	blockState, err := registry.GetBlockStateByIdentifier(identifier, properties...)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func NewBlockByRegistryID(registryID int32) (*block, error) {
	blockState, err := registry.GetBlockStateByRegistryID(registryID)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func (b *block) IsAir() bool {
	switch b.Identifier {
	case "minecraft:air", "minecraft:cave_air", "minecraft:void_air":
		return true
	default:
		return false
	}
}

func (b *block) IsFluid() bool {
	switch b.Identifier {
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
	blockInfo, err := registry.GetBlockByIdentifier(b.Identifier)
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

func newBlockFromBlockState(blockState *registry.BlockState) *block {
	return &block{
		Identifier: blockState.Block.Identifier,
		RegistryID: blockState.RegistryID,
	}
}
