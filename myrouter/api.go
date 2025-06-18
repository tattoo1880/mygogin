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

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users/:id", userController.GetUserByID)
		api.GET("/users", userController.GetAllUsers)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	return r
}
