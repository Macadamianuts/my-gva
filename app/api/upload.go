package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/service"
	"gva-lbx/response"
	"net/http"
)

// UploadFile 文件上传api
func UploadFile(c *gin.Context) response.Response {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: "文件超出限制大小", Error: err}
	}
	file, name, err := service.UploadFile(c.Request.Context(), header)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Message: "文件传输错误", Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: "上传成功", Data: gin.H{
		"path":     file,
		"filename": name,
	}}
}
