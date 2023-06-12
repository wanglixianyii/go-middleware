package main

import "github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"

func main() {
	rabbitmq := rabbitMq.NewRabbitMQSimple("simple_queue")
	rabbitmq.ConsumeSimple()
}
