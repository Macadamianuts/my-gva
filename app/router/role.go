package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type Role struct {
	router *gin.RouterGroup
}

func NewRoleRouter(router *gin.RouterGroup) *Role {
	return &Role{router: router}
}

func (r *Role) Init() {
	group := r.router.Group("role")
	{ // 不带日志中间件
		group.POST("first", response.Handler()(api.Role.First))
		group.POST("list", response.Handler()(api.Role.List))
		group.POST("menuTree", response.Handler()(api.Role.MenuTree))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.Role.Create))
		group.POST("copy", response.Handler()(api.Role.Copy))
		group.PUT("update", response.Handler()(api.Role.Update))
		group.DELETE("delete", response.Handler()(api.Role.Delete))
		group.POST("addMenus", response.Handler()(api.Role.AddMenus))
	}
}
