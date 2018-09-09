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

package cache

import (
	"github.com/sourabh1024/gollow/logging"
	"github.com/sourabh1024/gollow/sources"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var wg sync.WaitGroup

// RefreshCaches launches go routines to refresh the caches
// checks at the cache duration interval to fetch the snapshots
// returns after the cache has been loaded for the first time
// server should be started after RefreshCache has returned.
// should only be called one time
func RefreshCaches(ctx context.Context) {
	models := GetRegisteredModels()
	for model, cache := range models {
		wg.Add(1)
		go InitCache(ctx, model, cache)
		ticker := time.NewTicker(time.Duration(model.CacheDuration()))
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					// do stuff
					FetchSnapshot(ctx, model, cache)
					logging.GetLogger().Info("Updating  Snapshot : " + "-" + model.GetDataName())
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}
	logging.GetLogger().Info("Waiting for first snapshot to be fetched")
	wg.Wait()
}

// InitCache initialises the cache for the first time
func InitCache(ctx context.Context, model sources.DataModel, cache GollowCache) {
	defer wg.Done()
	FetchSnapshot(ctx, model, cache)
}
