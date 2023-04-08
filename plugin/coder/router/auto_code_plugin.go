package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/plugin/coder/api"
	"gva-lbx/response"
)

type AutoCodePlugin struct {
	router *gin.RouterGroup
}

func NewAutoCodePlugin(router *gin.RouterGroup) *AutoCodePlugin {
	return &AutoCodePlugin{router: router}
}

func (r *AutoCodePlugin) Init() {
	group := r.router.Group("plugin")
	{
		group.POST("create", response.Handler()(api.AutoCodePlugin.Create))
	}
}
