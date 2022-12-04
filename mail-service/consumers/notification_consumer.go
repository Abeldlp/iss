package consumers

import (
	"log"

	"github.com/Abeldlp/fullinfo/mail-service/config"
	"github.com/Abeldlp/fullinfo/mail-service/models"
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

			var users []models.User
			config.DB.Where("timezone = ?", timezone).Find(&users)

			if len(users) > 0 {
				for _, user := range users {
					to := user.Email
					utils.SendEmail(to)
				}
			}
		}
	}()

	<-openChannel
}
