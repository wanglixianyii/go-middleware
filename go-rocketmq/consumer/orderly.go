package consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"strconv"
)

// 顺序消费
func OrderlyConsumer() {

	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"101.42.237.244:9876"}),
		consumer.WithGroupName("orderly_test"),          // 多个实例
		consumer.WithConsumerModel(consumer.Clustering), // 一条消息在一个组中只有一个consumer消费
		consumer.WithConsumerOrder(true),                // *顺序接收 必须
	)
	if err != nil {
		fmt.Println("创建消费者失败")
	}

	defer func(newPushConsumer rocketmq.PushConsumer) {
		err := newPushConsumer.Shutdown()
		if err != nil {
			panic("关闭consumer失败")
		}
	}(c)

	err = c.Subscribe("OrderlyTopic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, m := range msgs {
				log.Printf("%s Body: %s", getBlank(m.GetShardingKey()), m.Body)
			}
			return consumer.ConsumeRetryLater, nil
		})
	if err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	fmt.Println("开始读取消息 queue")
	select {}

}

func getBlank(s string) (rs string) {
	n, _ := strconv.Atoi(s)
	for i := 0; i < n; i++ {
		rs += `    `
	}
	return
}
