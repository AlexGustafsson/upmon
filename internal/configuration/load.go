package configuration

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
)

func Load(filePath string) (*Configuration, error) {
	k := koanf.New(".")

	defaults := confmap.Provider(map[string]interface{}{
		"address":     "",
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

	return config, nil
}
