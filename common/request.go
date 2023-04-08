package common

import (
	"strconv"

	"gorm.io/gen"
)

type GormId struct {
	Id uint `json:"id" form:"id" swaggertype:"string" example:"uint 主键id"`
}

// ToRole GormId 转 Role
func (r *GormId) ToRole() Role {
	return Role{RoleId: r.Id}
}

type Role struct {
	RoleId uint `json:"roleId" form:"roleId" swaggertype:"string" example:"uint 角色id"`
}

func (r *Role) String() string {
	return strconv.Itoa(int(r.RoleId))
}

type GormIds struct {
	Ids []uint `json:"ids" form:"ids" swaggertype:"string" example:"[]uint 主键id"`
}

type PageInfo struct {
	Page     int  `json:"page" form:"page" swaggertype:"string" example:"int 页数"`
	PageSize int  `json:"pageSize" form:"pageSize" swaggertype:"string" example:"int 页大小"`
	Pass     bool `json:"-" form:"-"`
}

// Paginate 分页器共享函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *PageInfo) Paginate() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if p.Pass {
			return tx
		}
		if p.Page == 0 {
			p.Page = 1
		}
		switch {
		case p.PageSize > 100:
			p.PageSize = 100
		case p.PageSize <= 0:
			p.PageSize = 10
		}
		offset := (p.Page - 1) * p.PageSize
		return tx.Offset(offset).Limit(p.PageSize)
	}
}

// OrderInfo 用于查询语句排序
type OrderInfo struct {
	Desc  bool   `json:"desc" form:"desc"`
	Order string `json:"order" form:"order"`
}
