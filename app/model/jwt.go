package model

import "gva-lbx/global"

// JwtBlacklist jwt黑名单
type JwtBlacklist struct {
	global.Model
	Jwt string `json:"jwt" gorm:"type:text;column:jwt;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "system_jwt_blacklist"
}
