/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package cmd

import (
	"io/ioutil"

	"github.com/mixj93/helm-registry/pkg/common"
	"github.com/mixj93/helm-registry/pkg/log"
	"github.com/ghodss/yaml"
)

// Manager is a config of Space Manager
type Manager struct {
	// Name
	Name string `yaml:"name"`

	// Parameters of Space Manager
	Parameters map[string]interface{} `yaml:"parameters"`
}

// Config is a config of the application
type Config struct {
	// Listen address
	Listen string `yaml:"listen"`

	// Manager config
	Manager Manager `yaml:"manager"`
}

// newDefaultConfig creates a default config
func newDefaultConfig() *Config {
	return &Config{
		Listen: ":10080",
		Manager: Manager{
			Name: "simple",
			Parameters: map[string]interface{}{
				common.ParameterNameStorageDriver: "filesystem",
				common.ParameterNameRootDirectory: "/var/lib/helm",
				common.ParameterResourceLocker:    "memory",
			},
		},
	}
}

// newConfig creates config from file
func newConfig(filepath string) (*Config, error) {
	config := newDefaultConfig()
	if configPath != "" {
		file, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(file, config)
		if err != nil {
			return nil, err
		}
	} else {
		log.Printf("Using default config %v", config)
	}
	return config, nil
}
