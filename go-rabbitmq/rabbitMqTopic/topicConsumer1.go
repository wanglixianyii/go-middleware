package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	jay := rabbitMq.NewRabbitMQTopic("exchangeTopic", "Singer.*")
	jay.ConsumerTopic()
}
