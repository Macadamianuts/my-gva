package coder

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/global"
	"gva-lbx/plugin/coder/core"
	"gva-lbx/plugin/coder/model/dao"
	"gva-lbx/plugin/coder/router"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.RouterGroup) {
	core.Viper.Initialization()
	if global.Db != nil {
		dao.SetDefault(global.Db)
	}
	router.NewAutoCode(group).Init()
	router.NewAutoCodePlugin(group).Init()
}

func (p *plugin) RouterPath() string {
	return "coder"
}
