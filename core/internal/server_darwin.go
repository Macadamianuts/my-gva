package internal

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

// Initialization windows 自定义http配置
func (i *_server) Initialization(address string, engine *gin.Engine) server {
	engine.Static("/form-generator", "./resource/page")
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build, 并把dist文件夹放到resource文件夹下, 再打开下面1行注释
	// engine.Static("/admin", "./resource/dist")
	srv := endless.NewServer(address, engine)
	srv.ReadHeaderTimeout = 10 * time.Second
	srv.WriteTimeout = 10 * time.Second
	srv.MaxHeaderBytes = 1 << 20
	return srv
}
