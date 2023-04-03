package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/response"
)

type Base struct {
	router *gin.RouterGroup
}

func NewBaseRouter(router *gin.RouterGroup) *Base {
	return &Base{router: router}
}

func (r *Base) Init() {
	group := r.router.Group("base")
	{
		group.POST("login", response.Handler()(api.Base.Login))
		group.POST("captcha", response.Handler()(api.Base.Generate))
	}
}
