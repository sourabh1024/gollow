package cache

import (
	"fmt"
	"gollow/logging"
	"gollow/sources"
	"gollow/util"
	"time"
)

type GollowCache interface {
	Get(key string) (interface{}, error)
	Set(key string, value sources.DataModel)
	Delete(key string)
}

func BuildCache(data []sources.DataModel, cache GollowCache) {
	defer util.Duration(time.Now(), fmt.Sprintf("Build Cache for : %d ", len(data)))

	for i := 0; i < len(data); i++ {
		if i%1000 == 0 {
			logging.GetLogger().Info("i :", i)
		}
		cache.Set(data[i].GetPrimaryKey(), data[i])
	}
}
