package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/response"
	"net/http"
)

var Role = new(role)

type role struct{}

// Create
// @Tags    SystemRole
// @Summary 创建角色
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleCreate          true "请求参数"
// @Success 200  {object} response.Response{data=any} "创建成功!"
// @Failure 400  {object} response.Response{data=any}        "Bad Request"
// @Failure 500  {object} response.Response{data=any} ""创建失败!"
// @Router  /role/create [post]
func (a *role) Create(c *gin.Context) response.Response {
	var info request.RoleCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err, Message: common.ErrorCreated}
	}
	err = service.Role.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorCreated}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessCreated}
}

// First
// @Tags    SystemRole
// @Summary 根据 id 获取角色数据
// @accept  application/json
// @Produce application/json
// @Param   data body     common.GormId                      true "请求参数"
// @Success 200  {object} response.Response{data=model.Role} "获取单条数据成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any}        "获取单条数据失败!"
// @Router  /role/first [post]
func (a *role) First(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.Role.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorFirst}
	}
	return response.Response{Code: http.StatusOK, Data: data, Message: common.SuccessFirst}
}

// Update
// @Tags    SystemRole
// @Summary 增加menu和角色关联
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleUpdate          true "请求参数"
// @Success 200  {object} response.Response{data=any} "更新成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "更新失败!"
// @Router  /role/update [put]
func (a *role) Update(c *gin.Context) response.Response {
	var info request.RoleUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Role.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorUpdated}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// AddMenus
// @Tags    SystemRole
// @Summary 增加menu和角色关联
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleAddMenus        true "请求参数"
// @Success 200  {object} response.Response{data=any} "角色添加菜单成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "角色添加菜单失败!"
// @Router  /role/addMenus [post]
func (a *role) AddMenus(c *gin.Context) response.Response {
	var info request.RoleAddMenus
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Role.AddMenus(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "角色添加菜单失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "角色添加菜单成功!"}
}

// Copy
// @Tags    SystemRole
// @Summary 复制角色
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleAddMenus        true "请求参数"
// @Success 200  {object} response.Response{data=any} "角色复制成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "角色复制失败!"
// @Router  /role/copy [post]
func (a *role) Copy(c *gin.Context) response.Response {
	var info request.RoleCopy
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Role.Copy(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "角色复制失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "角色复制成功!"}
}

// Delete
// @Tags    SystemRole
// @Summary 增加menu和角色关联关
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleAddMenus        true "请求参数"
// @Success 200  {object} response.Response{data=any} "删除成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "删除失败!"
// @Router  /role/delete [delete]
func (a *role) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.Role.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorDeleted}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessDeleted}
}

// List
// @Tags    SystemRole
// @Summary 分页获取角色列表数据
// @accept  application/json
// @Produce application/json
// @Param   data body     request.RoleSearch                   true "请求参数"
// @Success 200  {object} response.Response{data=[]model.Role} "获取分页数据成功!"
// @Failure 400  {object} response.Response{data=any}          "Bad Request"
// @Failure 500  {object} response.Response{data=any}          "获取分页数据失败!"
// @Router  /role/list [post]
func (a *role) List(c *gin.Context) response.Response {
	var info request.RoleSearch
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, count, err := service.Role.List(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: common.ErrorList}
	}
	return response.Response{Code: http.StatusOK, Data: common.NewPageResult(list, count, info.PageInfo), Message: common.SuccessList}
}

// MenuTree
// @Tags     SystemRole
// @Summary  获取角色动态菜单树
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.Role                                     true "请求参数"
// @Success  200  {object} response.Response{data=[]response.RoleMenuTree} "获取列表数据成功!"
// @Failure  400  {object} response.Response{data=any}                     "Bad Request"
// @Failure  500  {object} response.Response{data=any}                     "获取角色动态菜单树失败!"
// @Router   /role/menuTree [post]
func (a *role) MenuTree(c *gin.Context) response.Response {
	var info common.Role
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, err := service.Role.MenuTree(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "获取角色动态菜单树失败!"}
	}
	return response.Response{Code: http.StatusOK, Data: list, Message: "获取角色动态菜单树成功!"}
}
