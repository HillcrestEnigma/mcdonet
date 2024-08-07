package config

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	blockDataByIdentifier map[string]*BlockData
	blockStateByID        map[int32]*BlockState
)

type BlockStateProperties map[string]string

type BlockData struct {
	Identifier   string
	Definition   map[string]any
	BlockType    string
	DefaultState *BlockState
	States       []*BlockState
	Properties   map[string][]string
}

type BlockState struct {
	Block      *BlockData
	ID         int32
	properties BlockStateProperties
}

func BlockDataByIdentifier(identifier string) (*BlockData, error) {
	b, ok := blockDataByIdentifier[identifier]
	if !ok {
		return nil, fmt.Errorf("block data not found for identifier %s", identifier)
	}

	return b, nil
}

func BlockStateByID(id int32) (*BlockState, error) {
	b, ok := blockStateByID[id]
	if !ok {
		return nil, fmt.Errorf("block state not found for id %d", id)
	}

	return b, nil
}

func BlockStateByIdentifier(identifier string, properties ...BlockStateProperties) (*BlockState, error) {
	b, ok := blockDataByIdentifier[identifier]
	if !ok {
		return nil, fmt.Errorf("block state not found for identifier %s", identifier)
	}

	if len(properties) == 0 {
		return b.DefaultState, nil
	}

blockState:
	for _, blockState := range b.States {
		for _, property := range properties {
			for key, value := range property {
				if blockState.properties[key] != value {
					continue blockState
				}
			}
		}

		return blockState, nil
	}

	return nil, fmt.Errorf("block state not found for identifier %s and given properties", identifier)
}

func loadBlockStates() {
	var blockStateRegistry map[string]struct {
		Definition map[string]any      `json:"definition"`
		Properties map[string][]string `json:"properties"`
		States     []struct {
			IsDefault  bool                 `json:"default"`
			Id         int32                `json:"id"`
			Properties BlockStateProperties `json:"properties"`
		} `json:"states"`
	}

	blockDataByIdentifier = make(map[string]*BlockData)
	blockStateByID = make(map[int32]*BlockState)

	blockStatesFileData, err := generatedFolder.ReadFile("generated/reports/blocks.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(blockStatesFileData, &blockStateRegistry)
	if err != nil {
		log.Fatal(err)
	}

	for identifier, blockTypeData := range blockStateRegistry {
		definition := blockTypeData.Definition
		blockType, ok := blockTypeData.Definition["type"].(string)
		if !ok {
			log.Fatalf("block.definition.type not found for block %s", identifier)
		}
		delete(definition, "type")

		block := &BlockData{
			Identifier: identifier,
			Definition: definition,
			BlockType:  blockType,
			Properties: blockTypeData.Properties,
		}

		blockStates := make([]*BlockState, len(blockTypeData.States))
		for blockStateIndex, blockStateData := range blockTypeData.States {
			blockState := &BlockState{
				Block:      block,
				ID:         blockStateData.Id,
				properties: blockStateData.Properties,
			}

			if blockStateData.IsDefault {
				block.DefaultState = blockState
			}

			blockStates[blockStateIndex] = blockState
			blockStateByID[blockState.ID] = blockState
		}

		block.States = blockStates
		blockDataByIdentifier[identifier] = block
	}
}
