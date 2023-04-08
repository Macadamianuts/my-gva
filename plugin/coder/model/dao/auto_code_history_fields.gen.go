// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"gva-lbx/plugin/coder/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newAutoCodeHistoryField(db *gorm.DB, opts ...gen.DOOption) autoCodeHistoryField {
	_autoCodeHistoryField := autoCodeHistoryField{}

	_autoCodeHistoryField.autoCodeHistoryFieldDo.UseDB(db, opts...)
	_autoCodeHistoryField.autoCodeHistoryFieldDo.UseModel(&model.AutoCodeHistoryField{})

	tableName := _autoCodeHistoryField.autoCodeHistoryFieldDo.TableName()
	_autoCodeHistoryField.ALL = field.NewAsterisk(tableName)
	_autoCodeHistoryField.ID = field.NewUint(tableName, "id")
	_autoCodeHistoryField.CreatedAt = field.NewTime(tableName, "created_at")
	_autoCodeHistoryField.UpdatedAt = field.NewTime(tableName, "updated_at")
	_autoCodeHistoryField.DeletedAt = field.NewTime(tableName, "deleted_at")
	_autoCodeHistoryField.IsDelete = field.NewField(tableName, "is_delete")
	_autoCodeHistoryField.Created = field.NewUint64(tableName, "created")
	_autoCodeHistoryField.Updated = field.NewUint64(tableName, "updated")
	_autoCodeHistoryField.Deleted = field.NewUint64(tableName, "deleted")
	_autoCodeHistoryField.AutoCodeHistoryID = field.NewUint(tableName, "auto_code_history_id")
	_autoCodeHistoryField.Name = field.NewString(tableName, "name")
	_autoCodeHistoryField.Type = field.NewString(tableName, "type")
	_autoCodeHistoryField.Json = field.NewString(tableName, "json")
	_autoCodeHistoryField.Description = field.NewString(tableName, "description")
	_autoCodeHistoryField.Size = field.NewString(tableName, "size")
	_autoCodeHistoryField.Where = field.NewString(tableName, "where")
	_autoCodeHistoryField.Column = field.NewString(tableName, "column")
	_autoCodeHistoryField.Comment = field.NewString(tableName, "comment")
	_autoCodeHistoryField.Sort = field.NewBool(tableName, "sort")
	_autoCodeHistoryField.ErrorText = field.NewString(tableName, "error_text")
	_autoCodeHistoryField.Dictionary = field.NewString(tableName, "dictionary")
	_autoCodeHistoryField.Require = field.NewBool(tableName, "require")
	_autoCodeHistoryField.Clearable = field.NewBool(tableName, "clearable")

	_autoCodeHistoryField.fillFieldMap()

	return _autoCodeHistoryField
}

type autoCodeHistoryField struct {
	autoCodeHistoryFieldDo autoCodeHistoryFieldDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Time
	IsDelete          field.Field
	Created           field.Uint64
	Updated           field.Uint64
	Deleted           field.Uint64
	AutoCodeHistoryID field.Uint
	Name              field.String
	Type              field.String
	Json              field.String
	Description       field.String
	Size              field.String
	Where             field.String
	Column            field.String
	Comment           field.String
	Sort              field.Bool
	ErrorText         field.String
	Dictionary        field.String
	Require           field.Bool
	Clearable         field.Bool

	fieldMap map[string]field.Expr
}

