package request

import "gva-lbx/app/model"

type Claims struct {
	UserId uint `json:"userId" example:"用户id"`
	RoleId uint `json:"roleId" example:"角色id"`
}

func NewClaims(user *model.User) Claims {
	return Claims{
		UserId: user.ID,
		RoleId: user.RoleId,
	}
}
