package core

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gva-lbx/config"
	"log"
	"os"
)

func (c *_gorm) Config(config *config.GormConfig) *gorm.Config {
	writer := NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags), config.LogZap)
	_logger := logger.New(writer, config.Config())
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NoLowerCase:   false,
		},
		Logger:                                   _logger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
