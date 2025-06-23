package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/model"
	"mygogin/service"
)

type ChatMsgController struct {
	service service.ChatMsgServiceInterface
}

func NewChatMsgController(service service.ChatMsgServiceInterface) *ChatMsgController {
	return &ChatMsgController{service: service}
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
	ctx.JSON(200, msgs)
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
