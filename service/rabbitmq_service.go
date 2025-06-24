//package service
//
//import (
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	. "mygogin/config"
//	"mygogin/model"
//)
//
//func ConsumeService(UserId string) {
//	// TODO
//
//	message, err := MyRabbitMQ.Consume(UserId)
//	if err != nil {
//		Logger.Error("消费消息失败")
//		return
//	}
//	Logger.Info("消费消息成功")
//	for d := range message {
//		Logger.Info("Received a message: ", zap.String("message", string(d.Body)))
//		var msg MsgBody
//		if err := json.Unmarshal(d.Body, &msg); err != nil {
//			Logger.Error("解析消息失败", zap.Error(err))
//			continue
//		}
//		Logger.Info("解析消息成功", zap.String("from_userid", msg.FromUserid), zap.String("msg", msg.Msg))
//		ClientConnLock.RLock()
//		conn, ok := ClientConnMap[UserId]
//		ClientConnLock.RUnlock()
//		if !ok {
//			Logger.Error("用户连接不存在", zap.String("user_id", UserId))
//			err := d.Nack(false, false)
//			if err != nil {
//				return
//			}
//			continue
//		}
//		if err := conn.WriteJSON(msg); err != nil {
//			Logger.Error("发送消息失败", zap.Error(err))
//			err := d.Nack(false, true) // requeue the message
//			if err != nil {
//				return
//			}
//			continue
//		}
//		// ! 手动确认消息
//		if err := d.Ack(false); err != nil {
//			Logger.Error("消息确认失败", zap.Error(err))
//			err := d.Nack(false, true) // requeue the message
//			if err != nil {
//				return
//			}
//			continue
//		}
//		Logger.Info("发送消息成功", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))
//
//		// TODO: 直接实例化 ChatMsgRepo 和 ChatMsgService
//		chatmsgrepo := model.NewChatMsgRepo(MySqlDB)
//		chatmsgservice := NewChatMsgService(chatmsgrepo)
//
//		chatmsg := &model.ChatMsg{
//			FromUserID: msg.FromUserid,
//			ToUserID:   UserId,
//			Msg:        msg.Msg,
//		}
//
//		if err := chatmsgservice.CreateChatMsg(chatmsg); err != nil {
//			Logger.Error("保存消息到数据库失败", zap.Error(err))
//			continue
//		}
//		Logger.Info("保存消息到数据库成功", zap.String("id", chatmsg.ID), zap.String("from_userid", chatmsg.FromUserID), zap.String("to_userid", chatmsg.ToUserID), zap.String("msg", chatmsg.Msg))
//
//		Logger.Info("消息确认成功", zap.String("message", string(d.Body)))
//		Logger.Info("消息处理完成", zap.String("to_user_id", UserId), zap.String("msg", msg.Msg))
//		Logger.Info("等待下一条消息...")
//
//	}
//}
//
//func ReciveRabbitMQ(ctx *gin.Context) {
//	type request struct {
//		ToUserId string  `json:"to_user_id"`
//		Content  MsgBody `json:"msg"`
//	}
//
//	var req request
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		Logger.Error("参数错误", zap.Error(err))
//		ctx.JSON(200, gin.H{
//			"code": 400,
//			"msg":  "参数错误",
//		})
//		return
//	}
//	Logger.Info("接收到消息", zap.String("给用户", req.ToUserId), zap.String("来自用户", req.Content.FromUserid), zap.String("msg", req.Content.Msg))
//	if err := MyRabbitMQ.Publish(req.ToUserId, req.Content); err != nil {
//		Logger.Error("发送消息失败")
//		ctx.JSON(200, gin.H{
//			"code": 400,
//			"msg":  "发送消息失败",
//		})
//		return
//	}
//	Logger.Info("发送消息成功")
//	ctx.JSON(200, gin.H{
//		"code": 200,
//		"msg":  "发送消息成功",
//	})
//}

package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "mygogin/config"
	"mygogin/model"
)

// NewChatMsgService 假设你在 service 包中定义了它
// func NewChatMsgService(repo *model.ChatMsgRepo) *ChatMsgService {
//     return &ChatMsgService{Repo: repo}
// }

