package internal

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model"
)

// Plugin 插件模式接口化
type Plugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)

	// RouterPath 用户返回注册路由
	RouterPath() string
}

type PluginPlus interface {
	// Menus 菜单注册
	Menus() []model.Menu
	// Viper 初始化配置文件
	Viper() any
	// Name 插件名称
	Name() string
}
