package block

import "github.com/HillcrestEnigma/mcdonet/config"

func (b *Block) IsAir() bool {
	switch b.Identifier {
	case "minecraft:air", "minecraft:cave_air", "minecraft:void_air":
		return true
	default:
		return false
	}
}

func (b *Block) IsFluid() bool {
	switch b.Identifier {
	case "minecraft:water", "minecraft:bubble_column", "minecraft:lava":
		return true
	default:
		return false
	}
}

// TODO: fix
func (b *Block) IsMotionBlocking() bool {
	return b.IsAir()
}

func (b *Block) IsLeaves() bool {
	blockInfo, err := config.BlockDataByIdentifier(b.Identifier)
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
