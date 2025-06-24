package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "mygogin/config"
	"mygogin/model"
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
			err := d.Nack(false, true)
			if err != nil {
				return
			}
			continue
		}
		if err := conn.WriteJSON(msg); err != nil {
			Logger.Error("发送消息失败", zap.Error(err))
			err := d.Nack(false, true) // requeue the message
			if err != nil {
				return
			}
			continue
		}
		// ! 手动确认消息
		if err := d.Ack(false); err != nil {
			Logger.Error("消息确认失败", zap.Error(err))
			err := d.Nack(false, true) // requeue the message
			if err != nil {
				return
			}
			continue
		}
		Logger.Info("发送消息成功", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))

		// TODO: 直接实例化 ChatMsgRepo 和 ChatMsgService
		chatmsgrepo := model.NewChatMsgRepo(MySqlDB)
		chatmsgservice := NewChatMsgService(chatmsgrepo)

		chatmsg := &model.ChatMsg{
			FromUserID: msg.FromUserid,
			ToUserID:   UserId,
			Msg:        msg.Msg,
		}

		if err := chatmsgservice.CreateChatMsg(chatmsg); err != nil {
			Logger.Error("保存消息到数据库失败", zap.Error(err))
			continue
		}
		Logger.Info("保存消息到数据库成功", zap.String("id", chatmsg.ID), zap.String("from_userid", chatmsg.FromUserID), zap.String("to_userid", chatmsg.ToUserID), zap.String("msg", chatmsg.Msg))

		Logger.Info("消息确认成功", zap.String("message", string(d.Body)))
		Logger.Info("消息处理完成", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))
		Logger.Info("等待下一条消息...")

	}
}

func ReciveRabbitMQ(ctx *gin.Context) {
	type request struct {
		ToUserId string  `json:"to_user_id"`
		Content  MsgBody `json:"msg"`
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		Logger.Error("参数错误", zap.Error(err))
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	Logger.Info("接收到消息", zap.String("给用户", req.ToUserId), zap.String("来自用户", req.Content.FromUserid), zap.String("msg", req.Content.Msg))
	if err := MyRabbitMQ.Publish(req.ToUserId, req.Content); err != nil {
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
