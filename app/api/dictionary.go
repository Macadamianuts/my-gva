package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/response"
	"net/http"
)

var Dictionary = new(dictionary)

type dictionary struct{}

// Create
// @Tags     SystemDictionary
// @Summary  创建字典
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.DictionaryCreate    true "请求参数"
// @Success  200  {object} response.Response{data=any} "创建成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} ""创建失败!"
// @Router   /dictionary/create [post]
func (d *dictionary) Create(c *gin.Context) response.Response {
	var info request.DictionaryCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Dictionary.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorCreated}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessCreated}
}

// First
// @Tags     SystemDictionary
// @Summary  获取字典数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId                            true "请求参数"
// @Success  200  {object} response.Response{data=model.Dictionary} "获取单条数据成功!"
// @Failure  400  {object} response.Response{data=any}              "Bad Request"
// @Failure  500  {object} response.Response{data=any}              "获取单条数据失败!"
// @Router   /dictionary/first [post]
func (d *dictionary) First(c *gin.Context) response.Response {
	var info request.DictionaryFirst
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.Dictionary.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorFirst}
	}
	return response.Response{Code: http.StatusOK, Data: data, Message: common.SuccessFirst}
}

// Update
// @Tags     SystemDictionary
// @Summary  更新字典
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.DictionaryUpdate    true "请求参数"
// @Success  200  {object} response.Response{data=any} "更新成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "更新失败!"
// @Router   /dictionary/update [put]
func (d *dictionary) Update(c *gin.Context) response.Response {
	var info request.DictionaryUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Dictionary.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorUpdated}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// Delete
// @Tags     SystemDictionary
// @Summary  删除字典
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId               true "请求参数"
// @Success  200  {object} response.Response{data=any} "删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "删除失败!"
// @Router   /dictionary/delete [delete]
func (d *dictionary) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Dictionary.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorDeleted}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessDeleted}
}

// List
// @Tags     SystemDictionary
// @Summary  分页获取字典列表数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.DictionarySearch                   true "请求参数"
// @Success  200  {object} response.Response{data=[]model.Dictionary} "获取分页数据成功!"
// @Failure  400  {object} response.Response{data=any}                "Bad Request"
// @Failure  500  {object} response.Response{data=any}                "获取分页数据失败!"
// @Router   /dictionary/list [post]
func (d *dictionary) List(c *gin.Context) response.Response {
	var info request.DictionarySearch
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, count, err := service.Dictionary.List(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorList}
	}
	return response.Response{Code: http.StatusOK, Data: common.NewPageResult(list, count, info.PageInfo), Message: common.SuccessList}
}
