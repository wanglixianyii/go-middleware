package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitMq.NewRabbitMQSimple("simple_queue")
	for i := 0; i < 10; i++ {
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i) + "---来自work模式")
		time.Sleep(1 * time.Second)
		fmt.Printf("work模式，共产生了%d条消息\n", i)
	}
}
