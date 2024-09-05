package block

import (
	"github.com/HillcrestEnigma/mcdonet/config"
)

type Block struct {
	Identifier string // block identifier, e.g. "minecraft:stone"
	ID         int32  // block state registry ID, e.g. 0 -> "minecraft:air"
}

func NewBlockByIdentifier(identifier string, properties ...config.BlockStateProperties) (*Block, error) {
	blockState, err := config.BlockStateByIdentifier(identifier, properties...)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func NewBlockByID(id int32) (*Block, error) {
	blockState, err := config.BlockStateByID(id)
	if err != nil {
		return nil, err
	}

	return newBlockFromBlockState(blockState), nil
}

func newBlockFromBlockState(blockState *config.BlockState) *Block {
	return &Block{
		Identifier: blockState.Block.Identifier,
		ID:         blockState.ID,
	}
}
