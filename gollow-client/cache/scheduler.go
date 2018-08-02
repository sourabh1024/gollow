package cache

import (
	"golang.org/x/net/context"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
	"gollow/cdd/sources/datamodel/dummy"
	"gollow/gollow-client/cache/client_datamodel"
	"sync"
	"time"
)

var wg sync.WaitGroup

func UpdateSnapshots(ctx context.Context) {
	models := GetRegisteredModels()
	for model, cache := range models {
		wg.Add(1)
		go UpdateFirstTimeSnapshot(ctx, model, cache)
		ticker := time.NewTicker(time.Duration(model.CacheDuration()))
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					// do stuff
					logging.GetLogger().Info("Updating  Snapshot : " + "-" + model.GetDataName())
					FetchSnapshot(ctx, model, cache)
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

func UpdateFirstTimeSnapshot(ctx context.Context, model sources.DataModel, cache GollowCache) {
	defer wg.Done()
	FetchSnapshot(ctx, model, cache)
}

// for demo purpose
func ReadValue() {
	ticker := time.NewTicker(time.Duration(30 * time.Second))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				logging.GetLogger().Info("Read data")
				val, err := client_datamodel.DummyDataCache.Get("1")
				//client_datamodel.DummyDataCache.Cache.Range(func(key, value interface{}) bool {
				//	logging.GetLogger().Info("value fetched : ", value)
				//	return true
				//})
				if err != nil {
					logging.GetLogger().Error("Error in reading value : ", err)
					continue
				}
				dummyData, ok := val.(*dummy.DummyData)
				if !ok {
					logging.GetLogger().Error("Error in parsing value : ")
					continue
				}
				logging.GetLogger().Info("created at : ", dummyData.FirstName)
				logging.GetLogger().Info("Value for id 1 : ", dummyData.MaxDebit)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
