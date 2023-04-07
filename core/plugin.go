package core

import (
	"context"
	"gva-lbx/app/model/dao"
	"gva-lbx/core/internal"
	"gva-lbx/plugin/oss"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Plugin = new(_plugin)

type _plugin struct{}

// Initialization 初始化插件
func (p *_plugin) Initialization(public, private *gin.RouterGroup) {
	p.PluginInit(oss.Plugin, public)
}

// PluginInit 插件初始化
func (p *_plugin) PluginInit(plugin internal.Plugin, group *gin.RouterGroup) {
	plus, ok := plugin.(internal.PluginPlus)
	if !ok {
		return
	}
	Viper.PluginConfig(plus.Viper(), plus.Name())
	menus := plus.Menus()
	for i := 0; i < len(menus); i++ {
		query := dao.Q.WithContext(context.Background()).Menu
		entity, err := query.Where(dao.Menu.Name.Eq(menus[i].Name)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) { // 数据库无记录
			err = query.Create(&menus[i])
			if err != nil {
				zap.L().Error("[PluginPlus]菜单注册失败!", zap.String("Name", menus[i].Name), zap.Error(err))
			}
		} else { // 数据库有记录则更新
			_, err = query.Where(dao.Menu.ID.Eq(entity.ID)).Updates(menus[i])
			if err != nil {
				zap.L().Error("[PluginPlus]菜单更新失败!", zap.String("Name", menus[i].Name), zap.Error(err))
			}
		}
	}

	group = group.Group(plugin.RouterPath())
	plugin.Register(group)
}
