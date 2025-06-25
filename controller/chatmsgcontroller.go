package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/model"
	"mygogin/service"
	"strconv"
)

type ChatMsgController struct {
	service     service.ChatMsgServiceInterface
	userservice service.UserService
}

func NewChatMsgController(service service.ChatMsgServiceInterface, userservice service.UserService) *ChatMsgController {
	return &ChatMsgController{service: service, userservice: userservice}
}

func (c *ChatMsgController) GetChatMsgsByFromUserId(ctx *gin.Context) {
	fromUserId := ctx.Param("fromUserId")

	msgs, err := c.service.FindChatMsgsByFromUserId(fromUserId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, msgs)
}

func (c *ChatMsgController) GetChatMsgsByToUserId(ctx *gin.Context) {
	toUserId := ctx.Param("toUserId")

	msgs, err := c.service.FindChatMsgsByToUserId(toUserId)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	//将 toUserId 转换为 int64 类型
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的用户ID"})
		return
	}

	var data []map[string]interface{}
	for _, msg := range msgs {
		myid, err := strconv.ParseInt(msg.ToUserID, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的用户ID"})
			return
		}
		fromid, err := strconv.ParseInt(msg.FromUserID, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的用户ID"})
			return
		}

		myname, err := c.userservice.GetUserByID(myid)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的用户ID"})
			return
		}
		fromname, err := c.userservice.GetUserByID(fromid)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的用户ID"})
			return
		}

		item := map[string]interface{}{
			"msg":            msg.Msg,
			"timestamp":      msg.TimeStamp,
			"from_user_name": fromname.Name,
			"to_user_name":   myname.Name,
			"id":             msg.ID,
		}

		data = append(data, item)
	}
	ctx.JSON(200, data)
}

func (c *ChatMsgController) CreateChatMsg(ctx *gin.Context) {
	var msg model.ChatMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateChatMsg(&msg); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "消息创建成功", "msg": msg})
}

func (c *ChatMsgController) GetAllChatMsgs(ctx *gin.Context) {
	msgs, err := c.service.FindAllChatMsgs()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, msgs)
}

func (c *ChatMsgController) DeleteChatMsg(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteChatMsg(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "消息删除成功"})
}
