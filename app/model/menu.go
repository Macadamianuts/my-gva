package model

import "gva-lbx/global"

// Menu 菜单
type Menu struct {
	global.Model
	Path      string `json:"path" gorm:"column:path;comment:路由路径"`
	Name      string `json:"name" gorm:"column:name;comment:路由名字"`
	Sort      int    `json:"sort" gorm:"column:sort;comment:排序标记"`
	Hidden    bool   `json:"hidden" gorm:"column:hidden;comment:是否在列表隐藏"`
	ParentId  uint   `json:"parentId" gorm:"column:parent_id;comment:父菜单Id"`
	Component string `json:"component" gorm:"column:component;comment:对应前端文件路径"`
	Meta      `json:"meta" gorm:"embedded;comment:附加属性"`

	Roles      []*Role          `json:"authorities" gorm:"many2many:system_Role_menus;foreignKey:ID;joinForeignKey:MenuId;references:ID;JoinReferences:RoleId"`
	Children   []*Menu          `json:"children" gorm:"-"`
	Parameters []*MenuParameter `json:"parameters" gorm:"foreignKey:MenuId;references:ID"`
}

// Meta 菜单附加属性
type Meta struct {
	Icon        string `json:"icon" gorm:"column:icon;comment:菜单图标"`
	Title       string `json:"title" gorm:"column:title;comment:菜单名"`
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;comment:自动关闭tab"`
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;comment:是否是基础路由"`
}

// TableName 自定义表名
func (m *Menu) TableName() string {
	return "system_menus"
}
