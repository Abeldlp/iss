package model

type SateliteLocation struct {
	Id         uint    `json:"id" gorm:"primaryKey;autoincrement"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude   float64 `json:"altitude"`
	Velocity   float64 `json:"velocity"`
	Visibility string  `json:"visibility"`
	Footprint  float64 `json:"footprint"`
	Timestamp  float64 `json:"timestamp"`
	Daynum     float64 `json:"daynum"`
	SolarLat   float64 `json:"solar_lat"`
	SolarLon   float64 `json:"solar_lon"`
	Units      string  `json:"units"`
}
