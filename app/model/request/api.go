package request

import (
	"go.uber.org/zap"
	"gorm.io/gen"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/common"
)

type ApiCreate struct {
	Path        string `json:"path" example:"路由路径"`
	Method      string `json:"method" example:"方法"`
	ApiGroup    string `json:"apiGroup" example:"路由组"`
	Description string `json:"description" example:"路由功能描述"`
}

func (r *ApiCreate) Create() model.Api {
	return model.Api{
		Path:        r.Path,
		Method:      r.Method,
		ApiGroup:    r.ApiGroup,
		Description: r.Description,
	}
}

type ApiUpdate struct {
	common.GormId
	Path        string `json:"path" example:"路由路径"`
	Method      string `json:"method" example:"方法"`
	ApiGroup    string `json:"apiGroup" example:"路由组"`
	Description string `json:"description" example:"路由功能描述"`
}

func (r *ApiUpdate) Update() model.Api {
	return model.Api{
		Path:        r.Path,
		Method:      r.Method,
		ApiGroup:    r.ApiGroup,
		Description: r.Description,
	}
}

func (r *ApiUpdate) Casbin() model.Casbin {
	return model.Casbin{
		Path:   r.Path,
		Method: r.Method,
	}
}

type ApiSearch struct {
	common.PageInfo
	common.OrderInfo
	Path        string `json:"path" example:"路由路径"`
	Method      string `json:"method" example:"方法"`
	ApiGroup    string `json:"apiGroup" example:"路由组"`
	Description string `json:"description" example:"路由功能描述"`
}

func (r *ApiSearch) Search() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Path != "" {
			tx = tx.Where(dao.Api.Path.Like("%" + r.Path + "%"))
		}
		if r.Description != "" {
			tx = tx.Where(dao.Api.Description.Like("%" + r.Description + "%"))
		}
		if r.Method != "" {
			tx = tx.Where(dao.Api.Method.Eq(r.Method))
		}
		if r.ApiGroup != "" {
			tx = tx.Where(dao.Api.ApiGroup.Eq(r.ApiGroup))
		}
		return tx
	}
}

func (r *ApiSearch) Order() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.OrderInfo.Order != "" {
			order, ok := dao.Api.GetFieldByName(r.OrderInfo.Order)
			if ok {
				if r.OrderInfo.Desc {
					tx = tx.Order(order.Desc())
				} else {
					tx.Order(order)
				}
			} else {
				zap.L().Error("非法的排序字段", zap.String("order", r.OrderInfo.Order))
			}
		}
		return tx
	}
}
