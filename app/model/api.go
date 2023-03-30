package model

import "gva-lbx/global"

type Api struct {
	global.Model
	global.Operator
	Path        string `json:"path" gorm:"column:path;comment:api路径"`
	Method      string `json:"method" gorm:"default:POST;column:method;comment:方法:POST(默认)|GET|PUT|DELETE"`
	ApiGroup    string `json:"apiGroup" gorm:"column:api_group;comment:api组"`
	Description string `json:"description" gorm:"column:description;comment:api中文描述"`
	IsEffective bool   `json:"isEffective" gorm:"column:is_effective;comment:是否有效"`
}

func (a *Api) TableName() string {
	return "system_apis"
}
