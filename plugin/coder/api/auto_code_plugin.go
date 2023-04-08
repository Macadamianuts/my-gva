package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/plugin/coder/model/request"
	"gva-lbx/plugin/coder/service"
	"gva-lbx/response"
	"net/http"
)

var AutoCodePlugin = new(autoCodePlugin)

type autoCodePlugin struct{}

// Create
// @Tags     AutoCodePlugin
// @Summary  获取模版列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.AutoCodePluginCreate true "请求参数"
// @Success  200  {object} response.Response{data=any}  "创建插件模版成功!"
// @Failure  400  {object} response.Response{data=any}  "Bad Request"
// @Failure  500  {object} response.Response{data=any}  ""创建插件模版失败!"
// @Router   /coder/create [post]
func (a *autoCodePlugin) Create(c *gin.Context) response.Response {
	var info request.AutoCodePluginCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.AutoCodePlugin.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "创建插件模版失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "创建插件模版成功!"}
}
