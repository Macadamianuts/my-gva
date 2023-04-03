package model

// RolesMenus Role 与 Menu 得多对多关联表
type RolesMenus struct {
	RoleId uint `json:"roleId" gorm:"column:role_id;comment:角色Id"`
	MenuId uint `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
}

func (r *RolesMenus) TableName() string {
	return "systems_roles_menus"
}
