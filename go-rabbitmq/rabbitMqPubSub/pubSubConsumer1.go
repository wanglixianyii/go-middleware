package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	rabbitmq := rabbitMq.NewRabbitMQPubSub("exchangePubSub")
	rabbitmq.ConsumePub()
}
