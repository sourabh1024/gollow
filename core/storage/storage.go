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

// Package storage provides methods related to storage handling
package storage

import (
	"github.com/sourabh1024/gollow/config"
)

//Storage interface provides all methods required for read/write/checkExistence
//of any snapshot or diff
type Storage interface {
	IsExist() bool

	Write(data []byte) (int, error)

	Read() ([]byte, error)
}

// NewStorage gives the current implementation of Storage
// returns storage object if initialised properly
// throws error if storage cannot be initialised
// Default it returns file storage
func NewStorage(key string) (Storage, error) {
	switch storage := config.GlobalConfig.Storage.StorageType; storage {
	case "file":
		return NewFileStorage(key)
	case "s3":
		return NewS3Storage(&Config{
			Region: config.GlobalConfig.Storage.AWSRegion,
			Bucket: config.GlobalConfig.Storage.S3Bucket,
			Key:    key,
		})
	default:
		return NewFileStorage(key)
	}
}
