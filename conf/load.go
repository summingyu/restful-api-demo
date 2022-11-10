package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path=%s, err=%s", filePath, err)
	}
	return nil

}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)

}
