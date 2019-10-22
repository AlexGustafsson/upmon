package main

import (
  "io/ioutil"
	"github.com/BurntSushi/toml"
	"github.com/AlexGustafsson/upmon/core"
)

// LoadConfig reads and parses a given TOML configuration file
func loadConfig(path string) (*core.Config, error) {
  core.LogDebug("Loading config from '%s'", path)
  dataBuffer, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }

  data := string(dataBuffer)

  var config = new(core.Config)
  _, err = toml.Decode(data, config);
  if err != nil {
    return nil, err
  }

  return config, nil
}
