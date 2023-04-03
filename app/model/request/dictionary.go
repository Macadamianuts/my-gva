package request

import (
	"gorm.io/gen"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/common"
)

type DictionaryCreate struct {
	Name        string `json:"name" example:"字典名(中)"`
	Type        string `json:"type" example:"字典名(英)"`
	Status      bool   `json:"status" swaggertype:"string" example:"状态"`
	Description string `json:"description" example:"描述"`
}

func (r *DictionaryCreate) Create() model.Dictionary {
	return model.Dictionary{
		Name:        r.Name,
		Type:        r.Type,
		Status:      r.Status,
		Description: r.Description,
	}
}

type DictionaryFirst struct {
	common.GormId
	Type   string `json:"type" example:"字典名(英)"`
	Status *bool  `json:"status" swaggertype:"string" example:"bool 状态"`
}

func (r *DictionaryFirst) First() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Type != "" {
			tx = tx.Where(dao.Dictionary.Type.Eq(r.Type))
		}
		if r.Status != nil {
			tx = tx.Where(dao.Dictionary.ID.Eq(r.GormId.Id)).Or(dao.Dictionary.Status.Is(*r.Status))
		} else {
			tx = tx.Where(dao.Dictionary.ID.Eq(r.GormId.Id)).Or(dao.Dictionary.Status.Is(*r.Status))
		}
		return tx
	}
}

type DictionaryUpdate struct {
	common.GormId
	Name        string `json:"name" example:"字典名(中)"`
	Type        string `json:"type" example:"字典名(英)"`
	Status      bool   `json:"status" swaggertype:"string" example:"bool 状态"`
	Description string `json:"description" example:"描述"`
}

func (r *DictionaryUpdate) Update() model.Dictionary {
	return model.Dictionary{
		Name:        r.Name,
		Type:        r.Type,
		Status:      r.Status,
		Description: r.Description,
	}
}

type DictionarySearch struct {
	common.PageInfo
	Name        string `json:"name" example:"字典名(中)"`
	Type        string `json:"type" example:"字典名(英)"`
	Status      *bool  `json:"status" swaggertype:"string" example:"bool 状态"`
	Description string `json:"description" example:"描述"`
}

func (r *DictionarySearch) Search() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Name != "" {
			tx = tx.Where(dao.Dictionary.Name.Like("%" + r.Name + "%"))
		}
		if r.Type != "" {
			tx = tx.Where(dao.Dictionary.Type.Like("%" + r.Type + "%"))
		}
		if r.Status != nil {
			tx = tx.Where(dao.Dictionary.Status.Is(*r.Status))
		}
		if r.Description != "" {
			tx = tx.Where(dao.Dictionary.Description.Like("%" + r.Description + "%"))
		}
		return tx
	}
}