func (a autoCodeHistoryField) Table(newTableName string) *autoCodeHistoryField {
	a.autoCodeHistoryFieldDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a autoCodeHistoryField) As(alias string) *autoCodeHistoryField {
	a.autoCodeHistoryFieldDo.DO = *(a.autoCodeHistoryFieldDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *autoCodeHistoryField) updateTableName(table string) *autoCodeHistoryField {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewTime(table, "deleted_at")
	a.IsDelete = field.NewField(table, "is_delete")
	a.Created = field.NewUint64(table, "created")
	a.Updated = field.NewUint64(table, "updated")
	a.Deleted = field.NewUint64(table, "deleted")
	a.AutoCodeHistoryID = field.NewUint(table, "auto_code_history_id")
	a.Name = field.NewString(table, "name")
	a.Type = field.NewString(table, "type")
	a.Json = field.NewString(table, "json")
	a.Description = field.NewString(table, "description")
	a.Size = field.NewString(table, "size")
	a.Where = field.NewString(table, "where")
	a.Column = field.NewString(table, "column")
	a.Comment = field.NewString(table, "comment")
	a.Sort = field.NewBool(table, "sort")
	a.ErrorText = field.NewString(table, "error_text")
	a.Dictionary = field.NewString(table, "dictionary")
	a.Require = field.NewBool(table, "require")
	a.Clearable = field.NewBool(table, "clearable")

	a.fillFieldMap()

	return a
}

func (a *autoCodeHistoryField) WithContext(ctx context.Context) IAutoCodeHistoryFieldDo {
	return a.autoCodeHistoryFieldDo.WithContext(ctx)
}

func (a autoCodeHistoryField) TableName() string { return a.autoCodeHistoryFieldDo.TableName() }

func (a autoCodeHistoryField) Alias() string { return a.autoCodeHistoryFieldDo.Alias() }

func (a *autoCodeHistoryField) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *autoCodeHistoryField) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 22)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["is_delete"] = a.IsDelete
	a.fieldMap["created"] = a.Created
	a.fieldMap["updated"] = a.Updated
	a.fieldMap["deleted"] = a.Deleted
	a.fieldMap["auto_code_history_id"] = a.AutoCodeHistoryID
	a.fieldMap["name"] = a.Name
	a.fieldMap["type"] = a.Type
	a.fieldMap["json"] = a.Json
	a.fieldMap["description"] = a.Description
	a.fieldMap["size"] = a.Size
	a.fieldMap["where"] = a.Where
	a.fieldMap["column"] = a.Column
	a.fieldMap["comment"] = a.Comment
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["error_text"] = a.ErrorText
	a.fieldMap["dictionary"] = a.Dictionary
	a.fieldMap["require"] = a.Require
	a.fieldMap["clearable"] = a.Clearable
}

func (a autoCodeHistoryField) clone(db *gorm.DB) autoCodeHistoryField {
	a.autoCodeHistoryFieldDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a autoCodeHistoryField) replaceDB(db *gorm.DB) autoCodeHistoryField {
	a.autoCodeHistoryFieldDo.ReplaceDB(db)
	return a
}

type autoCodeHistoryFieldDo struct{ gen.DO }

type IAutoCodeHistoryFieldDo interface {
	gen.SubQuery
	Debug() IAutoCodeHistoryFieldDo
	WithContext(ctx context.Context) IAutoCodeHistoryFieldDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAutoCodeHistoryFieldDo
	WriteDB() IAutoCodeHistoryFieldDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAutoCodeHistoryFieldDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAutoCodeHistoryFieldDo
	Not(conds ...gen.Condition) IAutoCodeHistoryFieldDo
	Or(conds ...gen.Condition) IAutoCodeHistoryFieldDo
	Select(conds ...field.Expr) IAutoCodeHistoryFieldDo
	Where(conds ...gen.Condition) IAutoCodeHistoryFieldDo
	Order(conds ...field.Expr) IAutoCodeHistoryFieldDo
	Distinct(cols ...field.Expr) IAutoCodeHistoryFieldDo
	Omit(cols ...field.Expr) IAutoCodeHistoryFieldDo
	Join(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo
	Group(cols ...field.Expr) IAutoCodeHistoryFieldDo
	Having(conds ...gen.Condition) IAutoCodeHistoryFieldDo
	Limit(limit int) IAutoCodeHistoryFieldDo
	Offset(offset int) IAutoCodeHistoryFieldDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoCodeHistoryFieldDo
	Unscoped() IAutoCodeHistoryFieldDo
	Create(values ...*model.AutoCodeHistoryField) error
	CreateInBatches(values []*model.AutoCodeHistoryField, batchSize int) error
	Save(values ...*model.AutoCodeHistoryField) error
	First() (*model.AutoCodeHistoryField, error)
	Take() (*model.AutoCodeHistoryField, error)
	Last() (*model.AutoCodeHistoryField, error)
	Find() ([]*model.AutoCodeHistoryField, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoCodeHistoryField, err error)
	FindInBatches(result *[]*model.AutoCodeHistoryField, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AutoCodeHistoryField) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAutoCodeHistoryFieldDo
	Assign(attrs ...field.AssignExpr) IAutoCodeHistoryFieldDo
	Joins(fields ...field.RelationField) IAutoCodeHistoryFieldDo
	Preload(fields ...field.RelationField) IAutoCodeHistoryFieldDo
	FirstOrInit() (*model.AutoCodeHistoryField, error)
	FirstOrCreate() (*model.AutoCodeHistoryField, error)
	FindByPage(offset int, limit int) (result []*model.AutoCodeHistoryField, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAutoCodeHistoryFieldDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a autoCodeHistoryFieldDo) Debug() IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Debug())
}

