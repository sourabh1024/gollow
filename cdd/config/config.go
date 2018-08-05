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
	// ENV_VAR represents the environment variable pointing to config file
	ENV_VAR = "GOLLOW_CF"
)

var (
	// GlobalConfig points to the config initialised with init
	GlobalConfig *Config
)

// Config is the config params for producer
type Config struct {

	// MySQLConfig  is the default mysql config
	MySQLConfig *data.MysqlConfig `json:"MySQLConfig"`

	// Storage configs for storage
	Storage *StorageConfig `json:"storage"`
}

// StorageConfig is the config params for Storage Stuff
type StorageConfig struct {
	// Type of Storage , only should have value supported
	// cdd/core/storage/storage.go : only supported configs should be passed
	// by default will fall back to file storage
	// by default will fall back to file storage
	StorageType string `json:"storageType"`

	// AnnouncedVersion should be announced version file name
	// this file is also accessed by the consumers , hence shouldn't be changed
	// without updating the consumers
	AnnouncedVersion string `json:"announcedVersion"`

	// BaseSnapshotPath should have the path of the folder where snapshots should be saved
	// this is used by fileStorage system
	BaseSnapshotPath string `json:"baseSnapshotPath"`
}

func init() {
	logging.GetLogger().Info("Config initialised")
	err := loadEnvFromJSON(ENV_VAR, &GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("error in loading the config : %v", err))
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
