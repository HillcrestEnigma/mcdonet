package registry

import (
	_ "embed"
	"encoding/json"
	"sync"
)

var (
	//go:embed login.json
	loginJSONData        []byte
	Login                map[string][]string
	loadRegistryDataOnce sync.Once
)

// TODO: consider loading the generated data folder from the data generator instead
func LoadRegistryData() (err error) {
	loadRegistryDataOnce.Do(func() {
		err = json.Unmarshal(loginJSONData, &Login)
		if err != nil {
			return
		}

		loadBlockStateRegistry()
	})

	return
}
