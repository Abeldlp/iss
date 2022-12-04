package controller

import (
	"github.com/Abeldlp/iss/scheduler-service/config"
	"github.com/Abeldlp/iss/scheduler-service/model"
)

func SaveSateliteLocation() model.SateliteLocation {

	var satelite model.SateliteLocation

	satelite.GetSateliteLocationData()
	satelite.GetSateliteCoordinatesTimzone()
	satelite.AddDateAndHour()

	config.DB.Create(&satelite)

	return satelite
}
