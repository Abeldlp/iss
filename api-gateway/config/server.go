package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	return r
}
