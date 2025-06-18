package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDB *gorm.DB

func Initdatabase() {
	dsn := "root:qwerty7788421@tcp(127.0.0.1:3306)/mygogin?charset=utf8mb4&parseTime=True&loc=Local"
	var dberror error
	MySqlDB, dberror = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dberror != nil {
		Logger.Error("连接数据库出现问题", zap.Error(dberror))
		panic("连接数据库失败")
	}
	Logger.Info("连接数据库成功")
}
