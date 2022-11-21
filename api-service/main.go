package main

import (
	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/route"
)

func main() {
	//Initial setup
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()
	r := config.InitializeServer()

	//Routes
	route.AppendSateliteLocationRoute(r)

	r.Run(":8000")
}
