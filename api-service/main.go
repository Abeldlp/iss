package main

import (
	"github.com/Abeldlp/iss/api-service/config"
	"github.com/Abeldlp/iss/api-service/route"
)

func main() {
	//Initial setup
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()
	r := config.InitializeServer()

	//Routes
	route.AppendSateliteLocationRoute(r)
	route.AppendUserRoute(r)

	r.Run(":8000")
}
