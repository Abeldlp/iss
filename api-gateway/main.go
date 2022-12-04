package main

import (
	"os"

	"github.com/Abeldlp/iss/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("PROD_ENV")

	if env == "" {
		godotenv.Load(".env")
	}

	r := gin.Default()

	r.Any("/*proxyPath", handlers.ProxyTo(os.Getenv("API_SERVICE_URL")))

	r.Run(":8080")
}
