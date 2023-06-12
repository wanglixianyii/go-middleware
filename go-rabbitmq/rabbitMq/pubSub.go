package rabbitMq

import (
	"github.com/streadway/amqp"
	"log"
)

// NewRabbitMQPubSub 订阅模式
// 指定交换机名称， 而队列名称则为空，随机生产队列名称
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

// PublishPub 实现订阅模式下生产者代码
func (r *RabbitMQ) PublishPub(message string) error {
	r.Lock()
	defer r.Unlock()
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", //设置交换机的类型，在订阅模式下交换机的类型为广播类型
		true,
		false,
		false, //true表示这个参数exchange不可以被client用来推送信息，仅用来exchange和exchange之间的绑定
		false,
		nil,
	)
	r.failOnErr(err, "创建交换机异常")

	//发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		"",
		//如果为true，根据exchange类型和routeKey类型，如果无法找到符合条件的队列，name会把发送的信息返回给发送者
		false,
		//如果为true，当exchange发送到消息队列后发现队列上没有绑定的消费者,则会将消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "创建交换机异常")

	return nil
}

// ConsumePub 实现订阅模式下的消费端代码
func (r *RabbitMQ) ConsumePub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", //设置交换机的类型，在订阅模式下交换机的类型为广播类型
		true,
		false,
		false, //true表示这个参数exchange不可以被client用来推送信息，仅用来exchange和exchange之间的绑定
		false,
		nil,
	)
	r.failOnErr(err, "创建交换机异常")

	//申请队列,如果不存在会自动创建，存在跳过创建，保证队列存在，消息能发送到队列中
	q, err := r.channel.QueueDeclare(
		"", //随机生成队列名称
		//控制消息是否持久化，true开启
		true,
		//是否为自动删除
		false,
		//是否具有排他性
		true,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	r.failOnErr(err, "生产队列异常")

	//绑定队列到Exchange中
	err = r.channel.QueueBind(
		q.Name,
		"", //在订阅模式下这个参数要为空,
		r.Exchange,
		false,
		nil,
	)

	r.failOnErr(err, "绑定队列异常")

	//接收消息
	messages, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息
		//传递给同一个connection的消费者
		false,
		//是否为阻塞
		false,
		nil,
	)
	r.failOnErr(err, "接收消息异常")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range messages {
			//实现我们的处理逻辑函数
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages,To exit press CTRAL+C")
	<-forever
}
