package middleware

import "github.com/gin-gonic/gin"

func OperationRecord() gin.HandlerFunc {
	Operator()
	return func(context *gin.Context) {
		context.Next()
	}
}
