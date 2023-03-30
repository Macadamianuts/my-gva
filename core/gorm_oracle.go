package core

import (
	"gorm.io/gorm"
	"gva-lbx/config"
)

type gormOracle struct {
	config *gorm.Config
}

func NewGormOracle(config *gorm.Config) Initialization {
	return &gormOracle{config: config}
}

func (c *gormOracle) Initialization(config *config.Gorm) (*gorm.DB, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gormOracle) Plugin(replicas []*config.GormReplica, config *config.Gorm) gorm.Plugin {
	//TODO implement me
	panic("implement me")
}
