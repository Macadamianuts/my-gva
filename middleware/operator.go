package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func Operator() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "Operator", uint64(time.Now().Unix()))
		c.Request = c.Request.WithContext(ctx)
	}
}
