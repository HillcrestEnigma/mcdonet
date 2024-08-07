package registry

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

var (
	//go:embed blocks.json
	blocksJSONData          []byte
	blocksByIdentifier      map[string]*block
	blockStatesByRegistryID map[int32]*BlockState
)

type block struct {
	Identifier   string
	definition   map[string]any
	BlockType    string
	defaultState *BlockState
	states       []*BlockState
	properties   map[string][]string
}

type BlockStateProperties map[string]string

type BlockState struct {
	Block      *block
	RegistryID int32
	properties BlockStateProperties
}

type blockStateRegistryBlock map[string]struct {
	Definition map[string]any      `json:"definition"`
	Properties map[string][]string `json:"properties"`
	States     []struct {
		IsDefault  bool                 `json:"default"`
		Id         int32                `json:"id"`
		Properties BlockStateProperties `json:"properties"`
	} `json:"states"`
}

func loadBlockStateRegistry() (err error) {
	blocksByIdentifier = make(map[string]*block)
	blockStatesByRegistryID = make(map[int32]*BlockState)

	var blockStateRegistry blockStateRegistryBlock
	err = json.Unmarshal(blocksJSONData, &blockStateRegistry)
	if err != nil {
		return
	}

	for identifier, blockTypeData := range blockStateRegistry {
		definition := blockTypeData.Definition
		blockType, ok := blockTypeData.Definition["type"].(string)
		if !ok {
			return fmt.Errorf("block.definition.type not found for block %s", identifier)
		}
		delete(definition, "type")

		block := &block{
			Identifier: identifier,
			definition: definition,
			BlockType:  blockType,
			properties: blockTypeData.Properties,
		}

		blockStates := make([]*BlockState, len(blockTypeData.States))
		for blockStateIndex, blockStateData := range blockTypeData.States {
			blockState := &BlockState{
				Block:      block,
				RegistryID: blockStateData.Id,
				properties: blockStateData.Properties,
			}

			if blockStateData.IsDefault {
				block.defaultState = blockState
			}

			blockStates[blockStateIndex] = blockState
			blockStatesByRegistryID[blockState.RegistryID] = blockState
		}

		block.states = blockStates
		blocksByIdentifier[identifier] = block
	}

	return
}

func BlockByIdentifier(identifier string) (*block, error) {
	b, ok := blocksByIdentifier[identifier]
	if !ok {
		return nil, fmt.Errorf("block %s not found", identifier)
	}
	return b, nil
}

func BlockStateByRegistryID(registryID int32) (*BlockState, error) {
	blockState, ok := blockStatesByRegistryID[registryID]
	if !ok {
		return nil, fmt.Errorf("block state %d not found", registryID)
	}
	return blockState, nil
}

func BlockStateByIdentifier(identifier string, properties ...BlockStateProperties) (*BlockState, error) {
	b, err := BlockByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	if len(properties) == 0 {
		return b.defaultState, nil
	}

blockState:
	for _, blockState := range b.states {
		for _, property := range properties {
			for key, value := range property {
				if blockState.properties[key] != value {
					continue blockState
				}
			}
		}

		return blockState, nil
	}

	return nil, fmt.Errorf("no matching block state found for specified properties")
}

func BlockStatePropertiesByRegistryID(registryID int32) (BlockStateProperties, error) {
	blockState, err := BlockStateByRegistryID(registryID)
	if err != nil {
		return nil, err
	}
	return blockState.properties, nil
}
