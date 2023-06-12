package main

import (
	"github.com/wanglixianyii/go-middleware/go-rocketmq/consumer"
)

func main() {
	//consumer.Consumer()
	//producer.AsyncProducer()
	//producer.OrderlyProducer()
	consumer.OrderlyConsumer()
	//producer.DelayProducer()
	//producer.TransactionProducer()

}
