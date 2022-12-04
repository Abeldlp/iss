package main

import (
	"github.com/Abeldlp/fullinfo/mail-service/config"
	"github.com/Abeldlp/fullinfo/mail-service/consumers"
)

func main() {
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()
	config.InitializeRabbitConnection()

	consumers.InitializeNotificationConsumer()
}
