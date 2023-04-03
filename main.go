package main

import (
	"go.uber.org/zap"
	"gva-lbx/core"
)

func main() {
	core.Viper.Initialization()
	core.Zap.Initialization()
	err := core.Gorm.Initialization()
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	core.Redis.Initialization()
	core.Casbin.Initialization()
	// 引擎初始化
	core.Engine.Initialization()
}
