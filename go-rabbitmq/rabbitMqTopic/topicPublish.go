package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"
	"strconv"
	"time"
)

func main() {
	one := rabbitMq.NewRabbitMQTopic("exchangeTopic", "Singer.Jay")
	two := rabbitMq.NewRabbitMQTopic("exchangeTopic", "All")
	for i := 0; i < 100; i++ {
		one.PublishTopic("topic模式，Jay," + strconv.Itoa(i))
		two.PublishTopic("topic模式，All," + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("topic模式的消息%v \n", i)
	}
}
