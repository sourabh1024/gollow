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

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sourabh1024/gollow/logging"
	"io/ioutil"
	"os"
)

const (
	ENV_VAR = "GOLLOW_DEMO_CF"
)

var (
	GlobalConfig *Config
)

type Config struct {
	AnnouncedVersion string         `json:"announcedVersion"`
	Storage          *StorageConfig `json:"storage"`
	ServerRPCPort    int            `json:"serverRpcPort"`
	ServerHttpPort   int            `json:"serverHttpPort"`
}

type StorageConfig struct {
	StorageType      string `json:"storageType"`
	AnnouncedVersion string `json:"announcedVersion"`
	AWSRegion        string `json:"awsRegion"`
	S3Bucket         string `json:"bucket"`
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
