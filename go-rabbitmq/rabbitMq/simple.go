package rabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 实现simple模式
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 实现简单模式下生产者代码
func (r *RabbitMQ) PublishSimple(message string) error {
	r.Lock()
	defer r.Unlock()
	//申请队列,如果不存在会自动创建，存在跳过创建，保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//控制消息是否持久化，true开启
		true,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		return err
	}
	//发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exchange类型和routekey类型，如果无法找到符合条件的队列，name会把发送的信息返回给发送者
		false,
		//如果为true，当exchange发送到消息队列后发现队列上没有绑定的消费者,则会将消息返还给发送者
		false,
		//发送信息
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		fmt.Println("生产简单消息失败", err)
		return err
	}
	return nil
}

// 实现简单模式下消费者代码
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列,如果不存在会自动创建，存在跳过创建，保证队列存在，消息能发送到队列中
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//控制消息是否持久化，true开启
		true,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//消费者流控，防止暴库
	err = r.channel.Qos(
		//每次只接受一个消息进行消费
		1,
		//服务器传递的最大容量（以8位字节为单位）
		0,
		//true对全局可用，false只对当前channel可用
		false,
	)
	if err != nil {
		fmt.Println("消费者流控失败")
		return
	}

	//接收消息
	messages, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		//是否自动应答
		false,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息
		//传递给同一个connection的消费者
		false,
		//是否为阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range messages {
			//实现我们的处理逻辑函数
			log.Printf("Received a message : %s", d.Body)

			// 进行业务操作。。。。。

			//message := &datamodels.Message{}
			//err := json.Unmarshal([]byte(d.Body), message)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//插入订单
			//_, err = orderService.InsertOrderByMessage(message)
			//if err != nil {
			//	fmt.Println(err)
			//}
			////扣除数量
			//err = productService.SubNumberOne(message.ProductID)
			//if err != nil {
			//	fmt.Println(err)
			//}

			//如果为true表示确认所有未确认的消息，false为当前消息
			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for messages,To exit press CTRAL+C")
	<-forever
}
