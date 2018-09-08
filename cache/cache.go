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

// Package cache provides all required methods for client side cache building
package cache

import (
	"fmt"
	"github.com/gollow/logging"
	"github.com/gollow/sources"
	"github.com/gollow/util"
	"time"
)

//GollowCache provides all method to fetch data
type GollowCache interface {

	//Get gets the value from the cache
	Get(key string) (interface{}, error)

	//Set sets the value in cache
	Set(key string, value sources.Message)

	//Delete deletes the key from cache
	Delete(key string)
}

//buildCache builds the cache
//puts all the elements of bag into the cache
//future : should have some random sleeps to avoid contention on reads
func buildCache(bag sources.Bag, cache GollowCache) {
	data := bag.GetEntries()
	defer util.Duration(time.Now(), fmt.Sprintf("Build Cache for : %d ", len(data)))

	for i := 0; i < len(data); i++ {
		if i > 0 && i%100000 == 0 {
			logging.GetLogger().Info("i :", i)
			//TODO : add some random sleep to avoid starvation for read
		}
		cache.Set(data[i].GetUniqueKey(), data[i])
	}
}
