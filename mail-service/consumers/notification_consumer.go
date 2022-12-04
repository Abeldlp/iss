package consumers

import (
	"log"

	"github.com/Abeldlp/fullinfo/mail-service/config"
	"github.com/Abeldlp/fullinfo/mail-service/utils"
)

func InitializeNotificationConsumer() {

	messages, err := config.RabbitChannel.Consume(
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

			to := "a.delapaz@presspage.com"

			utils.SendEmail(to)
		}
	}()

	<-openChannel
}
