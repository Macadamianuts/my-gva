package test

import (
	"gitea.madlocker.cn/SliverHorn/gin-vue-admin/global"
	"gitea.madlocker.cn/SliverHorn/gin-vue-admin/plugin/{{.Name}}/model/dao"
	"github.com/gin-gonic/gin"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.RouterGroup) {
	p.AutoMigrate()
}

func (p *plugin) RouterPath() string {
	return ""
}

// AutoMigrate 结构体生成表
func (p *plugin) AutoMigrate() {
	if global.Db != nil {
        dao.SetDefault(global.Db)
        err := global.Db.AutoMigrate()
		if err != nil {
			zap.L().Error("结构体生成表失败!", zap.Error(err))
			return
		}
	}
}