func ConsumeService(UserId string) {
	message, err := MyRabbitMQ.Consume(UserId)
	if err != nil {
		Logger.Error("消费消息失败，无法获取消息通道", zap.Error(err))
		return // 如果无法获取消息通道，则直接退出
	}
	Logger.Info("消费消息成功，等待消息...", zap.String("consumer_id", UserId))

	// 建议：在这里初始化服务，而不是在循环内部重复创建
	chatmsgrepo := model.NewChatMsgRepo(MySqlDB)
	chatmsgservice := NewChatMsgService(chatmsgrepo) // 确保 NewChatMsgService 已定义

	for d := range message {
		// Log delivery tag for better traceability
		Logger.Info("收到新消息", zap.Uint64("delivery_tag", d.DeliveryTag), zap.String("message_body", string(d.Body)))

		var msg MsgBody
		// 1. 解析消息
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			Logger.Error("解析消息失败，跳过并 Nack (不重新入队)", zap.Error(err),
				zap.String("message_body", string(d.Body)),
				zap.Uint64("delivery_tag", d.DeliveryTag))
			// 如果消息格式错误，通常不应该重试，直接拒绝并丢弃（或发送到DLQ）
			if nackErr := d.Nack(false, false); nackErr != nil {
				Logger.Error("Nack 失败 (消息解析失败后)", zap.Error(nackErr), zap.Uint64("delivery_tag", d.DeliveryTag))
			}
			continue // 处理下一条消息
		}
		Logger.Info("消息解析成功", zap.String("from_userid", msg.FromUserid), zap.String("msg_content", msg.Msg))

		// 2. 检查 WebSocket 连接并发送
		ClientConnLock.RLock()
		conn, ok := ClientConnMap[UserId]
		ClientConnLock.RUnlock()
		if !ok {
			Logger.Error("用户 WebSocket 连接不存在，消息将重新入队", zap.String("user_id", UserId), zap.Uint64("delivery_tag", d.DeliveryTag))
			// 连接不存在可能是暂时的，重新入队等待用户重新连接
			if nackErr := d.Nack(false, true); nackErr != nil {
				Logger.Error("Nack 失败 (WebSocket 连接不存在后)", zap.Error(nackErr), zap.Uint64("delivery_tag", d.DeliveryTag))
			}
			continue
		}

		if err := conn.WriteJSON(msg); err != nil {
			Logger.Error("发送消息到 WebSocket 失败，消息将重新入队", zap.Error(err),
				zap.String("to_user_id", UserId),
				zap.String("msg_content", msg.Msg),
				zap.Uint64("delivery_tag", d.DeliveryTag))
			// WebSocket 写入失败也可能是暂时的，重新入队
			if nackErr := d.Nack(false, true); nackErr != nil {
				Logger.Error("Nack 失败 (WebSocket 发送失败后)", zap.Error(nackErr), zap.Uint64("delivery_tag", d.DeliveryTag))
			}
			continue
		}
		Logger.Info("发送消息到 WebSocket 成功", zap.String("to_user_id", UserId), zap.String("msg_content", msg.Msg))

		// 3. 保存消息到数据库
		chatmsg := &model.ChatMsg{
			FromUserID: msg.FromUserid,
			ToUserID:   UserId,
			Msg:        msg.Msg,
		}

		if err := chatmsgservice.CreateChatMsg(chatmsg); err != nil {
			Logger.Error("保存消息到数据库失败，消息将重新入队", zap.Error(err),
				zap.String("from_userid", chatmsg.FromUserID),
				zap.String("to_userid", chatmsg.ToUserID),
				zap.Uint64("delivery_tag", d.DeliveryTag))
			// 数据库错误通常是暂时的（例如连接池耗尽，死锁），重新入队以便重试
			if nackErr := d.Nack(false, true); nackErr != nil {
				Logger.Error("Nack 失败 (数据库保存失败后)", zap.Error(nackErr), zap.Uint64("delivery_tag", d.DeliveryTag))
			}
			continue
		}
		Logger.Info("保存消息到数据库成功", zap.String("id", chatmsg.ID),
			zap.String("from_userid", chatmsg.FromUserID),
			zap.String("to_userid", chatmsg.ToUserID),
			zap.String("msg_content", chatmsg.Msg),
			zap.Uint64("delivery_tag", d.DeliveryTag))

		// 4. 最终确认消息
		if err := d.Ack(false); err != nil {
			// Ack 失败通常表示 RabbitMQ 连接已断开或不稳定。
			// 此时不需要 Nack，消息会在连接恢复后由 RabbitMQ 自动处理（重新发送）。
			Logger.Error("消息最终确认失败 (RabbitMQ 连接可能已断开)", zap.Error(err), zap.Uint64("delivery_tag", d.DeliveryTag))
			// 这里不使用 continue，因为如果 Ack 失败，后续操作的可靠性已经无法保证。
			// 通常，如果 Ack 失败，这意味着 AMQP Channel 已经关闭或连接已断开，
			// 该 deliveryTag 对应的消息将自动重新入队 (re-queued)
			// 并且当前 ConsumeService 可能会在下一次尝试读取消息时失败并退出。
		} else {
			Logger.Info("消息已成功处理并确认", zap.Uint64("delivery_tag", d.DeliveryTag), zap.String("to_user_id", UserId))
		}
		Logger.Info("消息处理流程完成，等待下一条消息...", zap.Uint64("delivery_tag", d.DeliveryTag))
	}
}

// ReciveRabbitMQ 函数保持不变，因为它是生产方，不涉及消息确认逻辑
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
