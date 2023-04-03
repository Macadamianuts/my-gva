package model

// UsersRoles User 与 Role 的多对多关联表
type UsersRoles struct {
	UserId uint `json:"userId" gorm:"column:user_id;comment:用户Id"`
	RoleId uint `json:"roleId" gorm:"column:role_id;comment:角色Id"`
}

func (u *UsersRoles) TableName() string {
	return "system_users_roles"
}
