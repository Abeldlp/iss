package cron

import (
	"time"

	"github.com/Abeldlp/fullinfo/scheduler-service/config"
	"github.com/Abeldlp/fullinfo/scheduler-service/controller"
	"github.com/go-co-op/gocron"
	"github.com/streadway/amqp"
)

func ScheduleSaveSateliteLocation(amountOfSeconds int) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(amountOfSeconds).Seconds().Do(func() {
		satelite := controller.SaveSateliteLocation()

		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(satelite.Timezone),
		}

		config.RabbitChannel.Publish(
			"",
			"NotificationQueue",
			false,
			false,
			message,
		)

	})

	s.StartBlocking()
}
