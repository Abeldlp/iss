package route

import (
	"github.com/Abeldlp/iss/api-service/controller"
	"github.com/gin-gonic/gin"
)

func AppendUserRoute(r *gin.Engine) {
	satelite := r.Group("/users")

	{
		satelite.POST("", controller.SubscribeNewUser())
	}
}
