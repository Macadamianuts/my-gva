package captcha

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/global"
	"time"
)

var _ base64Captcha.Store = (*captchaRedis)(nil)

type captchaRedis struct {
	Key        string
	Context    context.Context
	Expiration time.Duration
}

func NewCaptchaRedis(ctx context.Context) base64Captcha.Store {
	return &captchaRedis{
		Key:        "CAPTCHA_",
		Context:    ctx,
		Expiration: time.Second * 180,
	}
}

func (r *captchaRedis) Set(id string, value string) error {
	err := global.Redis.Set(r.Context, r.Key+id, value, r.Expiration).Err()
	if err != nil {
		zap.L().Error("redis写入数据失败!", zap.String("key", r.Key+id), zap.String("value", value))
		return errors.Wrap(err, "redis 写入验证码数据失败!")
	}
	return nil
}

func (r *captchaRedis) Get(id string, clear bool) string {
	value, err := global.Redis.Get(r.Context, id).Result()
	if err == redis.Nil {
		zap.L().Error("从redis获取数据失败!", zap.String("key", r.Key+id))
		return ""
	}
	if clear {
		err = global.Redis.Del(r.Context, id).Err()
		if err != nil {
			zap.L().Error("删除redis中的数据失败!", zap.String("key", r.Key+id))
			return ""
		}
	}
	return value
}

func (r *captchaRedis) Verify(id, answer string, clear bool) bool {
	v := r.Get(r.Key+id, clear)
	return v == answer
}
