package config

import (
	"gorm.io/gorm/logger"
	"gva-lbx/utils"
)

func (x *GormConfig) Config() logger.Config {
	level := logger.LogLevel(x.LogLevel)
	slowThreshold, _ := utils.Duration.Parse(x.SlowThreshold)
	return logger.Config{
		LogLevel:                  level,
		Colorful:                  x.Colorful,
		SlowThreshold:             slowThreshold,
		IgnoreRecordNotFoundError: x.IgnoreRecordNotFoundError,
	}
}
