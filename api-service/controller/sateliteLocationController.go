package controller

import (
	"net/http"

	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/gin-gonic/gin"
)

func GetAllSateliteLocations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sateliteLocations []model.SateliteLocation

		if err := model.GetAllLocations(&sateliteLocations).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusOK, sateliteLocations)
	}
}

func AggregateSateliteLocations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []model.AggregatedResult

		timezones := c.QueryArray("timezone[]")
		dates := c.QueryArray("date[]")
		hours := c.QueryArray("hour[]")

		if err := model.GetAggregatedFilteredLocations(&results, timezones, dates, hours).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}
