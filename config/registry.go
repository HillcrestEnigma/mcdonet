package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/HillcrestEnigma/mcdonet/datatype"
)

type RegistryEntry struct {
	Identifier string
	ID         int32
	Data       any
}

var (
	availableRegistries = [9]string{
		"minecraft:trim_material",
		"minecraft:trim_pattern",
		"minecraft:banner_pattern",
		"minecraft:worldgen/biome",
		"minecraft:chat_type",
		"minecraft:damage_type",
		"minecraft:dimension_type",
		"minecraft:wolf_variant",
		"minecraft:painting_variant",
	}
	RegistriesByIdentifier = make(map[string]map[string]*RegistryEntry)
	RegistriesByID         = make(map[string]map[int32]*RegistryEntry)
)

func RegistryByIdentifier(identifier string) (map[string]*RegistryEntry, error) {
	r, ok := RegistriesByIdentifier[identifier]
	if !ok {
		return nil, fmt.Errorf("registry not found for identifier %s", identifier)
	}

	return r, nil
}

func RegistryByID(identifier string) (map[int32]*RegistryEntry, error) {
	r, ok := RegistriesByID[identifier]
	if !ok {
		return nil, fmt.Errorf("registry not found for identifier %s", identifier)
	}

	return r, nil
}

func loadRegistryData() {
	for _, registryIdentifier := range availableRegistries {
		registryIdentifierNamespace, registryIdentifierValue, err := datatype.ParseIdentifier(registryIdentifier)
		if err != nil {
			log.Fatal(err)
		}

		RegistriesByIdentifier[registryIdentifier] = make(map[string]*RegistryEntry)
		RegistriesByID[registryIdentifier] = make(map[int32]*RegistryEntry)

		entryID := int32(0)
		registryDataPath := filepath.Join("generated", "data", registryIdentifierNamespace, filepath.FromSlash(registryIdentifierValue))
		err = fs.WalkDir(generatedFolder, registryDataPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			entryIdentifierValue := strings.TrimSuffix(filepath.Base(path), ".json")
			entryIdentifier := fmt.Sprintf("%s:%s", registryIdentifierNamespace, entryIdentifierValue)
			entry := RegistryEntry{
				Identifier: entryIdentifier,
				ID:         entryID,
			}

			entryFileData, err := fs.ReadFile(generatedFolder, path)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(entryFileData, &entry.Data)
			if err != nil {
				log.Fatal(err)
			}

			RegistriesByIdentifier[registryIdentifier][entryIdentifier] = &entry
			RegistriesByID[registryIdentifier][entryID] = &entry
			entryID++
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
