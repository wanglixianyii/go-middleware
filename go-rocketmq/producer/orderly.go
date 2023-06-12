package producer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"strconv"
)

func OrderlyProducer() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
	)
	if err != nil {
		panic("生成 producer 失败")
	}

	defer func(newProducer rocketmq.Producer) {
		err := newProducer.Shutdown()
		if err != nil {
			panic("关闭producer失败")
		}
	}(p)

	if err = p.Start(); err != nil {
		panic("启动 producer 失败")
	}

	// 能保证同一orderId下的消息是顺序的
	for i := 0; i < 3; i++ {
		orderId := strconv.Itoa(i)
		for j := 1; j < 5; j++ {
			msg := &primitive.Message{
				Topic: "OrderlyTopic",
				Body:  []byte("订单: " + orderId + " 步骤: " + strconv.Itoa(j)),
			}
			msg.WithShardingKey(orderId)                   // *关键 用于分片
			_, err = p.SendSync(context.Background(), msg) // 不能用单向
			if err != nil {
				log.Printf("send message err: %s", err)
				continue
			}
		}
	}
	log.Println("发布任务")
}
