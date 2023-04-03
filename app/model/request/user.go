package request

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gen"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/common"
)

type UserCreate struct {
	Phone       string `json:"phone" example:"用户手机号"`
	Email       string `json:"email" example:"用户邮箱"`
	Avatar      string `json:"avatar" example:"用户头像"`
	Username    string `json:"username" example:"用户登录名"`
	Password    string `json:"password" example:"用户密码"`
	Nickname    string `json:"nickname" example:"用户昵称"`
	SideMode    string `json:"sideMode" example:"用户侧边主题"`
	BaseColor   string `json:"baseColor" example:"基础颜色"`
	ActiveColor string `json:"activeColor" example:"活跃颜色"`
	RoleId      uint   `json:"RoleId" swaggertype:"string" example:"uint 角色Id"`
}

type UserLogin struct {
	Captcha   string `json:"captcha" example:"验证码"`
	Username  string `json:"username" example:"用户登录名"`
	Password  string `json:"password" example:"用户密码"`
	CaptchaId string `json:"captchaId" example:"验证码ID"`
}

func (r *UserCreate) Create() model.User {
	return model.User{
		Uuid:        uuid.NewV4().String(),
		Phone:       r.Phone,
		Email:       r.Email,
		Avatar:      r.Avatar,
		Username:    r.Username,
		Password:    r.Password,
		Nickname:    r.Nickname,
		SideMode:    r.SideMode,
		BaseColor:   r.BaseColor,
		ActiveColor: r.ActiveColor,
		Enable:      false,
		RoleId:      r.RoleId,
	}
}

type UserFirst struct {
	Id   uint   `json:"id" swaggertype:"string" example:"uint 主键id"`
	Uuid string `json:"uuid" example:"uuid"`
}

func (r *UserFirst) First() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Id != 0 {
			tx = tx.Where(dao.User.ID.Eq(r.Id))
		}
		if r.Uuid != "" {
			tx = tx.Where(dao.User.Uuid.Eq(r.Uuid))
		}
		return tx
	}
}

type UserUpdate struct {
	common.GormId
	Phone       string `json:"phone" example:"用户手机号"`
	Email       string `json:"email" example:"用户邮箱"`
	Avatar      string `json:"avatar" example:"用户头像"`
	Username    string `json:"username" example:"用户登录名"`
	Nickname    string `json:"nickname" example:"用户昵称"`
	SideMode    string `json:"sideMode" example:"用户侧边主题"`
	BaseColor   string `json:"baseColor" example:"基础颜色"`
	ActiveColor string `json:"activeColor" example:"活跃颜色"`
	Enable      bool   `json:"enable" swaggertype:"string" example:"bool 用户是否被冻结(false:正常,true:冻结)"`
}

func (r *UserUpdate) Update() model.User {
	return model.User{
		Phone:       r.Phone,
		Email:       r.Email,
		Avatar:      r.Avatar,
		Username:    r.Username,
		Nickname:    r.Nickname,
		SideMode:    r.SideMode,
		BaseColor:   r.BaseColor,
		ActiveColor: r.ActiveColor,
		Enable:      r.Enable,
	}
}

type UserChangePassword struct {
	UserId      uint   `json:"-"`
	Password    string `json:"password" example:"旧密码"`
	NewPassword string `json:"newPassword" example:"新密码"`
}

type UserSetRole struct {
	UserId uint `json:"userId" swaggertype:"string" example:"uint 用户Id"`
	RoleId uint `json:"RoleId" swaggertype:"string" example:"角色Id"`
}

type UserSetRoles struct {
	UserId  uint   `json:"userId" swaggertype:"string" example:"uint 用户Id"`
	RoleIds []uint `json:"RoleIds" swaggertype:"string" example:"角色Id"`
}

type UserSearch struct {
	common.PageInfo
	Username string `json:"username" example:"用户名"`
}

func (r *UserSearch) Search() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Username != "" {
			tx = tx.Where(dao.User.Username.Eq(r.Username))
		}
		return tx
	}
}
