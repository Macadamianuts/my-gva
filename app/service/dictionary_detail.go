package service

import (
	"context"
	"github.com/pkg/errors"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/common"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

// Create 创建字典详情
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Create(ctx context.Context, info request.DictionaryDetailCreate) error {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	create := info.Create()
	err := query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// First 获取字典数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) First(ctx context.Context, info common.GormId) (entity *model.DictionaryDetail, err error) {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	entity, err = query.Where(dao.DictionaryDetail.ID.Eq(info.Id)).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return entity, nil
}

// Update 更新字典详情
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Update(ctx context.Context, info request.DictionaryDetailUpdate) error {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	update := info.Update()
	_, err := query.Where(dao.DictionaryDetail.ID.Eq(info.Id)).Updates(&update)
	if err != nil {
		return errors.Wrap(err, common.ErrorUpdated)
	}
	return nil
}

// Delete 删除字典详情
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Delete(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	_, err := query.Where(dao.DictionaryDetail.ID.Eq(info.Id)).Delete()
	if err != nil {
		return errors.Wrap(err, common.ErrorDeleted)
	}
	return nil
}

// Deletes 批量删除字典详情
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Deletes(ctx context.Context, info common.GormIds) error {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	_, err := query.Where(dao.DictionaryDetail.ID.In(info.Ids...)).Delete()
	if err != nil {
		return errors.Wrap(err, common.ErrorBatchDeleted)
	}
	return nil
}

// List 获取字典详情列表数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) List(ctx context.Context, info request.DictionaryDetailSearch) (entities []*model.DictionaryDetail, count int64, err error) {
	query := dao.Q.WithContext(ctx).DictionaryDetail
	query = query.Scopes(info.Search())
	count, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorListCount)
	}
	entities, err = query.Scopes(info.Paginate()).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorList)
	}
	return entities, count, err
}
