package config

import (
	"os"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func InitializeMessageBroker() {
	rabbitUrl := os.Getenv("RABBIT_MQ_URL")

	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare(
		"NotificationQueue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	RabbitConn = conn
	RabbitChannel = channel
}
