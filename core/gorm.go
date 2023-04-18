package core

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/service"
	"gva-lbx/global"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Initialization gorm 初始化 默认走mysql
func (c *_gorm) Initialization() error {
	config := c.Config(global.Config.Gorm.OtherConfig)
	var db *gorm.DB
	var err error
	switch global.Config.Gorm.Type {
	case "mysql":
		db, err = NewGormMysql(config).Initialization(global.Config.Gorm)
	case "oracle":
		db, err = NewGormOracle(config).Initialization(global.Config.Gorm)
	default:
		db, err = NewGormMysql(config).Initialization(global.Config.Gorm)
	}
	if err != nil {
		return errors.Wrap(err, "gorm 初始化失败")
	}
	global.Db = db
	c.AutoMigrate()
	dao.SetDefault(db)
	service.Jwt.Load()
	return nil
}

func (c *_gorm) AutoMigrate() {
	err := global.Db.AutoMigrate(
		// 系统模块表
		new(model.Api),
		new(model.User),
		new(model.Menu),
		new(model.Role),
		new(model.Dictionary),
		new(model.JwtBlacklist),
		new(model.MenuParameter),
		new(model.OperationRecord),
		new(model.DictionaryDetail),
	)
	if err != nil {
		zap.L().Error("结构体生成表失败!", zap.Error(err))
		return
	}
	zap.L().Info("结构体生成表成功!")
}
