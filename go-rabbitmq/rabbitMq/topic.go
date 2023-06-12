package rabbitMq

import (
	"github.com/streadway/amqp"
	"log"
)

//topic模式
//与routing模式不同的是这个exchange的kind是"topic"类型的。
//topic模式的特别是可以以通配符的形式来指定与之匹配的消费者。
//"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。

// 创建rabbitmq实例
func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routingKey)
}

// PublishTopic 实现topic模式下生产者代码
func (r *RabbitMQ) PublishTopic(message string) error {
	r.Lock()
	defer r.Unlock()

	err := r.channel.ExchangeDeclare(r.Exchange, "topic", true, false, false, false, nil)
	r.failOnErr(err, "topic模式尝试创建exchange失败")

	//发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		//如果为true，根据exchange类型和routeKey类型，如果无法找到符合条件的队列，name会把发送的信息返回给发送者
		false,
		//如果为true，当exchange发送到消息队列后发现队列上没有绑定的消费者,则会将消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "topic模式尝试push message失败")

	return nil
}

// ConsumerTopic 实现topic模式下的消费端代码
func (r *RabbitMQ) ConsumerTopic() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic", //设置交换机的类型，将广播模式改为direct模式
		true,
		false,
		false, //true表示这个参数exchange不可以被client用来推送信息，仅用来exchange和exchange之间的绑定
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建exchange失败。")

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
	r.failOnErr(err, "topic模式，消费者创建queue失败")

	//绑定队列到Exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key, //在订阅模式下这个参数要为空,
		r.Exchange,
		false,
		nil,
	)

	r.failOnErr(err, "topic模式，消费者绑定队列异常")

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
	r.failOnErr(err, "topic模式，消费者接收消息异常")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range messages {
			//实现我们的处理逻辑函数
			log.Printf("Received a message : %s", d.Body)

			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for messages,To exit press CTRAL+C")
	<-forever
}
