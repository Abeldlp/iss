package controller

import (
	"net/http"

	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/gin-gonic/gin"
)

func GetAllSateliteLocations(c *gin.Context) {
	var sateliteLocations []model.SateliteLocation
	if err := config.DB.Find(&sateliteLocations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, sateliteLocations)
}

func AggregateSateliteLocations(c *gin.Context) {
	type Result struct {
		Timezone string `json:"timezone"`
		Date     string `json:"date"`
		Hour     string `json:"hour"`
		Minutes  int    `json:"minutes"`
	}

	var results []Result

	timezones := c.QueryArray("timezone[]")
	dates := c.QueryArray("date[]")
	hours := c.QueryArray("hour[]")

	clause := config.DB.Model(&model.SateliteLocation{}).
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

	if err := clause.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, results)
}
