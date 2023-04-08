package model

import "gva-lbx/global"

// AutoCodeHistory 代码生成器历史记录
type AutoCodeHistory struct {
	global.Model
	global.Operator
	Type          string                 `json:"type" gorm:"column:type;comment:模版类型"`
	Plugin        string                 `json:"plugin" gorm:"column:plugin;comment:插件名"`
	Struct        string                 `json:"struct" gorm:"column:struct;comment:结构体名称"`
	Filename      string                 `json:"filename" gorm:"column:filename;comment:文件夹名称"`
	TableName     string                 `json:"tableName" gorm:"column:table_name;comment:表名"`
	Description   string                 `json:"description" gorm:"column:description;comment:结构体中文名称"`
	Abbreviation  string                 `json:"abbreviation" gorm:"column:abbreviation;comment:结构体简称"`
	UnderlineName string                 `json:"underlineName" gorm:"column:underline_name;comment:go文件名(下划线)"`
	AutoMoveFile  bool                   `json:"autoMoveFile" gorm:"column:auto_move_file;comment:是否自动移动文件"`
	AutoCreateApi bool                   `json:"autoCreateApi" gorm:"column:auto_create_api;comment:是否自动创建api"`
	Fields        []AutoCodeHistoryField `json:"fields" gorm:"foreignKey:AutoCodeHistoryID;references:ID"`
}
