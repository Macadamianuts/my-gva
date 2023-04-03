package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/response"
	"net/http"
)

var Menu = new(menu)

type menu struct{}

// Create
// @Tags     SystemMenu
// @Summary  创建菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.MenuCreate          true "请求参数"
// @Success  200  {object} response.Response{data=any} "创建成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} ""创建失败!"
// @Router   /menu/create [post]
func (a *menu) Create(c *gin.Context) response.Response {
	var info request.MenuCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Menu.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: "创建成功!"}
}

// First
// @Tags     SystemMenu
// @Summary  创建菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId                      true "请求参数"
// @Success  200  {object} response.Response{data=model.Menu} "获取单条数据成功!"
// @Failure  400  {object} response.Response{data=any}        "Bad Request"
// @Failure  500  {object} response.Response{data=any}        "获取单条数据失败!"
// @Router   /menu/first [post]
func (a *menu) First(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.Menu.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: data}
}

// FindByRole
// @Tags     SystemMenu
// @Summary  根据角色获取菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.Role                          true "请求参数"
// @Success  200  {object} response.Response{data=[]model.Menu} "获取全部数据成功!"
// @Failure  400  {object} response.Response{data=any}          "Bad Request"
// @Failure  500  {object} response.Response{data=any}          "获取全部数据失败!"
// @Router   /menu/findByRole [post]
func (a *menu) FindByRole(c *gin.Context) response.Response {
	var info common.Role
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.Menu.FindByRole(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: data}
}

// Update
// @Tags     SystemMenu
// @Summary  更新菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.MenuUpdate          true "请求参数"
// @Success  200  {object} response.Response{data=any} "更新成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "更新失败!"
// @Router   /menu/update [put]
func (a *menu) Update(c *gin.Context) response.Response {
	var info request.MenuUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Menu.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK}
}

// Delete
// @Tags     SystemMenu
// @Summary  删除菜单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId               true "请求参数"
// @Success  200  {object} response.Response{data=any} "删除成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "删除失败!"
// @Router   /menu/delete [delete]
func (a *menu) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Menu.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK}
}

// List
// @Tags     SystemMenu
// @Summary  分页获取菜单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{data=[]model.Menu} "获取分页数据成功!"
// @Failure  400 {object} response.Response{data=any}          "Bad Request"
// @Failure  500 {object} response.Response{data=any}          "获取分页数据失败!"
// @Router   /menu/list [post]
func (a *menu) List(c *gin.Context) response.Response {
	list, err := service.Menu.List(c.Request.Context())
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: list}
}
