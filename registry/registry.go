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
