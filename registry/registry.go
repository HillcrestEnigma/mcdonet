package registry

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed login.json
	loginJSONData []byte
	Login         map[string][]string
)

// TODO: consider loading the generated data folder from the data generator instead
func init() {
	json.Unmarshal(loginJSONData, &Login)
	loadBlockStateRegistry()
}
