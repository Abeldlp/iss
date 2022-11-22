package controller

import (
	"github.com/Abeldlp/fullinfo/scheduler-service/config"
	"github.com/Abeldlp/fullinfo/scheduler-service/model"
)

func SaveSateliteLocation() {

	var satelite model.SateliteLocation

	satelite.GetSateliteLocationData()
	satelite.GetSateliteCoordinatesTimzone()

	config.DB.Create(&satelite)
}
