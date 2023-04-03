package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/app/model/response"
	"gva-lbx/common/captcha"
	"gva-lbx/global"
)

var Captcha = new(_captcha)

type _captcha struct{}

// Generate 生成验证码
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *_captcha) Generate(ctx context.Context, ip string) (entity *response.Captcha, err error) {
	entity = &response.Captcha{}
	if global.Redis != nil {
		global.CaptchaStore = captcha.NewCaptchaRedis(ctx)
	} else {
		global.CaptchaStore = base64Captcha.DefaultMemStore
	}
	entity.OpenCaptcha = s.ExplosionProof(ctx, ip)
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeightInt(), global.Config.Captcha.ImageWidthInt(), global.Config.Captcha.KeyLongInt(), 0.7, 80)
	entity.CaptchaId, entity.B64s, err = base64Captcha.NewCaptcha(driver, global.CaptchaStore).Generate()
	if err != nil {
		return nil, errors.Wrap(err, "生成验证码失败！")
	}
	entity.CaptchaLength = global.Config.Captcha.CaptchaLength()
	return entity, nil
}

// ExplosionProof 防爆破
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *_captcha) ExplosionProof(ctx context.Context, ip string) bool {
	if global.Redis != nil {
		times, err := global.Redis.Get(ctx, ip).Int64()
		if err == redis.Nil {
			err = global.Redis.Set(ctx, ip, 1, global.Config.Captcha.CacheTimeoutDuration()).Err()
			if err != nil {
				zap.L().Error("redis写入数据失败！", zap.Error(err))
				return false
			}
		}
		return global.Config.Captcha.OpenCaptcha(times)
	} else {
		value, ok := global.Cache.Get(ip)
		if !ok {
			global.Cache.Set(ip, 1, global.Config.Captcha.CacheTimeoutDuration())
		}
		times := value.(int)
		return global.Config.Captcha.OpenCaptcha(int64(times))
	}
}

// Explosion 验证码次数+1
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *_captcha) Explosion(ctx context.Context, ip string) {
	if global.Redis != nil {
		err := global.Redis.Incr(ctx, ip).Err()
		if err != nil {
			zap.L().Error("redis自增数据失败！", zap.Error(err))
		}
	} else {
		err := global.Cache.Increment(ip, 1)
		if err != nil {
			zap.L().Error("redis自增数据失败！", zap.Error(err))
		}
	}
}
