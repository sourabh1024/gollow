package cache

import (
	"gollow/data"
	"gollow/sources"
	"gollow/sources/datamodel"
	"sync"
)

//DummyDataRef is the reference object for HeatMap Data
var DummyDataRef = &datamodel.DummyData{}

var DummyDataCache = &dummyDataCache{}

type dummyDataCache struct {
	cache sync.Map
}

func (d *dummyDataCache) Get(key string) (interface{}, error) {

	val, ok := d.cache.Load(key)
	if !ok {
		return nil, data.ErrNoData
	}
	return val, nil
}

func (d *dummyDataCache) Set(key string, value sources.DataModel) {
	d.cache.Store(key, value)
}

func (d *dummyDataCache) Delete(key string) {
	d.cache.Delete(key)
}

func (d *dummyDataCache) Size() {
	return
}

//func newDummyDataCache() *heatMapDataCache {
//	return &heatMapDataCache{
//		cache: make(map[string]interface{}),
//	}
//}
//
//func GetDummyDataInstance() *dummyDataCache {
//	once.Do(func() {
//		instance = newDummyDataCache()
//	})
//	return instance
//}
//
//func BuildSnapshotDummy(lastAnnouncedSnapshot string) error {
//	// Unmarshal the data into the sources.DataModel
//	snapshotStorage := storage.NewStorage(lastAnnouncedSnapshot)
//	prevBytes, err := snapshotStorage.Read()
//
//	if err != nil {
//		logging.GetLogger().Error("Error in reading the last announced snapshot : ", lastAnnouncedSnapshot)
//		return err
//	}
//
//	data, err := producer.UnMarshalDataModelsBytes(prevBytes, HeatMapDataRef)
//
//	if err != nil {
//		logging.GetLogger().Error("Error in unmarshalling data bytes : ", err)
//		return err
//	}
//
//	logging.GetLogger().Info("Length of data from snapshot : ", len(data))
//
//	BuildCache(data)
//
//	return nil
//}
//
//func buildDummyData(data []sources.DataModel) {
//
//	defer util.Duration(time.Now(), fmt.Sprintf("Build Cache for : %d ", len(data)))
//	dummyData := GetDummyDataInstance()
//	dummyData.Lock()
//	defer dummyData.Unlock()
//	for i := 0; i < len(data); i++ {
//		if i%1000 == 0 {
//			logging.GetLogger().Info("i :", i)
//		}
//		dummyData.cache[data[i].GetPrimaryKey()] = data[i]
//	}
//}
//
//func (h *dummyDataCache) GetValue(key string) (interface{}, error) {
//	h.RLock()
//	defer h.RUnlock()
//	val, ok := h.cache[key]
//	if !ok {
//		return nil, data.ErrNoData
//	}
//	return val, nil
//}
//
//func (h *dummyDataCache) SetValue(key string, value interface{}) {
//	h.Lock()
//	defer h.Unlock()
//
//	h.cache[key] = value
//	return
//}
//
//func (h *dummyDataCache) DeleteValue(key string) {
//	h.Lock()
//	defer h.Unlock()
//	_, ok := h.cache[key]
//
//	if ok {
//		delete(h.cache, key)
//	}
//}
//
//func (h *dummyDataCache) Size() int {
//	return len(h.cache)
//}
