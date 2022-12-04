package model

import (
	"time"

	"github.com/Abeldlp/iss/api-service/config"
	"gorm.io/gorm"
)

type SateliteLocation struct {
	IdPkey     uint      `json:"id_pkey"`
	Name       string    `json:"name"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Altitude   float64   `json:"altitude"`
	Velocity   float64   `json:"velocity"`
	Visibility string    `json:"visibility"`
	Footprint  float64   `json:"footprint"`
	Timestamp  int64     `json:"timestamp"`
	Daynum     float64   `json:"daynum"`
	SolarLat   float64   `json:"solar_lat"`
	SolarLon   float64   `json:"solar_lon"`
	Units      string    `json:"units"`
	Timezone   string    `json:"timezone"`
	Date       string    `json:"date"`
	Hour       string    `json:"hour"`
	DateISO    time.Time `json:"date_iso"`
}

type AggregatedResult struct {
	Timezone string `json:"timezone"`
	Date     string `json:"date"`
	Hour     string `json:"hour"`
	Minutes  int    `json:"minutes"`
}

func GetAllLocations(satelite *[]SateliteLocation) *gorm.DB {
	return config.DB.Find(&satelite)
}

func GetAggregatedFilteredLocations(
	results *[]AggregatedResult,
	timezones []string,
	dates []string,
	hours []string) *gorm.DB {

	clause := config.DB.Model(SateliteLocation{}).
		Select("timezone, date, hour, count(*) as minutes").
		Group("timezone, date, hour").
		Order("minutes desc")

	if len(timezones) > 0 {
		clause.Where("timezone IN ?", timezones)
	}

	if len(dates) > 0 {
		clause.Where("date IN ?", dates)
	}

	if len(hours) > 0 {
		clause.Where("hour IN ?", hours)
	}

	return clause.Find(&results)
}
