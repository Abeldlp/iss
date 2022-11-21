package controller

import (
	"net/http"

	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/gin-gonic/gin"
)

func GetAllSateliteLocations(c *gin.Context) {
	var sateliteLocations []model.SateliteLocation
	config.DB.Find(&sateliteLocations)

	c.JSON(http.StatusOK, sateliteLocations)
}
