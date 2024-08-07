package config

import (
	"embed"
)

var (
	//go:embed generated/*
	generatedFolder embed.FS
)

// TODO: consider loading the generated data folder from the data generator instead
func init() {
	loadRegistryData()
	loadBlockStates()
}
