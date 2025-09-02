package controller

import (
	"fmt"
	"mygogin/model"
	"mygogin/service"
	"regexp"

	"github.com/gin-gonic/gin"
)

type XListController struct {
	service           service.XListService
	downloaddbservice service.DownloadServiceInterface
}

func NewXListController(service service.XListService, downloadservice service.DownloadServiceInterface) *XListController {
	return &XListController{
		service:           service,
		downloaddbservice: downloadservice,
	}
}

func (c *XListController) GetList(ctx *gin.Context) {
	url := ctx.Query("url")
	re := regexp.MustCompile(`/(\d{10,})`)

	var id string

	match := re.FindStringSubmatch(url)
	if len(match) > 1 {

		id = match[1]

	} else {
		id = ""
	}

	fmt.Println("url:", id)
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	list := c.service.GetList(id)
	if list == "" {
		ctx.JSON(404, gin.H{"error": "List not found"})
		return
	}

	var donwload = &model.DonwLoad{
		DonwLoadUrl: list,
	}

	err := c.downloaddbservice.CreateDonwLoad(donwload)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}

	//ctx.JSON(200, gin.H{"url": list})
	ctx.Redirect(302, list)
}
