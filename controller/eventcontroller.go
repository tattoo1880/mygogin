package controller

import (
	"github.com/gin-gonic/gin"
	"mygogin/model"
	"mygogin/service"
	"net/http"
	"strconv"
)

type EventController struct {
	service service.EventService
}

func NewEventController(service service.EventService) *EventController {
	return &EventController{service: service}
}

// 这里可以添加事件相关的处理方法，例如创建事件、获取事件等
func (c *EventController) CreateEvent(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.CreateEvent(&event); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create event"})
		return
	}
	ctx.JSON(http.StatusOK, event)

}

func (c *EventController) GetEventByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := c.service.GetEventByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Event not found"})
		return
	}
	ctx.JSON(200, event)
}

func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.service.GetAllEvents()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve events"})
		return
	}
	ctx.JSON(200, events)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}
	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	event.ID = idInt // 假设Event模型有ID字段
	if err := c.service.UpdateEvent(&event); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update event"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Event updated successfully"})
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}
	if err := c.service.DeleteEvent(idInt); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete event"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Event deleted successfully"})
}
