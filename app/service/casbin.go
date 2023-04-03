package service

import (
	"context"
	"github.com/pkg/errors"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/common"
	"gva-lbx/global"
)

var Casbin = new(_casbin)

type _casbin struct{}

// AddPolicies 添加api权限
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) AddPolicies(ctx context.Context, info request.CasbinAddPolicies) error {
	rules := info.Update()
	if err := global.Enforcer.LoadPolicy(); err != nil {
		return errors.Wrap(err, "casbin 加载策略失败!")
	}
	success, err := global.Enforcer.AddPolicies(rules)
	if !success || err != nil {
		return errors.Wrap(err, "存在相同api,添加失败,请联系管理员!")
	}
	return nil
}

// Find 获取权限列表
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) Find(ctx context.Context, info common.Role) ([]*model.Casbin, error) {
	if err := global.Enforcer.LoadPolicy(); err != nil {
		return nil, errors.Wrap(err, "casbin 加载策略失败!")
	}
	rules := global.Enforcer.GetFilteredPolicy(0, info.String())
	length := len(rules)
	entities := make([]*model.Casbin, 0, length)
	for i := 0; i < length; i++ {
		entities = append(entities, &model.Casbin{
			RoleId: info.String(),
			Path:   rules[i][1],
			Method: rules[i][2],
		})
	}
	return entities, nil
}

// Update 更新
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) Update(ctx context.Context, info request.CasbinUpdate) error {
	query := dao.Q.WithContext(ctx).Casbin
	update := info.Update()
	_, err := query.Where(dao.Casbin.Path.Eq(info.OldPath), dao.Casbin.Method.Eq(info.OldMethod)).Updates(&update)
	if err != nil {
		return errors.Wrap(err, "更新失败!")
	}
	return nil
}

// Clear 清除匹配的权限
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) Clear(ctx context.Context, info common.Role) error {
	if err := global.Enforcer.LoadPolicy(); err != nil {
		return errors.Wrap(err, "casbin 加载策略失败!")
	}
	success, err := global.Enforcer.RemoveFilteredPolicy(0, info.String())
	if !success || err != nil {
		return errors.Wrap(err, "清除匹配的权限失败!")
	}
	return nil
}
