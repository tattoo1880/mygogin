package config

import "go.uber.org/zap"

var Logger *zap.Logger

func Initlog() {
	Logger, _ = zap.NewDevelopment()
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {

		}
	}(Logger)
}
