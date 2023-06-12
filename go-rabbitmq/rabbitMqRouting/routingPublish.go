package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := rabbitMq.NewRabbitMQRouting("exchangeRouting", "one")
	rabbitmq2 := rabbitMq.NewRabbitMQRouting("exchangeRouting", "two")
	rabbitmq3 := rabbitMq.NewRabbitMQRouting("exchangeRouting", "three")
	for i := 0; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		rabbitmq3.PublishRouting("路由模式three" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("在路由模式下，routingKey为one,为two,为three的都分别生产了%d条消息\n", i)
	}
}
