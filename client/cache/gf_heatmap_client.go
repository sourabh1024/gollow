package cache

import (
	"fmt"
	"gollow/data"
	"gollow/logging"
	"gollow/producer"
	"gollow/sources"
	"gollow/sources/datamodel"
	"gollow/storage"
	"gollow/util"
	"sync"
	"time"
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

func BuildSnapshot(lastAnnouncedSnapshot string) error {
	// Unmarshal the data into the sources.DataModel
	snapshotStorage := storage.NewFileStorage(lastAnnouncedSnapshot)
	prevBytes, err := snapshotStorage.Read()

	if err != nil {
		logging.GetLogger().Error("Error in reading the last announced snapshot : ", lastAnnouncedSnapshot)
		return err
	}

	data, err := producer.UnMarshalDataModelsBytes(prevBytes, HeatMapDataRef)

	if err != nil {
		logging.GetLogger().Error("Error in unmarshalling data bytes : ", err)
		return err
	}

	logging.GetLogger().Info("Length of data from snapshot : ", len(data))

	BuildCache(data)

	return nil
}

func BuildCache(data []sources.DataModel) {

	defer util.Duration(time.Now(), fmt.Sprintf("Build Cache for : %d ", len(data)))
	heatMapData := GetHeatMapDataInstance()
	heatMapData.Lock()
	defer heatMapData.Unlock()
	for i := 0; i < len(data); i++ {
		if i%1000 == 0 {
			logging.GetLogger().Info("i :", i)
		}
		heatMapData.cache[data[i].GetPrimaryKey()] = data[i]
	}
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
