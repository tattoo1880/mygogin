package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/service"
)

type XVideoController struct {
	service service.XVideoService
}

type XvideoRequest struct {
	Url string `json:"url" binding:"required"`
}

func NewXVideoController(service service.XVideoService) *XVideoController {
	return &XVideoController{service: service}
}

func (controller *XVideoController) GetXVideoUrl(c *gin.Context) {

	var xvideoRequest XvideoRequest

	if err := c.ShouldBindJSON(&xvideoRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	downloadUrl, err := controller.service.GetUrl(xvideoRequest.Url)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"download_url": downloadUrl})

}
