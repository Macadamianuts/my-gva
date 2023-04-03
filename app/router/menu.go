package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type Menu struct {
	router *gin.RouterGroup
}

func NewMenuRouter(router *gin.RouterGroup) *Menu {
	return &Menu{router: router}
}

func (r *Menu) Init() {
	group := r.router.Group("menu")
	{ // 不带日志中间件
		group.POST("first", response.Handler()(api.Menu.First))
		group.POST("list", response.Handler()(api.Menu.List))
		group.POST("findByRole", response.Handler()(api.Menu.FindByRole))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.Menu.Create))
		group.PUT("update", response.Handler()(api.Menu.Update))
		group.DELETE("delete", response.Handler()(api.Menu.Delete))

	}
}
