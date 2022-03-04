package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog"
)

// Entrypoint main listen proxy settings
type Entrypoint struct {
	Host string
	Port int
}

// Logging logging settings
type Logging struct {
	Mode   string
	Logger zerolog.Logger
}

// Global presents global proxy settings
type Global struct {
	Entrypoint Entrypoint
	Logging    Logging
}

// LoadGlobal return Global
func LoadGlobal() (*Global, error) {
	var config Global
	configData, err := ioutil.ReadFile("./global.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
