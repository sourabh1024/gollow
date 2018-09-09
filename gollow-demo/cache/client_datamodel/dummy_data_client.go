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

// Package client_datamodel provides Cache for DummyData Cache which is populated by gollow
package client_datamodel

import (
	"github.com/sourabh1024/gollow/data"
	"github.com/sourabh1024/gollow/sources"
	"sync"
)

var DummyDataCache = &dummyDataCache{}

type dummyDataCache struct {
	Cache sync.Map
}

func (d *dummyDataCache) Get(key string) (interface{}, error) {

	val, ok := d.Cache.Load(key)
	if !ok {
		return nil, data.ErrNoData
	}
	return val, nil
}

func (d *dummyDataCache) Set(key string, value sources.Message) {
	d.Cache.Store(key, value)
}

func (d *dummyDataCache) Delete(key string) {
	d.Cache.Delete(key)
}
