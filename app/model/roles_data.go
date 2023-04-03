package model

type RolesData struct {
	Data   uint `json:"data" gorm:"column:data;comment:数据Id"`
	RoleId uint `json:"roleId" gorm:"column:role_id;comment:角色Id"`
}

func (r *RolesData) TableName() string {
	return "system_roles_data"
}
