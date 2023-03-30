package service

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/global"
)

var Jwt = new(_jwt)

type _jwt struct{}

// JsonInBlacklist 拉黑列表
func (j *_jwt) JsonInBlacklist(ctx context.Context, jwt string) error {
	query := dao.Q.WithContext(ctx).JwtBlacklist
	err := query.Create(&model.JwtBlacklist{Jwt: jwt})
	if err != nil {
		return errors.Wrap(err, "拉黑失败！")
	}
	global.Cache.SetDefault(jwt, struct{}{})
	return nil
}

// IsBlacklist 判断jwt师傅在黑名单内部
func (j *_jwt) IsBlacklist(jwt string) bool {
	_, ok := global.Cache.Get(jwt)
	return ok
}

// GetRedisJWT 从redis取jwt
func (j *_jwt) GetRedisJWT(ctx context.Context, userName string) (redisJWT string, err error) {
	redisJWT, err = global.Redis.Get(ctx, userName).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并且设置过期时间
func (j *_jwt) SetRedisJWT(ctx context.Context, jwt string, userName string) error {
	err := global.Redis.Set(ctx, userName, jwt, global.Config.Jwt.ExpiresAtDuration()).Err()
	if err != nil {
		return errors.Wrap(err, "设置登录状态失败!")
	}
	return err
}

// Load jwt黑名单 加入到Cach中
func (j *_jwt) Load() {
	var data []string
	err := dao.Q.WithContext(context.Background()).JwtBlacklist.Pluck(dao.JwtBlacklist.Jwt, &data)
	if err != nil {
		zap.L().Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.Cache.SetDefault(data[i], struct{}{})
	}
}
