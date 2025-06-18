package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/service"
)

type XListController struct {
	service service.XListService
}

func NewXListController(service service.XListService) *XListController {
	return &XListController{service: service}
}

func (c *XListController) GetList(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	list := c.service.GetList(id)
	if list == nil {
		ctx.JSON(404, gin.H{"error": "List not found"})
		return
	}

	ctx.JSON(200, gin.H{"list": list})
}
