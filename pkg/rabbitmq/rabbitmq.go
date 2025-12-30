package rabbitmq

import (
	"kumiko/pkg/logger"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

// InitRabbitMQ 初始化RabbitMQ连接
func InitRabbitMQ() {
	url := viper.GetString("rabbitmq.url") // 例如 "amqp://guest:guest@localhost:5672/"
	var err error
	Conn, err = amqp.Dial(url)
	if err != nil {
		logger.StdError("RabbitMQ连接失败: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		logger.StdError("RabbitMQ打开通道失败: %v", err)
	}
}

// Publish 发送消息到指定队列
func Publish(queueName string, body []byte) error {
	_, err := Channel.QueueDeclare(
		queueName, // 队列名
		true,      // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // no-wait
		nil,       // 参数
	)
	if err != nil {
		return err
	}

	return Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

// Consume 从指定队列消费消息
func Consume(queueName string) (<-chan amqp.Delivery, error) {
	_, err := Channel.QueueDeclare(
		queueName, // 队列名
		true,      // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // no-wait
		nil,       // 参数
	)
	if err != nil {
		return nil, err
	}

	msgs, err := Channel.Consume(
		queueName, // 队列名
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	return msgs, err
}
