package utils

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

var Duration = new(duration)

type duration struct{}

func (d *duration) Parse(str string) (time.Duration, error) {
	str = strings.TrimSpace(str)
	_duration, err := time.ParseDuration(str)
	if err != nil {
		if strings.Contains(str, "d") {
			index := strings.Index(str, "d")
			var hour int
			hour, err = strconv.Atoi(str[:index])
			if err != nil {
				return 0, errors.Wrap(err, "字符串转数字失败!")
			}
			_duration = time.Hour * 24 * time.Duration(hour)
			if str[index+1:] == "" {
				return _duration, nil
			}
			var __duration time.Duration
			__duration, err = time.ParseDuration(str[index+1:])
			if err != nil {
				return 0, errors.Wrap(err, "[utils][%d]转换time.Duration失败!")
			}
			return _duration + __duration, nil
		}
		var value int64
		value, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "字符串转数字失败!")
		}
		return time.Duration(value), err
	}
	return _duration, nil
}
