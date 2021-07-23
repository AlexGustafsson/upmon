package configuration

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
)

// Load a configuration from a YAML file
func Load(filePath string) (*Configuration, error) {
	k := koanf.New(".")

	defaults := confmap.Provider(map[string]interface{}{
		"address":     "127.0.0.1",
		"port":        "7070",
		"api.enabled": false,
	}, ".")
	err := k.Load(defaults, nil)
	if err != nil {
		return nil, err
	}

	provider := file.Provider(filePath)
	err = k.Load(provider, yaml.Parser())
	if err != nil {
		return nil, err
	}

	config := &Configuration{}
	err = k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "koanf"})
	if err != nil {
		return nil, err
	}

	// Set default names if none exist
	for _, service := range config.Services {
		for _, monitor := range service.Monitors {
			if monitor.Name == "" {
				monitor.Name = fmt.Sprintf("unnamed '%s' monitor", monitor.Type)
			}
		}
	}

	return config, nil
}
