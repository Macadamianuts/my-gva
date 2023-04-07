package router

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/api"
	"gva-lbx/response"
)

type Upload struct {
	router *gin.RouterGroup
}

func NewUploadRouter(router *gin.RouterGroup) *Upload {
	return &Upload{router: router}
}

func (r *Upload) Init() {
	group := r.router.Group("file")
	{ // 不带日志中间件
		group.POST("upload", response.Handler()(api.UploadFile))
	}
}
