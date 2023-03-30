package response

import "gva-lbx/app/model"

// RoleMenuTree 角色菜单响应体
type RoleMenuTree struct {
	*model.Menu
	MenuId   uint            `json:"menuId" swaggertype:"string" example:"uint 菜单id"`
	RoleId   uint            `json:"roleId" swaggertype:"string" example:"uint 角色id"`
	Children []*RoleMenuTree `json:"children"` // 子菜单
}
