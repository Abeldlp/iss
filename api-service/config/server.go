package config

import "github.com/gin-gonic/gin"

func InitializeServer() *gin.Engine {
	r := gin.Default()

	//Middlewares here

	return r
}
