package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"gorm.io/gorm"
	"gva-lbx/app/model"
	"gva-lbx/app/model/dao"
	"gva-lbx/global"
	"net/http"
	"strings"
)

var Router = new(_router)

type _router struct{}

// Initialization 初始化
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (r *_router) Initialization(engine *gin.Engine) {
	parser := swag.New()
	err := parser.ParseAPIMultiSearchDir([]string{"."}, "main.go", 0)
	if err != nil {
		fmt.Printf("[core][router][err:%s]解析api包注释失败!\n", err)
		return
	}
	ctx := context.Background()
	swagger := parser.GetSwagger()
	routes := engine.Routes()
	paths := swagger.SwaggerProps.Paths.Paths
	err = global.Db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.Api{}).UpdateColumn("IsEffective", false).Error
	if err != nil {
		fmt.Printf("[core][router][err:%s]设置全部api设置无效失败!\n", err)
		return
	}
	for i := 0; i < len(routes); i++ {
		if routes[i].Path == "/form-generator/*filepath" || routes[i].Path == "/swagger/*any" || routes[i].Path == "/files/*filepath" {
			continue
		}
		entity := model.Api{
			Path:        routes[i].Path,
			Method:      routes[i].Method,
			IsEffective: true,
		}
		var description, apiGroup string
		query := dao.Q.WithContext(ctx).Api
		_, err = query.Where(dao.Api.Path.Eq(routes[i].Path), dao.Api.Method.Eq(routes[i].Method)).First()
		value, ok := paths[routes[i].Path]
		if ok {
			switch routes[i].Method {
			case http.MethodGet:
				apiGroup = strings.Join(value.Get.Tags, ",")
				if value.Get.Description != "" {
					description = value.Get.Summary
				} else {
					description = value.Get.Description
				}
			case http.MethodPut:
				apiGroup = strings.Join(value.Put.Tags, ",")
				if value.Put.Description == "" {
					description = value.Put.Summary
				} else {
					description = value.Put.Description
				}
				if value.Put.Summary != "" {
				}
			case http.MethodPost:
				apiGroup = strings.Join(value.Post.Tags, ",")
				if value.Post.Description == "" {
					description = value.Post.Summary
				} else {
					description = value.Post.Description
				}
			case http.MethodPatch:
				apiGroup = strings.Join(value.Patch.Tags, ",")
				if value.Patch.Description == "" {
					description = value.Patch.Summary
				} else {
					description = value.Patch.Description
				}
			case http.MethodDelete:
				apiGroup = strings.Join(value.Delete.Tags, ",")
				if value.Delete.Description == "" {
					description = value.Delete.Summary
				} else {
					description = value.Delete.Description
				}
			}
			entity.ApiGroup = apiGroup
			entity.Description = description
		}
		if errors.Is(err, gorm.ErrRecordNotFound) { // api 不存在则创建
			err = query.Create(&entity)
			if err != nil {
				fmt.Printf("[core][router][err:%v]api记录创建失败!\n", err)
				return
			}
		} else { // api 存在则更新
			_, err = query.Where(dao.Api.Path.Eq(routes[i].Path), dao.Api.Method.Eq(routes[i].Method)).Updates(model.Api{Description: description, ApiGroup: apiGroup, IsEffective: true})
			if err != nil {
				fmt.Printf("[core][router][err:%v]api记录更新失败!\n", err)
				return
			}
		}
	}

	err = global.Db.WithContext(ctx).Where("is_effective = ?", false).Delete(&model.Api{}).Error
	if err != nil {
		fmt.Printf("[core][router][err:%s]设置无效api软删除失败!\n", err)
	}
}
