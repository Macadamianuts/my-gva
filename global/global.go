package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"gva-lbx/config"
)

var (
	Db           *gorm.DB
	Redis        *redis.Client
	Viper        *viper.Viper
	Config       = new(config.Config)
	SingleFlight = new(singleflight.Group)
)
