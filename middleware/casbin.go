package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gva-lbx/component/jwt"
	"gva-lbx/global"
	"gva-lbx/response"
	"net/http"
)

// Casbin 权限分配
func Casbin() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, _ := jwt.NewClaimsByGinContext(context)
		path := context.Request.URL.Path // 获取请求的PATH
		subject := claims.Subject        // 获取用户的角色
		method := context.Request.Method // 获取请求方法
		err := global.Enforcer.LoadPolicy()
		if err != nil {
			zap.L().Error("casbin 加载策略失败!", zap.Error(err))
		}
		success, _ := global.Enforcer.Enforce(subject, path, method)
		if gin.Mode() == gin.DebugMode || success {
			context.Next()
		} else {
			context.JSON(http.StatusOK, response.Response{Code: http.StatusBadRequest, Data: gin.H{}, Message: "权限不足!"})
			context.Abort()
			return
		}
	}
}
