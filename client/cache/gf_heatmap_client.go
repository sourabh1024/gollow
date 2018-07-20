package cache

import (
	"gollow/data"
	"gollow/sources/datamodel"
	"sync"
)

//HeatMapDataRef is the reference object for HeatMap Data
var HeatMapDataRef = &datamodel.HeatMapData{}

type heatMapDataCache struct {
	sync.RWMutex
	cache map[string]interface{}
}

func newHeatMapDataCache() *heatMapDataCache {
	return &heatMapDataCache{
		cache: make(map[string]interface{}),
	}
}

var instance *heatMapDataCache
var once sync.Once

func GetHeatMapDataInstance() *heatMapDataCache {
	once.Do(func() {
		instance = newHeatMapDataCache()
	})
	return instance
}

func ApplyDiff(diffVersion string) error {
	//d := diff.GetNewDiffObj("", "")

	// Unmarshal the data into the sources.DataModel
	//snapshotStorage := storage.NewFileStorage(lastAnnouncedSnapshot)
	//prevBytes, err := snapshotStorage.Read()
	//err = json.Unmarshal(data, &d)

	return nil
}

func (h *heatMapDataCache) GetValue(key string) (interface{}, error) {
	h.RLock()
	defer h.RUnlock()
	val, ok := h.cache[key]
	if !ok {
		return nil, data.ErrNoData
	}
	return val, nil
}

func (h *heatMapDataCache) SetValue(key string, value interface{}) {
	h.Lock()
	defer h.Unlock()

	h.cache[key] = value
	return
}

func (h *heatMapDataCache) DeleteValue(key string) {
	h.Lock()
	defer h.Unlock()
	_, ok := h.cache[key]

	if ok {
		delete(h.cache, key)
	}
}

func (h *heatMapDataCache) Size() int {
	return len(h.cache)
}
