package model

// RolesMenuButtons Role 与 MenuButton 得多对多关联表
type RolesMenuButtons struct {
	RoleId   uint `json:"roleId" gorm:"column:role_id;comment:角色Id"`
	MenuId   uint `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	ButtonId uint `json:"buttonId" gorm:"column:button_id;comment:菜单按钮Id"`
}

func (r *RolesMenuButtons) TableName() string {
	return "system_roles_menu_buttons"
}
