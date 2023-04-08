package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/plugin/coder/api"
	"gva-lbx/response"
)

type AutoCode struct {
	router *gin.RouterGroup
}

func NewAutoCode(router *gin.RouterGroup) *AutoCode {
	return &AutoCode{router: router}
}

func (r *AutoCode) Init() {
	group := r.router.Group("")
	{
		group.POST("create", response.Handler()(api.AutoCode.Create))
		group.POST("preview", response.Handler()(api.AutoCode.Preview))
		group.POST("templates", response.Handler()(api.AutoCode.Templates))
	}
}
