package model

import "time"

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
