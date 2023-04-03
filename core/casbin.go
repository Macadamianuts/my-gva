package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gva-lbx/global"
)

var Casbin = new(_casbin)

type _casbin struct{}

func (c *_casbin) Initialization() {
	adapter, err := gormAdapter.NewAdapterByDB(global.Db)
	if err != nil {
		zap.L().Error("casbin适配器报错!", zap.Error(err))
		return
	}
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, _ := model.NewModelFromString(text)
	if global.Redis != nil { // TODO 如果redis存在
	}
	global.Enforcer, err = casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		zap.L().Error("casbin 初始化失败!", zap.Error(err))
		return
	}
	global.Enforcer.SetExpireTime(60 * 60)
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		zap.L().Error("casbin 加载数据失败!", zap.Error(err))
	}
}
