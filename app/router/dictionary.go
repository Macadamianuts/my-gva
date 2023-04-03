package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type Dictionary struct {
	router *gin.RouterGroup
}

func NewDictionaryRouter(router *gin.RouterGroup) *Dictionary {
	return &Dictionary{router: router}
}

func (r *Dictionary) Init() {
	group := r.router.Group("dictionary")
	{ // 不带日志中间件
		group.POST("first", response.Handler()(api.Dictionary.First))
		group.POST("list", response.Handler()(api.Dictionary.List))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.Dictionary.Create))
		group.PUT("update", response.Handler()(api.Dictionary.Update))
		group.DELETE("delete", response.Handler()(api.Dictionary.Delete))
	}
}
