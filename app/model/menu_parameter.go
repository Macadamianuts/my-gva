package model

import "gva-lbx/global"

// MenuParameter 菜单参数
type MenuParameter struct {
	global.Model
	Key    string `json:"key" gorm:"column:key;comment:地址栏携带参数的key"`
	Type   string `json:"type" gorm:"column:type;comment:地址栏携带参数为params还是query"`
	Value  string `json:"value" gorm:"column:value;comment:地址栏携带参数的值"`
	MenuId uint   `json:"menuId" gorm:"column:menu_id;comment:菜单id"`
}

func (m *MenuParameter) TableName() string {
	return "system_menu_parameters"
}
