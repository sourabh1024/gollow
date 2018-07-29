package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"gollow/cdd/data"
	"gollow/cdd/logging"
	"io/ioutil"
	"os"
)

const (
	ENV_VAR = "GOLLOW_CF"
)

var (
	GlobalConfig *Config
	MySQLConfig  = &data.MysqlConfig{}
)

type Config struct {
	MySQLConfig      *data.MysqlConfig `json:"MySQLConfig"`
	AnnouncedVersion string            `json:"announcedVersion"`
}

func init() {
	logging.GetLogger().Info("Config initialised")
	err := loadEnvFromJSON(ENV_VAR, &GlobalConfig)
	if err != nil {
		panic(errors.New(fmt.Sprintf("error in loading the config : %v", err)))
	}
}

func loadEnvFromJSON(envVar string, config interface{}) error {
	filename := os.Getenv(envVar)
	logging.GetLogger().Info("Getting config from : %s", filename)
	if config == nil {
		return errors.New("config is nil")
	}

	if filename == "" {
		logging.GetLogger().Error("Falling back to development config ")
		filename = "../config-development.json"
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, config)
}
