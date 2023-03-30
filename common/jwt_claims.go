package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gva-lbx/app/model/request"
	"gva-lbx/global"
	"strconv"
	"time"
)

type Claims struct {
	UserId   uint             `json:"user_id"`
	RoleId   uint             `json:"role_id"`
	BufferAt *jwt.NumericDate `json:"buf,omitempty"`
	jwt.RegisteredClaims
}

func NewClaims(claims request.Claims) Claims {
	now := time.Now()
	expiresAt := now.Add(global.Config.Jwt.ExpiresAtDuration())
	bufferAt := expiresAt.Add(-global.Config.Jwt.BufferAtDuration())
	return Claims{
		UserId:   claims.UserId,
		RoleId:   claims.RoleId,
		BufferAt: jwt.NewNumericDate(bufferAt), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.Config.Jwt.Issuer,      // 签名的发行者
			IssuedAt:  jwt.NewNumericDate(now),       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 过期时间 7天  配置文件
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(-1000))),
		},
	}
}

// NewClaimsByGinContext .
func NewClaimsByGinContext(ctx *gin.Context) (claims *Claims, err error) {
	value, exists := ctx.Get("claims")
	if !exists {
		return nil, errors.New("未启用jwt中间件!")
	}
	claims, exists = value.(*Claims)
	if !exists {
		return nil, errors.New("类型断言失败!")
	}
	return claims, nil
}

// GetUserId 获取用户id
func (c Claims) GetUserId() uint {
	id, _ := strconv.ParseUint(c.ID, 10, 64)
	return uint(id)
}

// GetRoleId 获取用户角色id
func (c Claims) GetRoleId() uint {
	id, _ := strconv.ParseUint(c.Subject, 10, 64)
	return uint(id)
}
