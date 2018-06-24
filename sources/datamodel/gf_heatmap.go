package datamodel

import (
	"context"
	"fmt"
	"github.com/op/go-logging"
	"gollow/config"
	"gollow/data"
	"strconv"
	"time"
)

var log = logging.MustGetLogger("gollow")

//HeatMapData is the DataModelName being loaded
type HeatMapData struct {
	ID             int64     `sql-col:"id" sql-key:"id" sql-insert:"false"`
	VehicleTypeID  int64     `sql-col:"vehicle_type_id"`
	Start          time.Time `sql-col:"start_time"`
	End            time.Time `sql-col:"end_time"`
	Geohash        string    `sql-col:"geohash"`
	UnmetDemand    int       `sql-col:"unmet_demand"`
	TotalDemand    int       `sql-col:"total_demand"`
	Imbalance      int       `sql-col:"imbalance"`
	Surge          float64   `sql-col:"surge"`
	Score          float64   `sql-col:"score"`
	Sources        string    `sql-col:"sources"`
	IsOverSupplied bool      `sql-col:"is_over_supplied"`
	IsShown        bool      `sql-col:"is_shown"`
	Version        string    `sql-col:"version"`
}

// GetID implements the data.Entity interface.
func (hd *HeatMapData) GetID() string {
	return strconv.FormatInt(hd.ID, 10)
}

// SetID implements the data.Entity interface.
func (hd *HeatMapData) SetID(ID string) {
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Error("Error parsing heatmap data id: %v", err)
		return
	}
	hd.ID = id
}

// NewEntity implements the data.Entity interface.
func (hd HeatMapData) NewEntity() data.Entity {
	return &HeatMapData{}
}

var HeatMapDataRef = &HeatMapData{}

func (heatmapData *HeatMapData) LoadAll() {

	start := time.Now()
	fmt.Println("Starting load of data")
	config.Init()
	query := "SELECT * FROM heatmap_data"
	entities, err := data.NewMySQLConnectionRef().NativeQueryRows(context.Background(), config.MySQLConfig, query, &HeatMapData{})

	if err != nil {
		log.Error("Error in fetching data from DB : ", err)
	}

	log.Info("Length of data", len(entities))
	log.Info("Total time taken : %s", time.Since(start))
}
