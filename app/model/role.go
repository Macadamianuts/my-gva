package model

import "gva-lbx/global"

// Role 角色
type Role struct {
	global.Model
	Name          string `json:"name" gorm:"column:name;comment:角色名"`
	ParentId      uint   `json:"parentId" gorm:"column:parent_id;comment:父角色Id"`
	DefaultRouter string `json:"defaultRouter" gorm:"default:dashboard;column:default_router;comment:默认菜单"`
	// 关联
	Users    []*User `json:"users" gorm:"many2many:system_users_roles;foreignKey:ID;joinForeignKey:RoleId;references:ID;JoinReferences:UserId"`
	Menus    []*Menu `json:"menus" gorm:"many2many:system_roles_menus;foreignKey:ID;joinForeignKey:RoleId;references:ID;JoinReferences:MenuId"`
	Children []*Role `json:"children" gorm:"-"`
}

// TableName 自定义表名
func (r *Role) TableName() string {
	return "system_roles"
}
