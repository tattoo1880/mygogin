package main

import (
	"go.uber.org/zap"
	"mygogin/config"
	"mygogin/model"
	"mygogin/myrouter"
)

// todo 导入gogin

func main() {
	config.Initlog()
	config.Initdatabase()
	err := config.MySqlDB.AutoMigrate(&model.User{}, &model.Event{})

	if err != nil {
		config.Logger.Fatal("迁移数据库失败", zap.Error(err))
		panic("数据库迁移失败")
	}

	config.Logger.Info("数据库迁移成功")

	r := myrouter.Initrouter()
	r.Static("/static", "./static")
	err1 := r.Run("127.0.0.1:8080")
	if err1 != nil {

		config.Logger.Fatal("路由启动失败", zap.Error(err1))

	}
}
