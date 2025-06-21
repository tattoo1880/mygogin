package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "mygogin/config"
)

func ConsumeService(UserId string) {
	// TODO

	message, err := MyRabbitMQ.Consume(UserId)
	if err != nil {
		Logger.Error("消费消息失败")
		return
	}
	Logger.Info("消费消息成功")
	for d := range message {
		Logger.Info("Received a message: ", zap.String("message", string(d.Body)))
		var msg MsgBody
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			Logger.Error("解析消息失败", zap.Error(err))
			continue
		}
		Logger.Info("解析消息成功", zap.String("from_userid", msg.FromUserid), zap.String("msg", msg.Msg))
		ClientConnLock.RLock()
		conn, ok := ClientConnMap[UserId]
		ClientConnLock.RUnlock()
		if !ok {
			Logger.Error("用户连接不存在", zap.String("user_id", UserId))
			continue
		}
		if err := conn.WriteJSON(msg); err != nil {
			Logger.Error("发送消息失败", zap.Error(err))
			continue
		}
		// ! 手动确认消息
		if err := d.Ack(false); err != nil {
			Logger.Error("消息确认失败", zap.Error(err))
			continue
		}
		Logger.Info("发送消息成功", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))

		Logger.Info("消息确认成功", zap.String("message", string(d.Body)))
		Logger.Info("消息处理完成", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))
		Logger.Info("等待下一条消息...")

	}
}

func ReciveRabbitMQ(ctx *gin.Context) {
	type request struct {
		ToUserId string  `json:"to_user_id"`
		Msg      MsgBody `json:"msg"`
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
	Logger.Info("接收到消息", zap.String("给用户", req.ToUserId), zap.String("来自用户", req.Msg.FromUserid), zap.String("msg", req.Msg.Msg))
	if err := MyRabbitMQ.Publish(req.ToUserId, req.Msg); err != nil {
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
