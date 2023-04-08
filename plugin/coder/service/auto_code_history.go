package service

import (
	"context"
	"github.com/pkg/errors"
	"gva-lbx/common"
	"gva-lbx/plugin/coder/model"
	"gva-lbx/plugin/coder/model/dao"
	"gva-lbx/plugin/coder/model/request"
)

var AutoCodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// Create 创建
func (s *autoCodeHistory) Create(ctx context.Context, info request.AutoCodeCreate) error {
	create := info.Create()
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	err := query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// First 根据id获取记录
func (s *autoCodeHistory) First(ctx context.Context, info common.GormId) (entity *model.AutoCodeHistory, err error) {
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	entity, err = query.Where(dao.AutoCodeHistory.ID.Eq(info.Id)).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return nil, nil
}

// Repeat 检测重复
func (s *autoCodeHistory) Repeat(ctx context.Context, info request.AutoCodeHistoryRepeat) bool {
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	query = query.Where(dao.AutoCodeHistory.Plugin.Eq(info.Plugin), dao.AutoCodeHistory.Struct.Eq(info.Struct))
	count, err := query.Count()
	if err != nil {
		return false
	}
	return count > 0
}

// Rollback TODO 回滚
func (s *autoCodeHistory) Rollback(ctx context.Context, info request.AutoCodeHistoryRollback) error {
	return nil
}

// Delete 删除
func (s *autoCodeHistory) Delete(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	_, err := query.Where(dao.AutoCodeHistory.ID.Eq(info.Id)).Delete()
	if err != nil {
		return errors.Wrap(err, common.ErrorBatchDeleted)
	}
	return nil
}

// Deletes 批量删除
func (s *autoCodeHistory) Deletes(ctx context.Context, info common.GormIds) error {
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	_, err := query.Where(dao.AutoCodeHistory.ID.In(info.Ids...)).Delete()
	if err != nil {
		return errors.Wrap(err, common.ErrorBatchDeleted)
	}
	return nil
}

// List 列表
func (s *autoCodeHistory) List(ctx context.Context, info request.AutoCodeHistoryList) (entities []*model.AutoCodeHistory, count int64, err error) {
	query := dao.Q.WithContext(ctx).AutoCodeHistory
	query = query.Scopes(info.Search())
	count, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorListCount)
	}
	entities, err = query.Scopes(info.Paginate()).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorList)
	}
	return entities, count, nil
}
