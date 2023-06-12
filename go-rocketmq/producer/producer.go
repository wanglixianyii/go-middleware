package producer

import (
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"sync"
	"time"
)

import (
	"context"
)

type Sensor struct {
	Id       int64  `json:"id"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

// AsyncProducer 当发送的消息很重要，且对响应时间非常敏感的时候采用async方式；
func AsyncProducer() {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}), producer.WithRetry(2))
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
		fmt.Printf("start producer error: %s", err.Error())
	}

	var msg Sensor

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		msg.Id = int64(i)
		msg.Type = "text"
		msg.Location = "北京"

		dataBuffer, err := json.Marshal(msg)
		if err != nil {
			return
		}
		message := primitive.NewMessage("SimpleTopic", dataBuffer)
		message.WithTag("order")
		message.WithKeys([]string{"1", "2"})

		err = p.SendAsync(context.Background(),
			func(ctx context.Context, result *primitive.SendResult, e error) {
				if e != nil {
					fmt.Printf("receive message error: %s\n", err)
				} else {
					fmt.Printf("send message success: result=%s\n", result.String())
				}
				wg.Done()
			}, message)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		}

	}
	wg.Wait()
}

// SimpleProducer 简单消息
func SimpleProducer() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"101.42.237.244:9876"}),
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

	var msg Sensor

	for i := 1; i < 5; i++ {
		msg.Id = int64(i)
		msg.Type = "text"
		msg.Location = "河北"

		dataBuffer, err := json.Marshal(msg)
		if err != nil {
			return
		}

		message := primitive.NewMessage("SimpleTopic", dataBuffer)

		message.WithTag("other")
		message.WithKeys([]string{"1", "2"})
		message.WithShardingKey("other_key")

		res, err := p.SendSync(context.Background(), message)
		if err != nil {
			fmt.Printf("发送失败: %s\n", err)
		}
		fmt.Printf(" 消息: %s发送成功 \n", res.String())
	}

	//nowStr := time.Now().Format("2006-01-02 15:04:05")

}

// DelayProducer 延迟消息
func DelayProducer() {

	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"101.42.237.244:9876"}), producer.WithRetry(2))

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
		panic("启动producer失败")
	}

	message := primitive.NewMessage("DelayTopic", []byte("一条延时消息"))

	// WithDelayTimeLevel 设置要消耗的消息延迟时间。参考延迟等级定义：1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	// 延迟等级从1开始，例如设置param level=1，则延迟时间为1s。
	// 这里使用的是延时30s发送
	message.WithDelayTimeLevel(4)
	res, err := p.SendSync(context.Background(), message)
	if err != nil {
		panic("消息发送失败" + err.Error())
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 消息: %s发送成功 \n", nowStr, res.String())
}

type TransactionListener struct{}

// ExecuteLocalTransaction 执行本地事务
// primitive.CommitMessageState : 提交
// primitive.RollbackMessageState : 回滚
// primitive.UnknownState : 触发会查函数 CheckLocalTransaction
func (o *TransactionListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行")
	time.Sleep(time.Second * 3)
	fmt.Println("执行成功")
	//执行逻辑无缘无故失败 代码异常 宕机
	//return primitive.CommitMessageState // 执行成功的

	return primitive.RollbackMessageState

	// return primitive.UnknownState
}

// CheckLocalTransaction 回查函数
// primitive.CommitMessageState : 提交
// primitive.RollbackMessageState : 回滚
// primitive.UnknownState : 触发会查函数 CheckLocalTransaction
func (o *TransactionListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq 的消息回查")
	time.Sleep(time.Second * 15)
	return primitive.CommitMessageState
}

func TransactionProducer() {
	p, err := rocketmq.NewTransactionProducer(
		&TransactionListener{},
		producer.WithNameServer([]string{"101.42.237.244:9876"}),
	)
	if err != nil {
		panic("生成 producer 失败")
	}

	defer func(newProducer rocketmq.TransactionProducer) {
		err := newProducer.Shutdown()
		if err != nil {
			panic("关闭producer失败")
		}
	}(p)

	if err = p.Start(); err != nil {
		panic("启动 producer 失败")
	}
	res, err := p.SendMessageInTransaction(context.Background(),
		primitive.NewMessage("TransTopic", []byte("this is transaction message")))
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 消息: %s发送成功 \n", nowStr, res.String())

	time.Sleep(time.Hour)
}
