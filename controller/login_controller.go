package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/model"
	"mygogin/service"
)

type LoginController struct {
	service service.UserService
}

func NewLoginController(service service.UserService) *LoginController {
	return &LoginController{service: service}
}

func (c *LoginController) Login(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "参数错误"})
		return
	}
	existingUser, err := c.service.Login(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "登录成功",
		"user":    existingUser,
	})

}
