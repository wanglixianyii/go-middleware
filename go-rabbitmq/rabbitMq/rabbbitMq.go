package rabbitMq

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

// url格式：amqp://账号：密码@rabbitmq服务地址/vhost
const MQURL = "amqp://admin:admin@127.0.0.1:5672"

type RabbitMQ struct {
	sync.Mutex

	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 链接信息
	MqUrl string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqUrl:     MQURL,
	}
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl)
	if err != nil {
		rabbitMQ.failOnErr(err, "创建连接错误")
	}
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		rabbitMQ.failOnErr(err, "获取Channel失败")
	}
	return rabbitMQ
}

// 断开channel conn
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatal("%s:%s", msg, err)
	}
}
