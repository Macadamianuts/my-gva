package internal

import (
	"context"
	"github.com/pkg/errors"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/response"
	"gva-lbx/common"
)

var Role = new(role)

type role struct{}

// GetTreeMap 获取角色树
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *role) GetTreeMap(ctx context.Context) (map[uint][]*model.Role, error) {
	query := dao.Q.WithContext(ctx).Role
	entities, err := query.Where(dao.Role.ParentId.Neq(0)).Find()
	if err != nil {
		return nil, errors.Wrap(err, "获取子角色列表数据失败!")
	}
	length := len(entities)
	tree := make(map[uint][]*model.Role)
	for i := 0; i < length; i++ {
		value, ok := tree[entities[i].ParentId]
		if !ok {
			value = make([]*model.Role, 0, 5)
		}
		value = append(value, entities[i])
		tree[entities[i].ParentId] = value
	}
	return tree, nil
}

// GetChildren 获取子角色
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *role) GetChildren(entity *model.Role, tree map[uint][]*model.Role) {
	entity.Children = tree[entity.ID]
	for i := 0; i < len(entity.Children); i++ {
		r.GetChildren(entity.Children[i], tree)
	}
}

// GetMenuTreeMap 获取角色菜单树
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *role) GetMenuTreeMap(ctx context.Context, info common.Role) (map[uint][]*response.RoleMenuTree, error) {
	entities, err := dao.Q.WithContext(ctx).RolesMenus.Where(dao.RolesMenus.RoleId.Eq(info.RoleId)).Find()
	if err != nil {
		return nil, errors.Wrap(err, "角色菜单关联表数据查找失败!")
	}
	length := len(entities)
	ids := make([]uint, 0, length)
	for i := 0; i < length; i++ {
		ids = append(ids, entities[i].MenuId)
	}
	query := dao.Q.WithContext(ctx).Menu
	menus, err := query.Where(dao.Menu.ID.In(ids...)).Order(dao.Menu.Sort).Preload(dao.Menu.Parameters).Find()
	if err != nil {
		return nil, errors.Wrap(err, "获取角色菜单数据失败!")
	}
	length = len(menus)
	tree := make(map[uint][]*response.RoleMenuTree, length)
	for i := 0; i < length; i++ {
		value, ok := tree[menus[i].ParentId]
		if !ok {
			value = make([]*response.RoleMenuTree, 0, 7)
		}
		value = append(value, &response.RoleMenuTree{
			Menu:   menus[i],
			MenuId: menus[i].ID,
			RoleId: info.RoleId,
		})
		tree[menus[i].ParentId] = value
	}
	return tree, nil
}

// GetMenuChildren 获取角色子菜单
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *role) GetMenuChildren(entity *response.RoleMenuTree, tree map[uint][]*response.RoleMenuTree) {
	entity.Children = tree[entity.ID]
	for i := 0; i < len(entity.Children); i++ {
		r.GetMenuChildren(entity.Children[i], tree)
	}
}
