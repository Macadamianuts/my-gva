package service

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service/internal"
	"gva-lbx/common"
)

var Menu = new(menu)

type menu struct{}

// Create 新建菜单
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) Create(ctx context.Context, info request.MenuCreate) error {
	query := dao.Q.WithContext(ctx).Menu
	create := info.Create()
	_, err := query.Where(dao.Menu.Name.Eq(info.Name)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在重复菜单名，请修改菜单名!")
	}
	err = query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// First 根据id获取菜单数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) First(ctx context.Context, info common.GormId) (entity *model.Menu, err error) {
	query := dao.Q.WithContext(ctx).Menu
	entity, err = query.Where(dao.Menu.ID.Eq(info.Id)).Preload(dao.Menu.Parameters).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return entity, nil
}

// FindByRole 根据角色id获取菜单列表
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) FindByRole(ctx context.Context, info common.Role) ([]*model.Menu, error) {
	entity, err := dao.Q.WithContext(ctx).Role.Where(dao.Role.ID.Eq(info.RoleId)).Preload(dao.Role.Menus).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "角色不存在!")
	}
	return entity.Menus, nil
}

// Update 更新
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) Update(ctx context.Context, info request.MenuUpdate) error {
	query := dao.Q.WithContext(ctx).Menu
	entity, err := query.Where(dao.Menu.ID.Eq(info.Id)).Preload(field.Associations).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "菜单不存在!")
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		if entity.Name != info.Name {
			_, err = tx.WithContext(ctx).Menu.Where(dao.Menu.ID.Neq(info.Id), dao.Menu.Name.Eq(info.Name)).First()
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Wrap(err, "存在相同name修改失败!")
			}
		}
		err = tx.Menu.Parameters.Model(entity).Clear()
		if err != nil {
			return errors.Wrap(err, "菜单参数关联清空失败!")
		}
		_, err = tx.WithContext(ctx).MenuParameter.Unscoped().Delete(entity.Parameters...)
		if err != nil {
			return errors.Wrap(err, "菜单参数删除失败!")
		}
		update := info.Update()
		_, err = tx.WithContext(ctx).Menu.Where(dao.Menu.ID.Eq(info.Id)).Omit(field.AssociationFields).Updates(update)
		if err != nil {
			return errors.Wrap(err, common.ErrorUpdated)
		}
		return nil
	})
}

// Delete 删除
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) Delete(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).Menu
	entity, err := query.Where(dao.Menu.ParentId.Eq(info.Id)).Preload(dao.Menu.Parameters).Preload(dao.Menu.Roles).First()
	if errors.Is(err, gorm.ErrRecordNotFound) { // 不存在子菜单
		entity, err = query.Where(dao.Menu.ID.Eq(info.Id)).Preload(dao.Menu.Parameters).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "菜单不存在!")
		}
		err = dao.Menu.Parameters.Model(entity).Delete(entity.Parameters...)
		if err != nil {
			return errors.Wrap(err, "菜单参数关联引用删除失败!")
		}
		err = dao.Menu.Roles.Model(entity).Delete(entity.Roles...)
		if err != nil {
			return errors.Wrap(err, "菜单角色关联引用删除失败!")
		}
		_, err = dao.Q.WithContext(ctx).MenuParameter.Unscoped().Delete(entity.Parameters...)
		if err != nil {
			return errors.Wrap(err, "菜单参数删除失败!")
		}
		return nil
	}
	return errors.Wrap(err, "此菜单存在子菜单不可删除!")
}

// List 获取菜单列表
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) List(ctx context.Context) ([]*model.Menu, error) {
	tree, err := internal.Menu.GetTreeMap(ctx)
	if err != nil {
		return nil, err
	}
	entities := tree[0]
	for i := 0; i < len(entities); i++ {
		internal.Menu.GetChildren(entities[i], tree)
	}
	return entities, nil
}

// UserDefaultRouter 用户默认角色的默认菜单检查
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *menu) UserDefaultRouter(ctx context.Context, user *model.User) {
	query := dao.Q.WithContext(ctx).RolesMenus
	var menuIds []uint
	err := query.Where(dao.RolesMenus.RoleId.Eq(user.RoleId)).Pluck(dao.RolesMenus.MenuId, &menuIds)
	if err != nil {
		return
	}
	_, err = dao.Q.WithContext(ctx).Menu.Where(dao.Menu.Name.Eq(user.Role.DefaultRouter), dao.Menu.ID.In(menuIds...)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Role.DefaultRouter = "404"
	}
}
