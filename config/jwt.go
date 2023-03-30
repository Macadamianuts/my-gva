package config

import (
	"gva-lbx/utils"
	"time"
)

func (x *Jwt) ExpiresAtDuration() time.Duration {
	parse, err := utils.Duration.Parse(x.ExpiresAt)
	if err != nil {
		return 0
	}
	return parse
}

func (x *Jwt) BufferAtDuration() time.Duration {
	parse, err := utils.Duration.Parse(x.BufferAt)
	if err != nil {
		return 0
	}
	return parse
}
