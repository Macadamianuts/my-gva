package request

import (
	"gorm.io/gen"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/common"
)

type DictionaryDetailCreate struct {
	Label        string `json:"label" example:"展示值"`
	Sort         int    `json:"sort" swaggertype:"string" example:"int 排序标记"`
	Value        int    `json:"value" swaggertype:"string" example:"int 字典值"`
	Status       bool   `json:"status" swaggertype:"string" example:"bool 启用状态"`
	DictionaryId uint   `json:"dictionaryId" swaggertype:"string" example:"int 字典Id"`
}

func (r *DictionaryDetailCreate) Create() model.DictionaryDetail {
	return model.DictionaryDetail{
		Label:        r.Label,
		Sort:         r.Sort,
		Value:        r.Value,
		Status:       r.Status,
		DictionaryId: r.DictionaryId,
	}
}

type DictionaryDetailUpdate struct {
	common.GormId
	Label        string `json:"label" example:"展示值"`
	Sort         int    `json:"sort" swaggertype:"string" example:"int 排序标记"`
	Value        int    `json:"value" swaggertype:"string" example:"int 字典值"`
	Status       bool   `json:"status" swaggertype:"string" example:"bool 启用状态"`
	DictionaryId uint   `json:"dictionaryId" swaggertype:"string" example:"int 字典Id"`
}

func (r *DictionaryDetailUpdate) Update() model.DictionaryDetail {
	return model.DictionaryDetail{
		Label:        r.Label,
		Sort:         r.Sort,
		Value:        r.Value,
		Status:       r.Status,
		DictionaryId: r.DictionaryId,
	}
}

type DictionaryDetailSearch struct {
	common.PageInfo
	Label        string `json:"label" example:"展示值"`
	Value        *int   `json:"value" swaggertype:"string" example:"int 字典值"`
	Status       *bool  `json:"status" swaggertype:"string" example:"bool 启用状态"`
	DictionaryId *uint  `json:"dictionaryId" swaggertype:"string" example:"uint 字典Id"`
}

func (r *DictionaryDetailSearch) Search() func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if r.Label != "" {
			tx = tx.Where(dao.DictionaryDetail.Label.Like("%" + r.Label + "%"))
		}
		if r.Value != nil {
			tx = tx.Where(dao.DictionaryDetail.Value.Eq(*r.Value))
		}
		if r.Status != nil {
			tx = tx.Where(dao.DictionaryDetail.Status.Is(*r.Status))
		}
		if r.DictionaryId != nil {
			tx = tx.Where(dao.DictionaryDetail.DictionaryId.Eq(*r.DictionaryId))
		}
		return tx
	}
}
