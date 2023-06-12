package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	one := rabbitMq.NewRabbitMQRouting("exchangeRouting", "one")
	one.ConsumerRouting()
}
