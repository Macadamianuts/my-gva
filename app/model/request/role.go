package request

import (
	"context"
	"gorm.io/gen"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/common"
	"gva-lbx/global"
	"strconv"
)

type RoleCreate struct {
	Name          string `json:"authorityName" example:"角色名字"`
	RoleId        uint   `json:"authorityId" swaggertype:"string" example:"uint 角色id"`
	ParentId      uint   `json:"parentId" swaggertype:"string" example:"uint 父角色id"`
	DefaultRouter string `json:"defaultRouter" example:"默认菜单"`
}

func (r *RoleCreate) Create() model.Role {
	return model.Role{
		Model:         global.Model{ID: r.RoleId},
		Name:          r.Name,
		Menus:         r.Menus(),
		ParentId:      r.ParentId,
		DefaultRouter: r.DefaultRouter,
	}
}

func (r *RoleCreate) Menus() []*model.Menu {
	entities, _ := dao.Q.WithContext(context.Background()).Menu.Where(dao.Menu.Name.Eq("dashboard")).Find()
	return entities
}

func (r *RoleCreate) Casbin() []*model.Casbin {
	roleId := strconv.Itoa(int(r.RoleId))
	return []*model.Casbin{
		{Path: "/role/menu", Method: "POST", RoleId: roleId},
		{Path: "/jwt/jsonInBlacklist", Method: "POST", RoleId: roleId},
		{Path: "/user/login", Method: "POST", RoleId: roleId},
		{Path: "/user/admin_register", Method: "POST", RoleId: roleId},
		{Path: "/user/changePassword", Method: "POST", RoleId: roleId},
		{Path: "/user/setUserAuthority", Method: "POST", RoleId: roleId},
		{Path: "/user/update", Method: "PUT", RoleId: roleId},
		{Path: "/user/first", Method: "GET", RoleId: roleId},
	}
}

type RoleAddMenus struct {
	RoleId uint   `json:"roleId" swaggertype:"string" example:"uint 角色id"`
	MenuId []uint `json:"menuId" swaggertype:"string" example:"[]uint 菜单id"`
}

type RoleCopy struct {
	OldRoleId uint       `json:"oldRoleId" swaggertype:"string" example:"uint 旧角色ID"`
	Role      model.Role `json:"role"`
}

// ToCommonRole RoleCopy 转换 common.Role
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *RoleCopy) ToCommonRole() common.Role {
	return common.Role{RoleId: r.OldRoleId}
}

type RoleUpdate struct {
	Name          string `json:"name" example:"角色名字"`
	RoleId        uint   `json:"authorityId" swaggertype:"string" example:"uint 角色id"`
	ParentId      uint   `json:"parentId" swaggertype:"string" example:"uint 父角色id"`
	DefaultRouter string `json:"defaultRouter" example:"默认菜单"`
}

func (r *RoleUpdate) Update() model.Role {
	return model.Role{
		Model:         global.Model{ID: r.RoleId},
		ParentId:      r.ParentId,
		Name:          r.Name,
		DefaultRouter: r.DefaultRouter,
	}
}

type RoleSetData struct {
	RoleId uint         `json:"authorityId" swaggertype:"string" example:"uint 角色id"`
	Data   []RoleCreate `json:"data"`
}

func (r *RoleSetData) RolesData() []*model.Role {
	length := len(r.Data)
	entities := make([]*model.Role, 0, length)
	for i := 0; i < length; i++ {
		entity := r.Data[i].Create()
		entities = append(entities, &entity)
	}
	return entities
}

type RoleSetMenus struct {
	RoleId uint         `json:"authorityId" swaggertype:"string" example:"uint 角色id"`
	Menus  []MenuCreate `json:"menus"`
}

func (r *RoleSetMenus) RolesMenus() []*model.Menu {
	length := len(r.Menus)
	entities := make([]*model.Menu, 0, length)
	for i := 0; i < length; i++ {
		entity := r.Menus[i].Create()
		entities = append(entities, &entity)
	}
	return entities
}

type RoleSearch struct {
	common.PageInfo
	Name string `json:"name" example:"角色名字"`
}

func (r *RoleSearch) Search() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Name != "" {
			tx = tx.Where(dao.Role.Name.Like("%" + r.Name + "%"))
		}
		tx = tx.Where(dao.Role.ParentId.Eq(0))
		return tx
	}
}
