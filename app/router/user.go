package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type User struct {
	router *gin.RouterGroup
}

func NewUserRouter(router *gin.RouterGroup) *User {
	return &User{router: router}
}

func (r *User) Init() {
	group := r.router.Group("user")
	self := group.Group("self")
	{ // 不带日志中间件
		self.POST("first", response.Handler()(api.User.FirstSelf))
	}
	{ // 带日志中间件
		self.Use(middleware.OperationRecord())
		self.PUT("update", response.Handler()(api.User.UpdateSelf))
	}
	{ // 不带日志中间件
		group.POST("list", response.Handler()(api.User.List))
		group.POST("first", response.Handler()(api.User.First))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.User.Create))
		group.PUT("update", response.Handler()(api.User.Update))
		group.DELETE("delete", response.Handler()(api.User.Delete))
		group.DELETE("deletes", response.Handler()(api.User.Deletes))
		group.PATCH("setRole", response.Handler()(api.User.SetRole))
		group.PATCH("setRoles", response.Handler()(api.User.SetRoles))
		group.PATCH("changePassword", response.Handler()(api.User.ChangePassword))
		group.PATCH("resetPassword", response.Handler()(api.User.ResetPassword))
	}
}
