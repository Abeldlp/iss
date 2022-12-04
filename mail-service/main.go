package main

import (
	"github.com/Abeldlp/iss/mail-service/config"
	"github.com/Abeldlp/iss/mail-service/consumers"
)

func main() {
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()
	config.InitializeRabbitConnection()

	consumers.InitializeNotificationConsumer()
}
