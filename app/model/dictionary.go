package model

import "gva-lbx/global"

// Dictionary 字典
type Dictionary struct {
	global.Model
	global.Operator
	Name        string              `json:"name" gorm:"column:name;comment:字典名(中)"`
	Type        string              `json:"type" gorm:"column:type;comment:字典名(英)"`
	Status      bool                `json:"status" gorm:"column:status;comment:状态"`
	Description string              `json:"description" gorm:"column:description;comment:描述"`
	Details     []*DictionaryDetail `json:"details" gorm:"foreignKey:DictionaryId;references:ID"`
}

func (d *Dictionary) TableName() string {
	return "system_dictionaries"
}
