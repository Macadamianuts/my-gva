package service

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/app/model/response"
	"gva-lbx/app/service/internal"
	"gva-lbx/common"
)

var Role = new(role)

type role struct{}

// Create 创建
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) Create(ctx context.Context, info request.RoleCreate) error {
	query := dao.Q.WithContext(ctx).Role
	_, err := query.Where(dao.Role.ID.Eq(info.RoleId)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("角色已存在!")
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		create := info.Create()
		err = tx.WithContext(ctx).Role.Create(&create)
		if err != nil {
			return errors.Wrap(err, common.ErrorCreated)
		}
		casbin := info.Casbin()
		err = tx.WithContext(ctx).Casbin.Create(casbin...)
		if err != nil {
			return errors.Wrap(err, "创建角色默认权限失败!")
		}
		return nil
	})
}

// AddMenus 增加menu和角色关联关系
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) AddMenus(ctx context.Context, info request.RoleAddMenus) error {
	menus, err := dao.Q.WithContext(ctx).Menu.Where(dao.Menu.ID.In(info.MenuId...)).Find()
	if err != nil {
		return errors.Wrap(err, "获取菜单失败!")
	}
	entity, err := dao.Q.WithContext(ctx).Role.Where(dao.Role.ID.Eq(info.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "角色不存在！")
	}
	err = dao.Role.Menus.Model(entity).Append(menus...)
	if err != nil {
		return errors.Wrap(err, "角色添加菜单关联失败！")
	}
	return nil
}

// Copy 复制角色
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) Copy(ctx context.Context, info request.RoleCopy) error {
	query := dao.Q.WithContext(ctx).Role
	_, err := query.Where(dao.Role.ID.Eq(info.Role.ID)).Preload(dao.Role.Menus).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("角色已存在!")
	}
	menus, err := Menu.FindByRole(ctx, info.ToCommonRole())
	if err != nil {
		return err
	}
	info.Role.Menus = menus
	rules, err := Casbin.Find(ctx, info.ToCommonRole())
	if err != nil {
		return err
	}
	for i := 0; i < len(rules); i++ {
		rules[i].RoleId = info.Role.IDString()
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		err = tx.WithContext(ctx).Role.Create(&info.Role)
		if err != nil {
			return errors.Wrap(err, "角色创建失败!")
		}
		err = tx.WithContext(ctx).Casbin.Create(rules...)
		if err != nil {
			return errors.Wrap(err, "角色权限复制失败!")
		}
		return nil
	})
}

// First 根据 id 获取数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) First(ctx context.Context, info common.GormId) (entity *model.Role, err error) {
	query := dao.Q.WithContext(ctx).Role
	entity, err = query.Where(dao.Role.ID.Eq(info.Id)).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return entity, nil
}

// Update 更新角色
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) Update(ctx context.Context, info request.RoleUpdate) error {
	query := dao.Q.WithContext(ctx).Role
	update := info.Update()
	_, err := query.Where(dao.Role.ID.Eq(info.RoleId)).Updates(&update)
	if err != nil {
		return errors.Wrap(err, common.ErrorUpdated)
	}
	return nil
}

// Delete 删除角色
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) Delete(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).Role
	entity, err := query.Where(dao.Role.ID.Eq(info.Id)).Preload(dao.Role.Users).Preload(dao.Role.Menus).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "角色不存在!")
	}
	if len(entity.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除!")
	}
	entity, err = query.Where(dao.Role.ParentId.Eq(info.Id)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		_, err = tx.WithContext(ctx).Role.Unscoped().Select(field.AssociationFields).Delete(entity)
		if err != nil {
			return errors.Wrap(err, "删除角色失败!")
		}
		roleId := entity.IDString()
		_, err = tx.WithContext(ctx).Casbin.Where(dao.Casbin.RoleId.Eq(roleId)).Delete()
		if err != nil {
			return errors.Wrap(err, "删除角色权限失败!")
		}
		return nil
	})
}

// List 获取角色列表数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) List(ctx context.Context, info request.RoleSearch) (entities []*model.Role, count int64, err error) {
	query := dao.Q.WithContext(ctx).Role
	query = query.Scopes(info.Search())
	count, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorListCount)
	}
	entities, err = query.Scopes(info.Paginate()).Preload(dao.Role.Menus).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorList)
	}
	tree, err := internal.Role.GetTreeMap(ctx)
	if err != nil {
		return nil, 0, err
	}
	length := len(entities)
	for i := 0; i < length; i++ {
		internal.Role.GetChildren(entities[i], tree)
	}
	return entities, count, nil
}

// MenuTree 获取角色动态菜单树
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *role) MenuTree(ctx context.Context, info common.Role) ([]*response.RoleMenuTree, error) {
	tree, err := internal.Role.GetMenuTreeMap(ctx, info)
	if err != nil {
		return nil, err
	}
	entities := tree[0]
	for i := 0; i < len(entities); i++ {
		internal.Role.GetMenuChildren(entities[i], tree)
	}
	return entities, nil
}
