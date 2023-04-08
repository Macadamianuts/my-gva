package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/common"
	"gva-lbx/component/jwt"
	"gva-lbx/response"
	"net/http"
)

var User = new(user)

type user struct{}

// Create
// @Tags    SystemUser
// @Summary 创建用户
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserCreate          true "请求参数"
// @Success 200  {object} response.Response{data=any} "创建成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} ""创建失败!"
// @Router  /user/create [post]
func (u *user) Create(c *gin.Context) response.Response {
	var info request.UserCreate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.Create(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK}
}

// First
// @Tags    SystemUser
// @Summary 获取用户数据
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserFirst                  true "请求参数"
// @Success 200  {object} response.Response{data=model.User} "获取单条数据成功!"
// @Failure 400  {object} response.Response{data=any}        "Bad Request"
// @Failure 500  {object} response.Response{data=any}        "获取单条数据失败!"
// @Router  /user/first [post]
func (u *user) First(c *gin.Context) response.Response {
	var info request.UserFirst
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	data, err := service.User.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: data}
}

// FirstSelf
// @Tags    SystemUser
// @Summary 获取用户自身数据
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserFirst                  true "请求参数"
// @Success 200  {object} response.Response{data=model.User} "获取单条数据成功!"
// @Failure 400  {object} response.Response{data=any}        "Bad Request"
// @Failure 500  {object} response.Response{data=any}        "获取单条数据失败!"
// @Router  /user/self/first [post]
func (u *user) FirstSelf(c *gin.Context) response.Response {
	var info request.UserFirst
	claims, err := jwt.NewClaimsByGinContext(c)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	info.Id = claims.UserId
	data, err := service.User.First(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: data}
}

// Update
// @Tags    SystemUser
// @Summary 更新用户
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserUpdate          true "请求参数"
// @Success 200  {object} response.Response{data=any} "更新成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "更新失败!"
// @Router  /user/update [put]
func (u *user) Update(c *gin.Context) response.Response {
	var info request.UserUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// UpdateSelf
// @Tags    SystemUser
// @Summary 更新用户自身
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserUpdate          true "请求参数"
// @Success 200  {object} response.Response{data=any} "更新成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "更新失败!"
// @Router  /user/self/update [put]
func (u *user) UpdateSelf(c *gin.Context) response.Response {
	var info request.UserUpdate
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	claims, err := jwt.NewClaimsByGinContext(c)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	info.GormId.Id = claims.UserId
	err = service.User.Update(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessUpdated}
}

// SetRole
// @Tags     SystemUser
// @Summary  设置用户活跃角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.UserSetRole         true "请求参数"
// @Success  200  {object} response.Response{data=any} "设置角色成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "设置角色失败!"
// @Router   /user/setRole [patch]
func (u *user) SetRole(c *gin.Context) response.Response {
	var info request.UserSetRole
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.SetRole(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "设置角色失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "设置角色成功!"}
}

// SetRoles
// @Tags     SystemUser
// @Summary  设置用户多角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.UserSetRoles        true "请求参数"
// @Success  200  {object} response.Response{data=any} "设置用户多角色成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "设置用户多角色失败!"
// @Router   /user/setRoles [patch]
func (u *user) SetRoles(c *gin.Context) response.Response {
	var info request.UserSetRoles
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.SetRoles(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "设置用户多角色失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "设置用户多角色成功!"}
}

// ChangePassword
// @Tags     SystemUser
// @Summary  修改密码
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.ApiUpdate           true "请求参数"
// @Success  200  {object} response.Response{data=any} "修改密码成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "修改密码失败!"
// @Router   /user/changePassword [patch]
func (u *user) ChangePassword(c *gin.Context) response.Response {
	var info request.UserChangePassword
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.ChangePassword(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusOK, Message: "修改密码失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "修改密码成功!"}
}

// ResetPassword
// @Tags     SystemUser
// @Summary  重置密码
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     common.GormId               true "请求参数"
// @Success  200  {object} response.Response{data=any} "重置密码成功!"
// @Failure  400  {object} response.Response{data=any} "Bad Request"
// @Failure  500  {object} response.Response{data=any} "重置密码失败!"
// @Router   /user/resetPassword [patch]
func (u *user) ResetPassword(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.ResetPassword(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusOK, Message: "重置密码失败!"}
	}
	return response.Response{Code: http.StatusOK, Message: "重置密码成功!"}
}

// Delete
// @Tags    SystemUser
// @Summary 删除用户
// @accept  application/json
// @Produce application/json
// @Param   data body     common.GormId               true "请求参数"
// @Success 200  {object} response.Response{data=any} "删除成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "删除失败!"
// @Router  /user/delete [delete]
func (u *user) Delete(c *gin.Context) response.Response {
	var info common.GormId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.Delete(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: common.SuccessDeleted}
}

// Deletes
// @Tags    SystemUser
// @Summary 删除用户
// @accept  application/json
// @Produce application/json
// @Param   data body     common.GormIds              true "请求参数"
// @Success 200  {object} response.Response{data=any} "批量删除成功!"
// @Failure 400  {object} response.Response{data=any} "Bad Request"
// @Failure 500  {object} response.Response{data=any} "批量删除失败!"
// @Router  /user/deletes [delete]
func (u *user) Deletes(c *gin.Context) response.Response {
	var info common.GormIds
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	err = service.User.Deletes(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Message: "批量删除成功!"}
}

// List
// @Tags    SystemUser
// @Summary 更新用户
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserUpdate                   true "请求参数"
// @Success 200  {object} response.Response{data=[]model.User} "获取分页数据成功!"
// @Failure 400  {object} response.Response{data=any}          "Bad Request"
// @Failure 500  {object} response.Response{data=any}          "获取分页数据失败!"
// @Router  /user/list [post]
func (u *user) List(c *gin.Context) response.Response {
	var info request.UserSearch
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	list, count, err := service.User.List(c.Request.Context(), info)
	if err != nil {
		return response.Response{Code: http.StatusInternalServerError, Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: common.NewPageResult(list, count, info.PageInfo)}
}
