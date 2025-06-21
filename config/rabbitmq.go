package config

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type MsgBody struct {
	FromUserid string `json:"from_userid"`
	Msg        string `json:"msg"`
}

var MyRabbitMQ *RabbitMQ

const ExchangeName = "chat-exchange"

func NewRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		Logger.Error("连接RabbitMQ失败", zap.Error(err))
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		Logger.Error("创建RabbitMQ通道失败", zap.Error(err))
		return
	}

	// 声明一个交换机
	err = ch.ExchangeDeclare(
		ExchangeName,
		"direct", // 交换机类型
		true,     // 是否持久化
		false,    // 是否自动删除
		false,    // 是否内部使用
		false,    // 是否排外
		nil,      // 额外参数
	)

	if err != nil {
		Logger.Error("声明交换机失败", zap.Error(err))
		return
	}

	MyRabbitMQ = &RabbitMQ{
		conn:    conn,
		channel: ch,
	}

}

func (r *RabbitMQ) Publish(routingKey string, body MsgBody) error {

	msg_body, err := json.Marshal(body)
	if err != nil {
		Logger.Error("消息序列化失败", zap.Error(err))
		return err
	}

	return r.channel.Publish(
		ExchangeName, // exchange 名称
		routingKey,   // routing key 是用户 ID
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			// 消息持久化
			DeliveryMode: amqp.Persistent, // 消息持久化
			ContentType:  "text/plain",
			Body:         msg_body, // 消息体
		},
	)
}

func (r *RabbitMQ) Consume(userId string) (<-chan amqp.Delivery, error) {
	queueName := "queue_" + userId
	q, error := r.channel.QueueDeclare(
		queueName, // 队列名称
		true,      // durable
		false,     // auto-deleted
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if error != nil {
		return nil, error
	}

	// 绑定用户的 routingKey 到交换机
	err := r.channel.QueueBind(
		q.Name,
		userId,       // routing key 是用户 ID
		ExchangeName, // 交换机名称
		false,
		nil, // 额外参数
	)
	if err != nil {
		return nil, err
	}

	// 开始消费消息
	return r.channel.Consume(
		q.Name, // 队列名称
		"",     // consumer tag
		// ! auto-ack 改成 false  // 是否自动确认消息
		// true,   // auto-ack
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		err := r.channel.Close()
		if err != nil {
			return
		}
	}
	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			return
		}
	}
}
