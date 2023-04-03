package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type Api struct {
	router *gin.RouterGroup
}

func NewApiRouter(router *gin.RouterGroup) *Api {
	return &Api{router: router}
}

func (r *Api) Init() {
	group := r.router.Group("api")
	{ // 不带日志中间件
		group.POST("all", response.Handler()(api.Api.All))
		group.POST("list", response.Handler()(api.Api.List))
		group.POST("first", response.Handler()(api.Api.First))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.Api.Create))
		group.PUT("update", response.Handler()(api.Api.Update))
		group.DELETE("delete", response.Handler()(api.Api.Delete))
		group.DELETE("deletes", response.Handler()(api.Api.Deletes))

	}
}
