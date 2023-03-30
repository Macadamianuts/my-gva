package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Err     string      `json:"err,omitempty"`
	Code    int         `json:"-"`
	Data    interface{} `json:"data"`
	Error   error       `json:"-"`
	Message string      `json:"message"`
}

type handler func(c *gin.Context) Response

func Handler() func(handler handler) gin.HandlerFunc {
	return func(handler handler) gin.HandlerFunc {
		return func(context *gin.Context) {
			response := handler(context)
			if response.Data == nil {
				response.Data = struct{}{}
			}
			if response.Error != nil {
				if response.Message == "" {
					response.Message = response.Error.Error()
				}
				if gin.Mode() == gin.DebugMode { // 只有debug模式下才输出err
					response.Err = response.Error.Error()
				}
				zap.L().Error(fmt.Sprintf("%+v", response.Error))
			}
			if response.Message == "" {
				response.Message = http.StatusText(response.Code)
			}
			context.JSON(response.Code, response)
		}
	}
}
