package config

import (
	"go.uber.org/zap"
	"gva-lbx/utils"
	"time"
)

func (x *Captcha) KeyLongInt() int {
	return int(x.KeyLong)
}

func (x *Captcha) CaptchaLength() int {
	return x.KeyLongInt()
}

func (x *Captcha) ImageWidthInt() int {
	return int(x.ImageWidth)
}

func (x *Captcha) ImageHeightInt() int {
	return int(x.ImageHeight)
}

func (x *Captcha) OpenCaptcha(times int64) bool {
	if x.ExplosionProof == 0 || x.ExplosionProof < times {
		return true
	}
	return false
}

func (x *Captcha) CacheTimeoutDuration() time.Duration {
	parse, err := utils.Duration.Parse(x.CacheTimeout)
	if err != nil {
		zap.L().Error("获取验证码配置文件的缓存超时时间失败！", zap.Error(err))
		return 0
	}
	return parse
}
