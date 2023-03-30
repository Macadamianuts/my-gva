package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/global"
	"gva-lbx/response"
	"net/http"
	"strconv"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusOK, response.Response{Message: "未登录或非法访问!"})
			c.Abort()
			return
		}
		if service.Jwt.IsBlacklist(token) { // 判断jwt是否在黑名单
			c.JSON(http.StatusOK, response.Response{Message: "您的帐户异地登陆或令牌失效!"})
			c.Abort()
			return
		}
		_jwt := common.NewJwt()
		claims, err := _jwt.Parse(token)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{Error: err})
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferAt.Unix() {
			expiresAt := global.Config.Jwt.ExpiresAtDuration()
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expiresAt))
			newToken, _ := _jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := _jwt.Parse(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if global.Config.System.UseMultipoint {
				var redisJwtToken string
				redisJwtToken, err = service.Jwt.GetRedisJWT(c.Request.Context(), newClaims.ID)
				if err != nil {
					zap.L().Error("get redis jwt failed", zap.Error(err))
					return
				}
				_ = service.Jwt.JsonInBlacklist(c.Request.Context(), redisJwtToken)      // 当之前的取成功时才进行拉黑操作
				_ = service.Jwt.SetRedisJWT(c.Request.Context(), newToken, newClaims.ID) // 无论如何都要记录当前的活跃状态
			}
		}
		c.Set("claims", claims)
	}
}
