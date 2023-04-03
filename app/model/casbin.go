package model

type Casbin struct {
	Path   string `json:"path" gorm:"column:v2;comment:路由路径"`
	RoleId string `json:"roleId" gorm:"column:v1;comment:角色id"`
	Method string `json:"method" gorm:"column:v3;comment:请求方法"`
}

func (c *Casbin) TableName() string {
	return "casbin_rules"
}
