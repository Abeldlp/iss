package config

import (
	"os"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func InitializeRabbitConnection() {
	rabbitUrl := os.Getenv("RABBIT_MQ_URL")

	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	RabbitChannel = channel
}
