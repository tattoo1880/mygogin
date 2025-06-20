package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygogin/config"
	"mygogin/model"
	"mygogin/myrouter"
	"mygogin/service"
)

// todo 导入gogin

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.Initlog()
	config.Initdatabase()
	err := config.MySqlDB.AutoMigrate(&model.User{}, &model.Event{})

	if err != nil {
		config.Logger.Fatal("迁移数据库失败", zap.Error(err))
		panic("数据库迁移失败")
	}

	config.Logger.Info("数据库迁移成功")

	// 初始化RabbitMQ连接
	config.NewRabbitMQ()
	defer config.MyRabbitMQ.Close()
	go service.ConsumeService("2")
	go service.ConsumeService("1")

	// 初始化路由
	r := myrouter.Initrouter()
	err1 := r.Run("127.0.0.1:8080")
	if err1 != nil {

		config.Logger.Fatal("路由启动失败", zap.Error(err1))

	}
}
