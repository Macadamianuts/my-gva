package internal

import (
	"context"
	"github.com/pkg/errors"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
)

var Menu = new(menu)

type menu struct{}

// GetTreeMap 获取菜单树
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (m *menu) GetTreeMap(ctx context.Context) (map[uint][]*model.Menu, error) {
	query := dao.Q.WithContext(ctx).Menu
	entities, err := query.Order(dao.Menu.Sort).Preload(dao.Menu.Parameters).Find()
	if err != nil {
		return nil, errors.Wrap(err, "获取列表数据失败!")
	}
	length := len(entities)
	tree := make(map[uint][]*model.Menu, length)
	for i := 0; i < length; i++ {
		value, ok := tree[entities[i].ParentId]
		if !ok {
			value = make([]*model.Menu, 0, 5)
		}
		value = append(value, entities[i])
		tree[entities[i].ParentId] = value
	}
	return tree, nil
}

// GetChildren 递归获取子菜单
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (m *menu) GetChildren(menu *model.Menu, tree map[uint][]*model.Menu) {
	menu.Children = tree[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		m.GetChildren(menu.Children[i], tree)
	}
}
