package core

import (
	"github.com/gin-gonic/gin"
	ginFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gva-lbx/app/router"
	"gva-lbx/core/internal"
	"gva-lbx/global"
	"gva-lbx/middleware"
)

var Engine = new(_engine)

type _engine struct{}

// Initialization 引擎初始化
func (c *_engine) Initialization() {
	engine := gin.Default()

	address := global.Config.System.Address()
	server := internal.Server.Initialization(address, engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(ginFiles.Handler))

	public := engine.Group("")
	{
		router.NewBaseRouter(public).Init()
	}
	private := engine.Group("")
	// private.Use(middleware.Jwt(), middleware.Casbin())
	private.Use(middleware.Operator())
	{
		// 路由引擎初始化
		router.NewApiRouter(private).Init()
		router.NewRoleRouter(private).Init()
		router.NewMenuRouter(private).Init()
		router.NewUserRouter(private).Init()
		router.NewDictionaryRouter(private).Init()
		router.NewDictionaryDetailRouter(private).Init()
		router.NewUploadRouter(private).Init()
	}

	Plugin.Initialization(public, private)

	// Router.Initialization(engine)
	zap.L().Error(server.ListenAndServe().Error())
}
