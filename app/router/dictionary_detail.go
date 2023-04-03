package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/middleware"
	"gva-lbx/response"
)

type DictionaryDetail struct {
	router *gin.RouterGroup
}

func NewDictionaryDetailRouter(router *gin.RouterGroup) *DictionaryDetail {
	return &DictionaryDetail{router: router}
}

func (r *DictionaryDetail) Init() {
	group := r.router.Group("dictionaryDetail")
	{ // 不带日志中间件
		group.POST("first", response.Handler()(api.DictionaryDetail.First))
		group.POST("list", response.Handler()(api.DictionaryDetail.List))
	}
	{ // 带日志中间件
		group.Use(middleware.OperationRecord())
		group.POST("create", response.Handler()(api.DictionaryDetail.Create))
		group.PUT("update", response.Handler()(api.DictionaryDetail.Update))
		group.DELETE("delete", response.Handler()(api.DictionaryDetail.Delete))
		group.DELETE("deletes", response.Handler()(api.DictionaryDetail.Deletes))
	}
}
