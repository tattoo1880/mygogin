package myrouter

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mygogin/config"
	"mygogin/controller"
	"mygogin/model"
	"mygogin/myutils"
	"mygogin/service"
)

func Initrouter() *gin.Engine {

	userepo := model.NewUserRepo(config.MySqlDB)
	userService := service.NewUserService(userepo)
	userController := controller.NewUserController(userService)
	loginController := controller.NewLoginController(userService)

	eventRepo := model.NewEventRepo(config.MySqlDB)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	myutilsimpl := myutils.NewGetListUtils()
	myutilservice := service.NewXListService(myutilsimpl)
	xlistController := controller.NewXListController(myutilservice)

	myutilsximpl := myutils.NewXVideoDownload()
	myxvideoservice := service.NewXVideoService(myutilsximpl)
	xvideoController := controller.NewXVideoController(myxvideoservice)

	chatmsgrepo := model.NewChatMsgRepo(config.MySqlDB)
	chatmsgservice := service.NewChatMsgService(chatmsgrepo)
	chatmsgcontroller := controller.NewChatMsgController(chatmsgservice, userService)

	rabbitservice := service.NewConsumeService(chatmsgservice)
	websocketservice := service.NewWebSocketService(rabbitservice)

	r := gin.Default()
	// CORS配置
	r.Use(cors.Default())

	r.StaticFile("/mygogin/", "./static/index.html") // Matches /
	r.Static("/mygogin/static", "./static")
	r.Static("/mygogin/assets", "./static/assets") // Matches /assets

	api := r.Group("/mygogin/api")
	users := api.Group("/users")
	// users 子路由
	{
		users.POST("", userController.CreateUser)
		users.GET("/:id", userController.GetUserByID)
		users.GET("", userController.GetAllUsers)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
		users.POST("/login", loginController.Login) // 登录路由
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
	// RabbitMQ相关的路由
	rabbitmq := api.Group("/rabbitmq")
	{
		rabbitmq.POST("", rabbitservice.ReciveRabbitMQ)
		rabbitmq.GET("/ws", websocketservice.WebSocketHandler) // WebSocket连接
	}
	// Utils相关的路由
	myutilsroute := api.Group("/utils")
	{
		myutilsroute.GET("/:id", xlistController.GetList)
		myutilsroute.POST("/xvideo", xvideoController.GetXVideoUrl)

	}
	chatmsg := api.Group("/chatmsg")
	{
		chatmsg.GET("/from/:fromUserId", chatmsgcontroller.GetChatMsgsByFromUserId)
		chatmsg.GET("/to/:toUserId", chatmsgcontroller.GetChatMsgsByToUserId)
		chatmsg.POST("", chatmsgcontroller.CreateChatMsg)
		chatmsg.GET("", chatmsgcontroller.GetAllChatMsgs)
		chatmsg.DELETE("/:id", chatmsgcontroller.DeleteChatMsg)
	}

	return r
}
