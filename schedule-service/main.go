package main

import (
	"github.com/Abeldlp/fullinfo/scheduler-service/config"
	"github.com/Abeldlp/fullinfo/scheduler-service/cron"
)

func main() {
	config.InitializeDatabase()
	config.InitializeEnvironmentVariables()
	config.InitializeMessageBroker()

	cron.ScheduleSaveSateliteLocation(60)
}
