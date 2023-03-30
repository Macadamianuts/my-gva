package oss

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model"
	"gva-lbx/plugin/oss/global"
)

var Plugin = new(plugin)

type plugin struct{}

func (o *plugin) Menus() []model.Menu {
	return []model.Menu{}
}

func (o *plugin) Viper() any {
	return &global.Config
}

func (o *plugin) Name() string {
	return "oss"
}

func (o *plugin) Register(group *gin.RouterGroup) {

	// group.StaticFS(global.Config.LocalStorage.Path, http.Dir(global.Config.LocalStorage.Path)) // 为用户头像和文件提供静态地址
}

func (o *plugin) RouterPath() string {
	return ""
}
