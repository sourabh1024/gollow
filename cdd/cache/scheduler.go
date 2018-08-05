package cache

import (
	"golang.org/x/net/context"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
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
