package main

import (
	"github.com/Abeldlp/iss/scheduler-service/config"
	"github.com/Abeldlp/iss/scheduler-service/cron"
)

func main() {
	config.InitializeDatabase()
	config.InitializeEnvironmentVariables()
	config.InitializeMessageBroker()

	cron.ScheduleSaveSateliteLocation(60)
}
