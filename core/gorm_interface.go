package core

import (
	"gorm.io/gorm"
	"gva-lbx/config"
)

type Initialization interface {
	Initialization(config *config.Gorm) (*gorm.DB, error)
	Plugin(replicas []*config.GormReplica, config *config.Gorm) gorm.Plugin
}
