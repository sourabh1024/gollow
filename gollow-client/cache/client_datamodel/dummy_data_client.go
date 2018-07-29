package client_datamodel

import (
	"gollow/cdd/data"
	"gollow/cdd/sources"
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
