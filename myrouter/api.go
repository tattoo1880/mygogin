package myrouter

import (
	"github.com/gin-gonic/gin"
	"mygogin/config"
	"mygogin/controller"
	"mygogin/model"
	"mygogin/service"
)

func Initrouter() *gin.Engine {

	userepo := model.NewUserRepo(config.MySqlDB)
	userService := service.NewUserService(userepo)
	userController := controller.NewUserController(userService)

	eventRepo := model.NewEventRepo(config.MySqlDB)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	r := gin.Default()

	api := r.Group("/api")
	users := api.Group("/users")
	// users 子路由
	{
		users.POST("", userController.CreateUser)
		users.GET("/:id", userController.GetUserByID)
		users.GET("", userController.GetAllUsers)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
	// 事件相关的路由
	events := api.Group("/events")
	{
		events.POST("", eventController.CreateEvent)
		events.GET("/:id", eventController.GetEventByID)
		events.GET("", eventController.GetAllEvents)
		events.PUT("/:id", eventController.UpdateEvent)
		events.DELETE("/:id", eventController.DeleteEvent)
	}

	return r
}
