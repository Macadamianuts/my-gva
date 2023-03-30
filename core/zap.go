package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gva-lbx/core/internal"
	"gva-lbx/global"
	"gva-lbx/utils"
	"os"
)

var Zap = new(_zap)

type _zap struct{}

func (c *_zap) Initialization() {
	ok, _ := utils.Directory.PathExists(global.Config.Zap.Director)
	if !ok { // 判断是否有 global.Config.Zap.Director 文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()         // 获取 zap 核心切片
	logger := zap.New(zapcore.NewTee(cores...)) // 初始化 zap.Logger
	if global.Config.Zap.ShowLine {             // 判断是否显示行
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger) // logger 注册到全局, 通过 zap.L() 调用日志组件
}
