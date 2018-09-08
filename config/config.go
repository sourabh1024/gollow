//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package config provides method to load the config
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sourabh1024/gollow/data"
	"github.com/sourabh1024/gollow/logging"
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

	ProducerRPCPort int `json:"producerRpcPort"`

	ProducerHttpPort int `json:"producerHttpPort"`
}

// StorageConfig is the config params for Storage Stuff
type StorageConfig struct {
	// Type of Storage , only should have value supported in
	// cdd/core/storage/storage.go : only supported configs should be passed
	// by default will fall back to file storage
	StorageType string `json:"storageType"`

	// AnnouncedVersion should be announced version file name
	// this file is also accessed by the consumers , hence shouldn't be changed
	// without updating the consumers
	AnnouncedVersion string `json:"announcedVersion"`

	// BaseSnapshotPath should have the path of the folder where snapshots should be saved
	// this is used by fileStorage system
	BaseSnapshotPath string `json:"baseSnapshotPath"`

	AWSRegion string `json:"awsRegion"`
	S3Bucket  string `json:"bucket"`
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
