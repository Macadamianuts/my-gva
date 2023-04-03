package main

import (
	"gorm.io/gen"
	"gva-lbx/app/model"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前文件夹路径
	pwd, _ := os.Getwd()

	g := gen.NewGenerator(gen.Config{
		// 文件输出的位置（根据自己的情况而定）
		OutPath: filepath.Join(pwd, "app", "model", "dao"),
		// 输出的文件名
		OutFile: "query.go",
		// 输出模式
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		FieldCoverable: true,

		WithUnitTest: false,
	})

	// 挂在自己的结构体在这里（根据自己的业务而定）
	g.ApplyBasic(
		new(model.Api),
		new(model.User),
		new(model.Menu),
		new(model.Role),
		new(model.Casbin),
		new(model.RolesMenus),
		new(model.Dictionary),
		new(model.UsersRoles),
		new(model.JwtBlacklist),
		new(model.MenuParameter),
		new(model.OperationRecord),
		new(model.RolesMenuButtons),
		new(model.DictionaryDetail),
	)

	g.Execute()
}