func (a autoCodeHistoryFieldDo) WithContext(ctx context.Context) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a autoCodeHistoryFieldDo) ReadDB() IAutoCodeHistoryFieldDo {
	return a.Clauses(dbresolver.Read)
}

func (a autoCodeHistoryFieldDo) WriteDB() IAutoCodeHistoryFieldDo {
	return a.Clauses(dbresolver.Write)
}

func (a autoCodeHistoryFieldDo) Session(config *gorm.Session) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Session(config))
}

func (a autoCodeHistoryFieldDo) Clauses(conds ...clause.Expression) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a autoCodeHistoryFieldDo) Returning(value interface{}, columns ...string) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a autoCodeHistoryFieldDo) Not(conds ...gen.Condition) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a autoCodeHistoryFieldDo) Or(conds ...gen.Condition) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a autoCodeHistoryFieldDo) Select(conds ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a autoCodeHistoryFieldDo) Where(conds ...gen.Condition) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a autoCodeHistoryFieldDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAutoCodeHistoryFieldDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a autoCodeHistoryFieldDo) Order(conds ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a autoCodeHistoryFieldDo) Distinct(cols ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a autoCodeHistoryFieldDo) Omit(cols ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a autoCodeHistoryFieldDo) Join(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a autoCodeHistoryFieldDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a autoCodeHistoryFieldDo) RightJoin(table schema.Tabler, on ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a autoCodeHistoryFieldDo) Group(cols ...field.Expr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a autoCodeHistoryFieldDo) Having(conds ...gen.Condition) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a autoCodeHistoryFieldDo) Limit(limit int) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a autoCodeHistoryFieldDo) Offset(offset int) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a autoCodeHistoryFieldDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a autoCodeHistoryFieldDo) Unscoped() IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Unscoped())
}

func (a autoCodeHistoryFieldDo) Create(values ...*model.AutoCodeHistoryField) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a autoCodeHistoryFieldDo) CreateInBatches(values []*model.AutoCodeHistoryField, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a autoCodeHistoryFieldDo) Save(values ...*model.AutoCodeHistoryField) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a autoCodeHistoryFieldDo) First() (*model.AutoCodeHistoryField, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoCodeHistoryField), nil
	}
}

func (a autoCodeHistoryFieldDo) Take() (*model.AutoCodeHistoryField, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoCodeHistoryField), nil
	}
}

func (a autoCodeHistoryFieldDo) Last() (*model.AutoCodeHistoryField, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoCodeHistoryField), nil
	}
}

func (a autoCodeHistoryFieldDo) Find() ([]*model.AutoCodeHistoryField, error) {
	result, err := a.DO.Find()
	return result.([]*model.AutoCodeHistoryField), err
}

func (a autoCodeHistoryFieldDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoCodeHistoryField, err error) {
	buf := make([]*model.AutoCodeHistoryField, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a autoCodeHistoryFieldDo) FindInBatches(result *[]*model.AutoCodeHistoryField, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a autoCodeHistoryFieldDo) Attrs(attrs ...field.AssignExpr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a autoCodeHistoryFieldDo) Assign(attrs ...field.AssignExpr) IAutoCodeHistoryFieldDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a autoCodeHistoryFieldDo) Joins(fields ...field.RelationField) IAutoCodeHistoryFieldDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a autoCodeHistoryFieldDo) Preload(fields ...field.RelationField) IAutoCodeHistoryFieldDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a autoCodeHistoryFieldDo) FirstOrInit() (*model.AutoCodeHistoryField, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoCodeHistoryField), nil
	}
}

func (a autoCodeHistoryFieldDo) FirstOrCreate() (*model.AutoCodeHistoryField, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoCodeHistoryField), nil
	}
}

func (a autoCodeHistoryFieldDo) FindByPage(offset int, limit int) (result []*model.AutoCodeHistoryField, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a autoCodeHistoryFieldDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a autoCodeHistoryFieldDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a autoCodeHistoryFieldDo) Delete(models ...*model.AutoCodeHistoryField) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *autoCodeHistoryFieldDo) withDO(do gen.Dao) *autoCodeHistoryFieldDo {
	a.DO = *do.(*gen.DO)
	return a
}
