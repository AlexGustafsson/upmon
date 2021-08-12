package configuration

import (
	"fmt"
	"os"

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

	// If no name is set, default to the hostname
	if config.Name == "" {
		name, err := os.Hostname()
		if err == nil {
			config.Name = name
		}
	}

	for _, service := range config.Services {
		service.Origin = config.Name
		if service.Name == "" {
			service.Name = "unnamed service"
		}
		for _, monitor := range service.Monitors {
			if monitor.Name == "" {
				monitor.Name = fmt.Sprintf("unnamed '%s' monitor", monitor.Type)
			}
		}
	}

	return config, nil
}
