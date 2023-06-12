package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	two := rabbitMq.NewRabbitMQRouting("exchangeRouting", "two")
	two.ConsumerRouting()
}
