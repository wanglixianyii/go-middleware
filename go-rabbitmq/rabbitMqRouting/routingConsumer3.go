package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	three := rabbitMq.NewRabbitMQRouting("exchangeRouting", "three")
	three.ConsumerRouting()
}
