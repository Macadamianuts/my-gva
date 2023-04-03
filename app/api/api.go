package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/response"
	"net/http"
)

var Api = new(api)

type api struct{}

// Create
// @Tags     SystemApi
// @Summary  创建基础api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.ApiCreate  true "请求参数"
// @Success  200 {object}  response.Response{data=any} "创建成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} ""创建失败!"
// @Router   /api/create [post]
func (a *api) Create(c *gin.Context) response.Response {
	var info request.ApiCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Api.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusCreated, Message: common.SuccessCreated}
}

// First
// @Tags     SystemApi
// @Summary  根据id获取api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId                     true "请求参数"
// @Success  200  {object} response.Response{data=model.Api} "获取单条数据成功!"
// @Failure  400  {object} response.Response{data=any}       "Bad Request"
// @Failure  500  {object} response.Response{data=any}       "获取单条数据失败!"
// @Router   /api/first [post]
func (a *api) First(c *gin.Context) response.Response {
	var info common.GormId
	if err := c.ShouldBindQuery(&info); err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.Api.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: data, Message: common.SuccessFirst}
}

// Update
// @Tags     SystemApi
// @Summary  更新基础api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.ApiUpdate           true "请求参数"
// @Success  200  {object} response.Response{data=any} "更新成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "更新失败!"
// @Router   /api/update [put]
func (a *api) Update(c *gin.Context) response.Response {
	var info request.ApiUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Api.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// Delete
// @Tags     SystemApi
// @Summary  删除api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId               true "请求参数"
// @Success  200  {object} response.Response{data=any} "删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "删除失败!"
// @Router   /api/delete [delete]
func (a *api) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Api.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessDeleted}
}

// Deletes
// @Tags     SystemApi
// @Summary  批量删除api
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormIds              true "请求参数"
// @Success  200  {object} response.Response{data=any} "批量删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "批量删除失败!"
// @Router   /api/deletes [delete]
func (a *api) Deletes(c *gin.Context) response.Response {
	var info common.GormIds
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Api.Deletes(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessBatchDeleted}
}

// List
// @Tags     SystemApi
// @Summary  分页获取API列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.ApiSearch                   true "请求参数"
// @Success  200  {object} response.Response{data=[]model.Api} "获取分页数据成功!"
// @Failure  400  {object} response.Response{data=any}         "Bad Request"
// @Failure  500  {object} response.Response{data=any}         "获取分页数据失败!"
// @Router   /api/list [post]
func (a *api) List(c *gin.Context) response.Response {
	var info request.ApiSearch
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, count, err := service.Api.List(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: common.NewPageResult(list, count, info.PageInfo), Message: common.SuccessList}
}

// All
// @Tags     SystemApi
// @Summary  获取API列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{data=[]model.Api} "获取全部数据成功!"
// @Failure  400 {object} response.Response{data=any}         "Bad Request"
// @Failure  500 {object} response.Response{data=any}         "获取全部数据失败!"
// @Router   /api/all [post]
func (a *api) All(c *gin.Context) response.Response {
	list, err := service.Api.All(c.Request.Context())
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorAll}
	}
	return response.Response{Code: http.StatusOK, Data: gin.H{"apis": list}, Message: common.SuccessAll}
}
