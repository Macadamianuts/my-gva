package model

import "gva-lbx/global"

// AutoCodeHistoryField 代码生成器历史字段
type AutoCodeHistoryField struct {
	global.Model
	global.Operator
	AutoCodeHistoryID uint `json:"autoCodeHistoryID" gorm:"column:auto_code_history_id;comment:历史记录ID"`
	// 后端相关
	Name        string `json:"name" gorm:"column:name;comment:字段名"`
	Type        string `json:"type" gorm:"column:type;comment:字段数据类型"`
	Json        string `json:"json" gorm:"column:json;comment:字段 json tag"`
	Description string `json:"description" gorm:"column:description;comment:字段中文名"`
	// 数据库相关
	Size    string `json:"size" gorm:"column:size;comment:数据库字段长度"`
	Where   string `json:"where" gorm:"column:where;comment:数据库字段搜索条件"`
	Column  string `json:"column" gorm:"column:column;comment:数据库字段列名"`
	Comment string `json:"comment" gorm:"column:comment;comment:数据库字段描述"`
	Sort    bool   `json:"sort" gorm:"column:sort;comment:是否增加排序"`
	// 前端相关
	ErrorText  string `json:"errorText" gorm:"column:error_text;comment:校验失败文字"`
	Dictionary string `json:"dictionary" gorm:"column:dictionary;comment:关联字典"`
	Require    bool   `json:"require" gorm:"column:require;comment:是否必填"`
	Clearable  bool   `json:"clearable" gorm:"column:clearable;comment:是否可清空"`
}

func (a *AutoCodeHistoryField) TableName() string {
	return "auto_code_history_fields"
}
