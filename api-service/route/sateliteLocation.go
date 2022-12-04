package route

import (
	"github.com/Abeldlp/iss/api-service/controller"
	"github.com/gin-gonic/gin"
)

func AppendSateliteLocationRoute(r *gin.Engine) {
	satelite := r.Group("/")

	{
		satelite.GET("/", controller.GetAllSateliteLocations())
		satelite.GET("/grouped", controller.AggregateSateliteLocations())
	}
}
