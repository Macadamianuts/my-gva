package request

import (
	"gorm.io/gen"
	"gva-lbx/common"
	"gva-lbx/plugin/coder/model/dao"
)

type AutoCodeHistoryRepeat struct {
	Struct string `json:"struct" example:"结构体名称"`
	Plugin string `json:"plugin" example:"插件名"`
}

type AutoCodeHistoryRollback struct {
	common.GormId
	DeleteTable bool `json:"deleteTable" swaggertype:"string" example:"bool 是否删除表"`
}

type AutoCodeHistoryList struct {
	common.PageInfo
	Plugin string `json:"plugin" example:"插件名"`
}

func (r *AutoCodeHistoryList) Search() common.GenScopes {
	return func(tx gen.Dao) gen.Dao {
		if r.Plugin != "" {
			tx = tx.Where(dao.AutoCodeHistory.Plugin.Eq(r.Plugin))
		}
		return tx
	}
}
