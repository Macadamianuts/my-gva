package service

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/common"
	"gva-lbx/global"
)

var Api = new(api)

type api struct{}

// Create 创建 api
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Create(ctx context.Context, info request.ApiCreate) error {
	query := dao.Q.WithContext(ctx).Api
	_, err := query.Where(dao.Api.Path.Eq(info.Path), dao.Api.Method.Eq(info.Method)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "存在相同api")
	}
	create := info.Create()
	err = query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// First 获取api数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) First(ctx context.Context, info common.GormId) (entity *model.Api, err error) {
	query := dao.Q.WithContext(ctx).Api
	entity, err = query.Where(dao.Api.ID.Eq(info.Id)).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return entity, nil
}

// Update 更新 api
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Update(ctx context.Context, info request.ApiUpdate) error {
	query := dao.Q.WithContext(ctx).Api
	entity, err := query.Where(dao.Api.ID.Eq(info.Id)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "api不存在!")
	}
	if entity.Path != info.Path || entity.Method != info.Method {
		_, err = query.Where(dao.Api.Path.Eq(info.Path), dao.Api.Method.Eq(info.Method)).First()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "存在相同api路径")
		}
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		update := info.Update()
		_, err = tx.WithContext(ctx).Api.Where(dao.Api.ID.Eq(info.Id)).Updates(&update)
		if err != nil {
			return errors.Wrap(err, common.ErrorUpdated)
		}
		_, err = tx.WithContext(ctx).Casbin.Where(dao.Casbin.Path.Eq(entity.Path), dao.Casbin.Method.Eq(entity.Method)).Updates(info.Casbin())
		if err != nil {
			return errors.Wrap(err, "casbin规则更新失败!")
		}
		return nil
	})
}

// Delete 删除 api
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Delete(ctx context.Context, info common.GormId) error {
	entity, err := dao.Q.WithContext(ctx).Api.Where(dao.Api.ID.Eq(info.Id)).First()
	if err != nil {
		return errors.Wrap(err, common.ErrorFirst)
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		_, err = tx.WithContext(ctx).Api.Delete(entity)
		if err != nil {
			return errors.Wrap(err, common.ErrorDeleted)
		}
		_, err = tx.WithContext(ctx).Casbin.Where(dao.Casbin.Path.Eq(entity.Path), dao.Casbin.Method.Eq(entity.Method)).Delete()
		if err != nil {
			return errors.Wrap(err, "casbin规则删除失败!")
		}
		return nil
	})
}

// Deletes 批量删除 api
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Deletes(ctx context.Context, info common.GormIds) error {
	query := dao.Q.WithContext(ctx).Api
	entities, err := query.Where(dao.Api.ID.In(info.Ids...)).Find()
	if err != nil {
		return errors.Wrap(err, common.ErrorFind)
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		for i := 0; i < len(entities); i++ {
			_, err = tx.WithContext(ctx).Api.Delete(entities[i])
			if err != nil {
				return errors.Wrap(err, common.ErrorDeleted)
			}
			_, err = tx.WithContext(ctx).Casbin.Where(dao.Casbin.Path.Eq(entities[i].Path), dao.Casbin.Method.Eq(entities[i].Method)).Delete()
			if err != nil {
				return errors.Wrap(err, "casbin规则删除失败!")
			}
		}
		err = global.Enforcer.InvalidateCache()
		if err != nil {
			return errors.Wrap(err, "casbin规则缓存无效失败!")
		}
		return nil
	})
}

// List 获取api列表分页数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) List(ctx context.Context, info request.ApiSearch) (entities []*model.Api, count int64, err error) {
	query := dao.Q.WithContext(ctx).Api
	query = query.Scopes(info.Search())
	count, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorListCount)
	}
	entities, err = query.Scopes(info.Paginate(), info.Order()).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorList)
	}
	return entities, count, nil
}

// All 获取api列表数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) All(ctx context.Context) (entities []*model.Api, err error) {
	query := dao.Q.WithContext(ctx).Api
	entities, err = query.Find()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorList)
	}
	return entities, nil
}
