package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"strconv"
	"sync"
)

//go:embed years.json
var yearsJSON []byte

var (
	once     sync.Once
	instance *BSConfig
)

func GetConfig() BSConfig {
	once.Do(func() {
		var raw map[string]BSYearConfig
		if err := json.Unmarshal(yearsJSON, &raw); err != nil {
			panic(fmt.Sprintf("config: failed to parse years.json: %s", err))
		}

		cfg := &BSConfig{Years: make(map[int]BSYearConfig, len(raw))}
		for yearStr, yc := range raw {
			year, err := strconv.Atoi(yearStr)
			if err != nil {
				panic(fmt.Sprintf("config: invalid year key %q: %s", yearStr, err))
			}
			cfg.Years[year] = yc
		}

		instance = cfg
	})

	configCopy := make(map[int]BSYearConfig, len(instance.Years))
	// yearConfig contains array and int so no deep clone is needed
	maps.Copy(configCopy, instance.Years)
	return BSConfig{Years: configCopy}
}
