package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gva-lbx/global"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type Jwt struct {
	SigningKey []byte
}

func NewJwt() *Jwt {
	return &Jwt{
		SigningKey: []byte(global.Config.Jwt.SigningKey),
	}
}

// Create 创建 token
func (j *Jwt) Create(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *Jwt) CreateTokenByOldToken(oldToken string, claims Claims) (string, error) {
	v, err, _ := global.SingleFlight.Do("Jwt:"+oldToken, func() (interface{}, error) {
		return j.Create(claims)
	})
	return v.(string), err
}

// Parse 解析 token
func (j *Jwt) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
