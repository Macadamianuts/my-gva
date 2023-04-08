package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/app/model/request"
	"gva-lbx/app/model/response"
	"gva-lbx/common"
	"gva-lbx/component/jwt"
	"gva-lbx/global"
)

var User = new(user)

type user struct{}

// Create 创建用户
func (s *user) Create(ctx context.Context, info request.UserCreate) error {
	query := dao.Q.WithContext(ctx).User
	_, err := query.Where(dao.User.Username.Eq(info.Username)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "用户名已注册!")
	}
	create := info.Create()
	err = create.EncryptedPassword()
	if err != nil {
		return errors.Wrap(err, "加密密码失败!")
	}
	err = query.Create(&create)
	if err != nil {
		return errors.Wrap(err, common.ErrorCreated)
	}
	return nil
}

// Login 登录
func (s *user) Login(ctx context.Context, info request.UserLogin) (*response.UserLogin, error) {
	query := dao.Q.WithContext(ctx).User
	entity, err := query.Where(dao.User.Username.Eq(info.Username)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "用户不存在!")
	}
	err = entity.CompareHashAndPassword(info.Password)
	if err != nil {
		return nil, err
	}
	Menu.UserDefaultRouter(ctx, entity)
	return s.Token(ctx, entity)
}

// First 获取用户信息
func (s *user) First(ctx context.Context, info request.UserFirst) (entity *model.User, err error) {
	query := dao.Q.WithContext(ctx).User
	entity, err = query.Scopes(info.First()).Preload(dao.User.Role).Preload(dao.User.Roles).First()
	if err != nil {
		return nil, errors.Wrap(err, "用户不存在!")
	}
	Menu.UserDefaultRouter(ctx, entity)
	return entity, nil
}

// Update 设置用户信息
func (s *user) Update(ctx context.Context, info request.UserUpdate) error {
	query := dao.Q.WithContext(ctx).User
	update := info.Update()
	_, err := query.Where(dao.User.ID.Eq(info.Id)).Updates(&update)
	if err != nil {
		return errors.Wrap(err, common.ErrorUpdated)
	}
	return nil
}

// ResetPassword 重置密码
func (s *user) ResetPassword(ctx context.Context, info common.GormId) error {
	query := dao.Q.WithContext(ctx).User
	first, err := query.Where(dao.User.ID.Eq(info.Id)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "用户不存在!")
	}
	first.Password = "123456"
	err = first.EncryptedPassword()
	if err != nil {
		return err
	}
	_, err = query.Where(dao.User.ID.Eq(info.Id)).UpdateColumn(dao.User.Password, first.Password)
	if err != nil {
		return errors.Wrap(err, "重置密码失败!")
	}
	return err
}

// ChangePassword 修改密码
func (s *user) ChangePassword(ctx context.Context, info request.UserChangePassword) error {
	query := dao.Q.WithContext(ctx).User
	first, err := query.WithContext(ctx).Where(dao.User.ID.Eq(info.UserId)).First()
	if err != nil {
		return errors.Wrap(err, "用户不存在!")
	}
	err = first.CompareHashAndPassword(info.Password)
	if err != nil {
		return err
	}
	first.Password = info.NewPassword
	err = first.EncryptedPassword()
	if err != nil {
		return errors.Wrap(err, "加密密码失败!")
	}
	_, err = query.WithContext(ctx).Where(dao.User.ID.Eq(info.UserId)).UpdateColumn(dao.User.Password, first.Password)
	if err != nil {
		return errors.Wrap(err, "修改密码失败!")
	}
	return nil
}

// SetRole 设置用户角色
func (s *user) SetRole(ctx context.Context, info request.UserSetRole) error {
	_, err := dao.Q.WithContext(ctx).UsersRoles.Where(dao.UsersRoles.UserId.Eq(info.UserId), dao.UsersRoles.RoleId.Eq(info.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "该用户无此角色")
	}
	query := dao.Q.WithContext(ctx).User
	_, err = query.Where(dao.User.ID.Eq(info.UserId)).UpdateColumn(dao.User.RoleId, info.RoleId)
	if err != nil {
		return errors.Wrap(err, "设置用户角色失败!")
	}
	return nil
}

