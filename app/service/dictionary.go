package service

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/common"
)

var Dictionary = new(dictionary)

type dictionary struct{}

// Create 创建字典
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Create(ctx context.Context, info request.DictionaryCreate) error {
	query := dao.Q.WithContext(ctx).Dictionary
	_, err := query.Where(dao.Dictionary.Type.Eq(info.Type)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的字典类型，不允许创建!")
	}
	create := info.Create()
	err = query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// First 获取字典数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) First(ctx context.Context, info request.DictionaryFirst) (entity *model.Dictionary, err error) {
	query := dao.Q.WithContext(ctx).Dictionary
	entity, err = query.Scopes(info.First()).Preload(dao.Dictionary.Details).First()
	if err != nil {
		return nil, errors.Wrap(err, common.ErrorFirst)
	}
	return entity, nil
}

// Update 更新字典
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Update(ctx context.Context, info request.DictionaryUpdate) error {
	query := dao.Q.WithContext(ctx).Dictionary
	first, err := query.Where(dao.Dictionary.ID.Eq(info.Id)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "字典不存在!")
	}
	if first.Type != info.Type {
		_, err = query.Where(dao.Dictionary.Type.Eq(info.Type)).First()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "存在相同的字典名，不允许创建!")
		}
	}
	update := info.Update()
	_, err = query.Where(dao.Dictionary.ID.Eq(info.Id)).Updates(&update)
	if err != nil {
		return errors.Wrap(err, common.ErrorUpdated)
	}
	return nil
}

// Delete 删除字典
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Delete(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).Dictionary
	entity, err := query.Where(dao.Dictionary.ID.Eq(info.Id)).Preload(dao.Dictionary.Details).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, common.ErrorDeleted)
	}
	_, err = query.Select(field.AssociationFields).Delete(entity)
	if err != nil {
		return errors.Wrap(err, common.ErrorDeleted)
	}
	return nil
}

// List 获取字典列表数据
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) List(ctx context.Context, info request.DictionarySearch) (entities []*model.Dictionary, count int64, err error) {
	query := dao.Q.WithContext(ctx).Dictionary
	query = query.Scopes(info.Search())
	count, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorListCount)
	}
	entities, err = query.Scopes(info.Paginate()).Preload(dao.Dictionary.Details).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, common.ErrorList)
	}
	return
}
