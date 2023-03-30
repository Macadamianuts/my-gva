package model

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gva-lbx/global"
)

// User 用户
type User struct {
	global.Model
	Uuid        string `json:"uuid" gorm:"column:uuid;comment:用户UUID"`
	Phone       string `json:"phone" gorm:"column:phone;comment:用户手机号"`
	Email       string `json:"email" gorm:"column:email;comment:用户邮箱"`
	Avatar      string `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;column:avatar;comment:用户头像"`
	Username    string `json:"username" gorm:"column:username;comment:用户登录名"`
	Password    string `json:"-" gorm:"column:password;comment:用户登录密码"`
	Nickname    string `json:"nickname" gorm:"column:nickname;comment:用户昵称"`
	SideMode    string `json:"sideMode" gorm:"column:side_mode;comment:用户侧边主题"`
	BaseColor   string `json:"baseColor" gorm:"column:base_color;comment:基础颜色"`
	ActiveColor string `json:"activeColor" gorm:"column:active_color;comment:活跃颜色"`
	Enable      bool   `json:"enable" gorm:"default:1;column:enable;comment:用户是否被冻结"`
	// 关联
	RoleId uint   `json:"RoleId" gorm:"default:888;column:Role_id;comment:用户角色ID"`
	Role   Role   `json:"Role" gorm:"foreignKey:RoleId;references:ID"` // 用户角色
	Roles  []Role `json:"users" gorm:"many2many:system_users_authorities;foreignKey:ID;joinForeignKey:UserId;references:ID;JoinReferences:RoleId"`
}

// CompareHashAndPassword 密码检查 false 校验失败, true 校验成功
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) CompareHashAndPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.Wrap(err, "密码校验失败!")
	}
	return nil
}

// EncryptedPassword 加密密码
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) EncryptedPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "密码加密失败!")
	}
	u.Password = string(password)
	return nil
}

// TableName 自定义表名
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) TableName() string {
	return "system_users"
}
