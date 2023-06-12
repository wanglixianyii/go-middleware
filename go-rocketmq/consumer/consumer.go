package consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func Consumer() {

	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("simple_test"), // 多个实例
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

	err = c.Subscribe("SimpleTopic", consumer.MessageSelector{},
		func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, i := range msg {

				//nowStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("读取到一条普通消息,消息内容: %s \n", string(i.Body))

			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println("读取普通消息失败")
	}

	//err = c.Subscribe("DelayTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//	for _, msg := range msgs {
	//		nowStr := time.Now().Format("2006-01-02 15:04:05")
	//		fmt.Printf("%s 读取到一条延迟消息,消息内容: %s \n", nowStr, string(msg.Body))
	//	}
	//	return consumer.ConsumeSuccess, nil
	//})
	//if err != nil {
	//	fmt.Println("读取延迟消息失败")
	//}
	//
	//err = c.Subscribe("TransTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//	for _, msg := range msgs {
	//		nowStr := time.Now().Format("2006-01-02 15:04:05")
	//		fmt.Printf("%s 读取到一条事务消息,消息内容: %s \n", nowStr, string(msg.Body))
	//	}
	//	return consumer.ConsumeSuccess, nil
	//})
	//if err != nil {
	//	fmt.Println("读取延迟消息失败")
	//}

	if err = c.Start(); err != nil {
		panic("启动consumer失败")
	}
	//不能让主 goroutine 退出，不然就马上结束了
	time.Sleep(time.Hour)

}
