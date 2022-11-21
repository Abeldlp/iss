package main

import (
	"github.com/Abeldlp/fullinfo/api-gateway/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/satelite/*proxyPath", handlers.ProxyTo("http://localhost:8000"))

	r.Run(":8080")
}
