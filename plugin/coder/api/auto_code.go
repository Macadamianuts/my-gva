package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gva-lbx/plugin/coder/global"
	"gva-lbx/plugin/coder/model/request"
	"gva-lbx/plugin/coder/service"
	"gva-lbx/response"
	"net/http"
	"net/url"
	"os"
)

var AutoCode = new(autoCode)

type autoCode struct{}

// Create
// @Tags     AutoCode
// @Summary  代码生成器生成代码
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.AutoCodeCreate      true "请求参数"
// @Success  200  {object} response.Response{data=any} "创建代码成功并移动文件成功!"
// @Failure  400 {object} response.Response{data=any} "Bad Request"
// @Failure  500 {object} response.Response{data=any} ""获取模版列表失败!"
// @Router   /coder/create [post]
func (a *autoCode) Create(c *gin.Context) response.Response {
	var info request.AutoCodeCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.AutoCode.Create(c.Request.Context(), info)
	defer func() {
		_ = os.Remove("./plugin.zip")
	}()
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	if info.AutoMoveFile {
		c.Writer.Header().Add("success", "true")
		c.Writer.Header().Add("msg", url.QueryEscape("创建代码成功并移动文件成功!"))
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "plugin.zip"))
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./plugin.zip")
	}
	return response.Response{Code: http.StatusOK}
}

// Preview
// @Tags     AutoCode
// @Summary  代码生成器预览代码
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.AutoCodeCreate      true "请求参数"
// @Success  200  {object} response.Response{data=any} "预览成功!"
// @Failure  400 {object} response.Response{data=any} "Bad Request"
// @Failure  500 {object} response.Response{data=any} ""预览失败!"
// @Router   /coder/preview [post]
func (a *autoCode) Preview(c *gin.Context) response.Response {
	var info request.AutoCodeCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.AutoCode.Preview(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "预览失败!"}
	}
	return response.Response{Code: http.StatusOK, Data: data, Message: "预览成功!"}
}

// Templates
// @Tags     AutoCode
// @Summary  获取模版列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{data=any} "获取模版列表成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} ""获取模版列表失败!"
// @Router   /coder/templates [post]
func (a *autoCode) Templates(c *gin.Context) response.Response {
	entries, err := os.ReadDir(global.Config.Server.TemplatePath())
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "获取模版列表失败!"}
	}
	length := len(entries)
	dirs := make([]string, 0, length)
	for i := 0; i < length; i++ {
		dirs = append(dirs, entries[i].Name())
	}
	return response.Response{Code: http.StatusOK, Data: dirs, Message: "获取模版列表成功!"}
}
