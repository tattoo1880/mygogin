package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "mygogin/config"
)

func ConsumeService(queueName string) {
	// TODO

	message, err := MyRabbitMQ.Consume(queueName)
	if err != nil {
		Logger.Error("消费消息失败")
		return
	}
	Logger.Info("消费消息成功")
	for d := range message {
		Logger.Info("Received a message: ", zap.String("message", string(d.Body)))
	}
}

func ReciveRabbitMQ(ctx *gin.Context) {
	type request struct {
		QueueName string `json:"queue_name"`
		Msg       string `json:"msg"`
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		Logger.Error("参数错误")
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	Logger.Info("接收到消息", zap.String("queue_name", req.QueueName), zap.String("msg", req.Msg))
	if err := MyRabbitMQ.Publish(req.QueueName, req.Msg); err != nil {
		Logger.Error("发送消息失败")
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "发送消息失败",
		})
		return
	}
	Logger.Info("发送消息成功")
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "发送消息成功",
	})
}
