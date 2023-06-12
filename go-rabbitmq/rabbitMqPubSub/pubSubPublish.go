package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-rabbitmq/rabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitMq.NewRabbitMQPubSub("exchangePubSub")
	for i := 0; i < 100; i++ {
		err := rabbitmq.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
		time.Sleep(1 * time.Second)
	}
}
