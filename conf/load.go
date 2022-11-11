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
	// return loadGlobal()

}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return fmt.Errorf("load config from env error, err=%s", err)
	}
	return nil
	// return loadGlobal()

}

// func loadGlobal() error {
// 	var err error
// 	db, err = config.MySQL.getDBConn()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
