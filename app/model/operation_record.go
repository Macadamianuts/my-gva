package model

import (
	"gva-lbx/global"
	"time"
)

// OperationRecord 操作记录
type OperationRecord struct {
	global.Model
	Ip           string        `json:"ip" gorm:"column:ip;comment:请求ip"`
	Path         string        `json:"path" gorm:"column:path;comment:请求路径"`
	Agent        string        `json:"agent" gorm:"column:agent;comment:代理"`
	Method       string        `json:"method" gorm:"column:method;comment:请求方法"`
	Request      string        `json:"request" gorm:"type:text;column:request;comment:请求内容"`
	Response     string        `json:"response" gorm:"type:text;column:response;comment:响应内容"`
	ErrorMessage string        `json:"error_message" gorm:"column:error_message;comment:错误信息"`
	Status       int           `json:"status" gorm:"column:status;comment:请求状态"`
	Latency      time.Duration `json:"latency" gorm:"column:latency;comment:延迟"`
	// 关联
	UserId int  `json:"user_id" gorm:"column:user_id;comment:用户id"`
	User   User `json:"user" gorm:"foreignKey:ID;references:UserId"`
}

func (o *OperationRecord) TableName() string {
	return "system_operation_records"
}
