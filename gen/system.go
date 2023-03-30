package main

import (
	"gorm.io/gen"
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
	g.ApplyBasic()

	g.Execute()
}
