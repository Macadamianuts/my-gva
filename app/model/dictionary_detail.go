package model

import "gva-lbx/global"

// DictionaryDetail 字典详情
type DictionaryDetail struct {
	global.Model
	global.Operator
	Label        string `json:"label" gorm:"column:label;comment:展示值"`
	Sort         int    `json:"sort" gorm:"column:sort;comment:排序标记"`
	Value        int    `json:"value" gorm:"column:value;comment:字典值"`
	Status       bool   `json:"status" gorm:"column:status;comment:启用状态"`
	DictionaryId uint   `json:"dictionaryId" gorm:"column:dictionary_id;comment:字典Id"`
}

func (d *DictionaryDetail) TableName() string {
	return "system_dictionary_details"
}
