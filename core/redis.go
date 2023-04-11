package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gva-lbx/global"
)

var Redis = new(_redis)

type _redis struct{}

// Initialization 初始化
func (c *_redis) Initialization() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Username: global.Config.Redis.Username,
		Password: global.Config.Redis.Password, // no password set
		DB:       int(global.Config.Redis.DB),  // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("链接redis失败", zap.String("address", global.Config.Redis.Address), zap.Error(err))
		return
	}
	global.Redis = client
}
