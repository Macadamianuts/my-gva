package request

import (
	"gva-lbx/app/model"
	"gva-lbx/common"
)

type MenuCreateMeta struct {
	Icon        string `json:"icon" example:"菜单图标"`
	Title       string `json:"title" example:"菜单名"`
	CloseTab    bool   `json:"closeTab" swaggertype:"string" example:"bool 自动关闭tab"`
	KeepAlive   bool   `json:"keepAlive" swaggertype:"string" example:"bool 是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" swaggertype:"string" example:"bool 是否是基础路由"`
}

type MenuCreateParameter struct {
	Key   string `json:"key" example:"地址栏携带参数的key"`
	Type  string `json:"type" example:"地址栏携带参数为params还是query"`
	Value string `json:"value" example:"地址栏携带参数的值"`
}

type MenuCreate struct {
	Path       string                `json:"path" example:"路由路径"`
	Name       string                `json:"name" example:"路由名字"`
	Sort       int                   `json:"sort" swaggertype:"string" example:"int 排序标记"`
	Hidden     bool                  `json:"hidden" swaggertype:"string" example:"bool 是否在列表隐藏"`
	ParentId   uint                  `json:"parentId" swaggertype:"string" example:"uint 父菜单Id"`
	Component  string                `json:"component" example:"对应前端文件路径"`
	Parameters []MenuCreateParameter `json:"parameters"`
	MenuCreateMeta
}

func (r *MenuCreate) Create() model.Menu {
	length := len(r.Parameters)
	parameters := make([]*model.MenuParameter, 0, length)
	for i := 0; i < length; i++ {
		parameters = append(parameters, &model.MenuParameter{
			Key:   r.Parameters[i].Key,
			Type:  r.Parameters[i].Type,
			Value: r.Parameters[i].Value,
		})
	}
	return model.Menu{
		Path:      r.Path,
		Name:      r.Name,
		Sort:      r.Sort,
		Hidden:    r.Hidden,
		ParentId:  r.ParentId,
		Component: r.Component,
		Meta: model.Meta{
			Icon:        r.MenuCreateMeta.Icon,
			Title:       r.MenuCreateMeta.Title,
			CloseTab:    r.MenuCreateMeta.CloseTab,
			KeepAlive:   r.MenuCreateMeta.KeepAlive,
			DefaultMenu: r.MenuCreateMeta.DefaultMenu,
		},
		Parameters: parameters,
	}
}

type MenuUpdateMeta struct {
	Icon        string             `json:"icon" example:"菜单图标"`
	Title       string             `json:"title" example:"菜单名"`
	CloseTab    bool               `json:"closeTab" swaggertype:"string" example:"bool 自动关闭tab"`
	KeepAlive   bool               `json:"keepAlive" swaggertype:"string" example:"bool 是否缓存"`
	DefaultMenu bool               `json:"defaultMenu" swaggertype:"string" example:"bool 是否是基础路由"`
	Buttons     []MenuUpdateButton `json:"buttons"`
}

type MenuUpdateButton struct {
	Title       string `json:"title" example:"标题"`
	ApiId       uint   `json:"apiId" swaggertype:"string" example:"uint api关联Id"`
	MenuId      uint   `json:"menuId" swaggertype:"string" example:"uint 菜单id"`
	Description string `json:"description" example:"描述"`
}

type MenuUpdateParameter struct {
	Key    string `json:"key" example:"地址栏携带参数的key"`
	Type   string `json:"type" example:"地址栏携带参数为params还是query"`
	Value  string `json:"value" example:"地址栏携带参数的值"`
	MenuId uint   `json:"menuId" swaggertype:"string" example:"uint 菜单id"`
}
type MenuUpdate struct {
	common.GormId
	Path       string                `json:"path" example:"路由路径"`
	Name       string                `json:"name" example:"路由名字"`
	Sort       int                   `json:"sort" swaggertype:"string" example:"int 排序标记"`
	Hidden     bool                  `json:"hidden" swaggertype:"string" example:"bool 是否在列表隐藏"`
	ParentId   uint                  `json:"parentId" swaggertype:"string" example:"uint 父菜单Id"`
	Component  string                `json:"component" example:"对应前端文件路径"`
	Parameters []MenuUpdateParameter `json:"parameters"`
	MenuUpdateMeta
}

func (r *MenuUpdate) Update() model.Menu {
	length := len(r.Parameters)
	parameters := make([]*model.MenuParameter, 0, length)
	for i := 0; i < length; i++ {
		parameters = append(parameters, &model.MenuParameter{
			Key:    r.Parameters[i].Key,
			Type:   r.Parameters[i].Type,
			Value:  r.Parameters[i].Value,
			MenuId: r.Parameters[i].MenuId,
		})
	}
	return model.Menu{
		Path:      r.Path,
		Name:      r.Name,
		Sort:      r.Sort,
		Hidden:    r.Hidden,
		ParentId:  r.ParentId,
		Component: r.Component,
		Meta: model.Meta{
			Icon:        r.MenuUpdateMeta.Icon,
			Title:       r.MenuUpdateMeta.Title,
			CloseTab:    r.MenuUpdateMeta.CloseTab,
			KeepAlive:   r.MenuUpdateMeta.KeepAlive,
			DefaultMenu: r.MenuUpdateMeta.DefaultMenu,
		},
		Parameters: parameters,
	}
}
