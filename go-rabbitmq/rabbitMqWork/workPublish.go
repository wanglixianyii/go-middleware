package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"
	"strconv"
	"time"
)

//simple模式和work模式其实用的是一套逻辑代码，只是work模式是可以有多个消费者的，work模式起到一个负载均衡的作用。

func main() {
	rabbitmq := rabbitMq.NewRabbitMQSimple("" +
		"")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i) + "---来自work模式")
		time.Sleep(1 * time.Second)
		fmt.Printf("work模式，共产生了%d条消息\n", i)
	}
}
