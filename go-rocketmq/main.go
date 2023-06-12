package main

import (
	"go-rocketmq/consumer"
)

func main() {
	//consumer.Consumer()
	//producer.AsyncProducer()
	//producer.OrderlyProducer()
	consumer.OrderlyConsumer()
	//producer.DelayProducer()
	//producer.TransactionProducer()

}
