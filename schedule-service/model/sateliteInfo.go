package model

import (
	"fmt"
	"time"

	"github.com/Abeldlp/fullinfo/scheduler-service/util"
	"gorm.io/gorm"
)

type SateliteLocation struct {
	IdPkey     uint           `gorm:"primaryKey"`
	Name       string         `json:"name"`
	Latitude   float64        `json:"latitude"`
	Longitude  float64        `json:"longitude"`
	Altitude   float64        `json:"altitude"`
	Velocity   float64        `json:"velocity"`
	Visibility string         `json:"visibility"`
	Footprint  float64        `json:"footprint"`
	Timestamp  float64        `json:"timestamp"`
	Daynum     float64        `json:"daynum"`
	SolarLat   float64        `json:"solar_lat"`
	SolarLon   float64        `json:"solar_lon"`
	Units      string         `json:"units"`
	Timezone   string         `json:"timezone_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (satelite *SateliteLocation) GetSateliteLocationData() {
	util.GetJson("https://api.wheretheiss.at/v1/satellites/25544", &satelite)
}

func (satelite *SateliteLocation) GetSateliteCoordinatesTimzone() {
	var latitude string = fmt.Sprintf("%f", satelite.Altitude)
	var longitude string = fmt.Sprintf("%f", satelite.Latitude)

	util.GetJson("https://api.wheretheiss.at/v1/coordinates/"+latitude+","+longitude, &satelite)
}
