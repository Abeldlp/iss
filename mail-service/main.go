package main

import (
	"log"
	"os"

	"github.com/Abeldlp/fullinfo/mail-service/config"
	"github.com/Abeldlp/fullinfo/mail-service/utils"
	"github.com/streadway/amqp"
)

func main() {
	config.InitializeEnvironmentVariables()

	to := "a.delapaz@presspage.com"

	rabbitUrl := os.Getenv("RABBIT_MQ_URL")

	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	messages, err := channel.Consume(
		"NotificationQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	openChannel := make(chan bool)

	go func() {
		for message := range messages {
			timezone := string(message.Body)
			log.Printf(" > Received message: %s\n", timezone)

			utils.SendEmail(to)
		}
	}()

	<-openChannel
}
