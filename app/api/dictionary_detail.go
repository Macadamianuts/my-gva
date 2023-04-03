package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/response"
	"net/http"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

// Create
// @Tags     SystemDictionaryDetail
// @Summary  创建字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.DictionaryDetailCreate true "请求参数"
// @Success  200  {object} response.Response{data=any}    "创建成功!"
// @Failure  400  {object} response.Response{data=any}    "Bad Request"
// @Failure  500  {object} response.Response{data=any}    ""创建失败!"
// @Router   /dictionaryDetail/create [post]
func (d *dictionaryDetail) Create(c *gin.Context) response.Response {
	var info request.DictionaryDetailCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.DictionaryDetail.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorCreated}
	}
	return response.Response{Code: http.StatusCreated, Message: common.SuccessCreated}
}

// First
// @Tags     SystemDictionaryDetail
// @Summary  获取字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId                                  true "请求参数"
// @Success  200  {object} response.Response{data=model.DictionaryDetail} "获取单条数据成功!"
// @Failure  400  {object} response.Response{data=any}                    "Bad Request"
// @Failure  500  {object} response.Response{data=any}                    "获取单条数据失败!"
// @Router   /dictionaryDetail/first [post]
func (d *dictionaryDetail) First(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.DictionaryDetail.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorFirst}
	}
	return response.Response{Code: http.StatusOK, Data: data, Message: common.SuccessFirst}
}

// Update
// @Tags     SystemDictionaryDetail
// @Summary  更新字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.DictionaryDetailUpdate true "请求参数"
// @Success  200  {object} response.Response{data=any}    "更新成功!"
// @Failure  400  {object} response.Response{data=any}    "Bad Request"
// @Failure  500  {object} response.Response{data=any}    "更新失败!"
// @Router   /dictionaryDetail/update [put]
func (d *dictionaryDetail) Update(c *gin.Context) response.Response {
	var info request.DictionaryDetailUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.DictionaryDetail.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorUpdated}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// Delete
// @Tags     SystemDictionaryDetail
// @Summary  删除字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId               true "请求参数"
// @Success  200  {object} response.Response{data=any} "删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "删除失败!"
// @Router   /dictionaryDetail/delete [delete]
func (d *dictionaryDetail) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.DictionaryDetail.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorDeleted}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessDeleted}
}

// Deletes
// @Tags     SystemDictionaryDetail
// @Summary  批量删除字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormIds              true "请求参数"
// @Success  200  {object} response.Response{data=any} "批量删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "批量删除失败!"
// @Router   /dictionaryDetail/deletes [delete]
func (d *dictionaryDetail) Deletes(c *gin.Context) response.Response {
	var info common.GormIds
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.DictionaryDetail.Deletes(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorBatchDeleted}
	}
	return response.Response{Code: http.StatusOK, Message: common.ErrorBatchDeleted}
}

// List
// @Tags     SystemDictionaryDetail
// @Summary  删除字典详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormIds                                   true "请求参数"
// @Success  200  {object} response.Response{data=[]model.DictionaryDetail} "获取分页数据成功!"
// @Failure  400  {object} response.Response{data=any}                      "Bad Request"
// @Failure  500  {object} response.Response{data=any}                      "获取分页数据失败!"
// @Router   /dictionaryDetail/list [post]
func (d *dictionaryDetail) List(c *gin.Context) response.Response {
	var info request.DictionaryDetailSearch
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, count, err := service.DictionaryDetail.List(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorList}
	}
	return response.Response{Code: http.StatusOK, Data: common.NewPageResult(list, count, info.PageInfo), Message: common.SuccessList}
}
