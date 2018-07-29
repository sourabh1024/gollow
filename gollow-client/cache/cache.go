package cache

import (
	"fmt"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
	"gollow/cdd/util"
	"time"
)

type GollowCache interface {
	Get(key string) (interface{}, error)
	Set(key string, value sources.Message)
	Delete(key string)
}

func BuildCache(bag sources.Bag, cache GollowCache) {
	data := bag.GetEntries()
	defer util.Duration(time.Now(), fmt.Sprintf("Build Cache for : %d ", len(data)))

	for i := 0; i < len(data); i++ {
		if i%100000 == 0 {
			logging.GetLogger().Info("i :", i)
			//TODO : add some random sleep to avoid starvation for read
		}
		cache.Set(data[i].GetPrimaryID(), data[i])
	}
}
