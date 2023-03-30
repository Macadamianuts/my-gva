package core

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gva-lbx/config"
)

type gormMysql struct {
	config *gorm.Config
}

func NewGormMysql(config *gorm.Config) Initialization {
	return &gormMysql{config: config}
}

// Initialization 初始化
func (c *gormMysql) Initialization(config *config.Gorm) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), c.config)
	if err != nil {
		return nil, errors.Wrap(err, "链接mysql失败!")
	}
	sql, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "获取mysql数据库连接失败!")
	}
	sql.SetMaxIdleConns(config.GetMaxIdleConesInt())
	sql.SetMaxOpenConns(config.GetMaxOpenConesInt())
	sql.SetConnMaxLifetime(config.GetConnMaxLifetimeDuration())
	sql.SetConnMaxIdleTime(config.GetConnMaxIdleTimeDuration())
	err = db.Use(c.Plugin(config.Replicas, config))
	if err != nil {
		return nil, errors.Wrap(err, "注册读写分离插件失败!")
	}
	return db, nil
}

// Plugin 读写分离插件
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (c *gormMysql) Plugin(replicas []*config.GormReplica, config *config.Gorm) gorm.Plugin {
	length := len(replicas)
	plugin := dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.New(mysql.Config{
			DSN:                       config.Dsn(),
			SkipInitializeWithVersion: true,
		})},
		Policy: dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})

	for j := 0; j < length; j++ {
		plugin.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.New(mysql.Config{
				DSN:                       replicas[j].MysqlDsn(),
				SkipInitializeWithVersion: true,
			})},
		}, replicas[j].DataInterfaces()...)
	}
	return plugin
}
