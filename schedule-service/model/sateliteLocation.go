package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Abeldlp/iss/scheduler-service/util"
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
	Timestamp  int64          `json:"timestamp"`
	Daynum     float64        `json:"daynum"`
	SolarLat   float64        `json:"solar_lat"`
	SolarLon   float64        `json:"solar_lon"`
	Units      string         `json:"units"`
	Timezone   string         `json:"timezone_id"`
	Date       string         `json:"date"`
	Hour       string         `json:"hour"`
	DateISO    time.Time      `json:"date_iso"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (satelite *SateliteLocation) GetSateliteLocationData() {
	util.GetJson("https://api.wheretheiss.at/v1/satellites/25544", &satelite)
}

func (satelite *SateliteLocation) GetSateliteCoordinatesTimzone() {
	var latitude string = fmt.Sprintf("%f", satelite.Latitude)
	var longitude string = fmt.Sprintf("%f", satelite.Longitude)

	url := "https://api.wheretheiss.at/v1/coordinates/" + latitude + "," + longitude

	util.GetJson(url, &satelite)
}

func (satelite *SateliteLocation) AddDateAndHour() {
	timeStamp := strconv.FormatInt(satelite.Timestamp, 10)

	i, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	satelite.DateISO = tm
	satelite.Date = tm.Format("02-01-2006")
	satelite.Hour = tm.Format("15")
}