// SetRoles 设置用户多角色
func (s *user) SetRoles(ctx context.Context, info request.UserSetRoles) error {
	entity, err := dao.Q.WithContext(ctx).User.Where(dao.User.ID.Eq(info.UserId)).Preload(dao.User.Roles).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "用户不存在!")
	}
	roles, err := dao.Q.WithContext(ctx).Role.Where(dao.Role.ID.In(info.RoleIds...)).Find()
	if err != nil {
		return errors.Wrap(err, "获取角色列表失败!")
	}
	return dao.Q.Transaction(func(tx *dao.Query) error {
		err = tx.User.Roles.Model(entity).Clear()
		if err != nil {
			return errors.Wrap(err, "清空用户角色失败!")
		}
		err = tx.User.Roles.Model(entity).Append(roles...)
		if err != nil {
			return errors.Wrap(err, "设置用户多角色失败!")
		}
		return nil
	})
}

// Delete 删除用户
func (s *user) Delete(ctx context.Context, info common.GormId) error {
	return dao.Q.Transaction(func(tx *dao.Query) error {
		_, err := tx.User.WithContext(ctx).Where(dao.User.ID.Eq(info.Id)).Delete()
		if err != nil {
			return errors.Wrap(err, common.ErrorDeleted)
		}
		_, err = tx.UsersRoles.WithContext(ctx).Where(dao.UsersRoles.UserId.Eq(info.Id)).Delete()
		if err != nil {
			return errors.Wrap(err, "清空用户角色失败!")
		}
		return nil
	})
}

// Deletes 批量删除用户
func (s *user) Deletes(ctx context.Context, info common.GormIds) error {
	return dao.Q.Transaction(func(tx *dao.Query) error {
		_, err := tx.User.WithContext(ctx).Where(dao.User.ID.In(info.Ids...)).Delete()
		if err != nil {
			return errors.Wrap(err, common.ErrorBatchDeleted)
		}
		_, err = tx.UsersRoles.WithContext(ctx).Where(dao.UsersRoles.UserId.In(info.Ids...)).Delete()
		if err != nil {
			return errors.Wrap(err, "清空用户角色失败!")
		}
		return nil
	})
}

// List 获取用户列表数据
func (s *user) List(ctx context.Context, info request.UserSearch) (entities []*model.User, total int64, err error) {
	query := dao.Q.WithContext(ctx).User
	query = query.Scopes(info.Search())
	total, err = query.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, "获取数据数量失败!")
	}
	entities, err = query.Scopes(info.Paginate()).Preload(dao.User.Role).Preload(dao.User.Roles).Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, "获取数据列表失败!")
	}
	return entities, total, nil
}

// Token 用户登录成功获取的token
func (s *user) Token(ctx context.Context, user *model.User) (*response.UserLogin, error) {
	base := request.NewClaims(user)
	claims := jwt.NewClaims(base)
	token, err := jwt.NewJwt().Create(claims)
	if err != nil {
		return nil, err
	}
	success := response.UserLogin{
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix(),
		User:      user,
	}
	if !global.Config.System.UseMultipoint {
		return &success, nil
	}
	var redisJwt string
	redisJwt, err = Jwt.GetRedisJWT(ctx, user.Username)
	if err != nil || err == redis.Nil {
		err = Jwt.SetRedisJWT(ctx, token, user.Username)
		if err != nil {
			return nil, err
		}
		return &success, nil
	}
	err = Jwt.JsonInBlacklist(ctx, redisJwt)
	if err != nil {
		return nil, err
	}
	err = Jwt.SetRedisJWT(ctx, token, user.Username)
	if err != nil {
		return nil, err
	}
	return &success, nil
}
