package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"strconv"
	"sync"
)

type BSYearConfig struct {
	DaysInMonths    [12]int `json:"days_in_months"`
	StartDateUnixTs int64   `json:"start_date"`
}

type BSConfig struct {
	Years map[int]BSYearConfig
}

//go:embed years.json
var yearsJSON []byte

var (
	once     sync.Once
	instance *BSConfig
)

func GetConfig() BSConfig {
	// Load once
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
