package cache

import (
	"fmt"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
	"gollow/cdd/util"
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
		cache.Set(data[i].GetPrimaryID(), data[i])
	}
}
