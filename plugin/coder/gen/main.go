package main

import (
	"gorm.io/gen"
	"gva-lbx/plugin/coder/model"
	"os"
	"path/filepath"
)

// generate code
func main() {
	pwd, _ := os.Getwd()
	g := gen.NewGenerator(gen.Config{
		OutPath:        filepath.Join(pwd, "plugin", "coder", "model", "dao"),
		OutFile:        "query.go",
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldCoverable: true,
		WithUnitTest:   false,
	})

	g.ApplyBasic(
		new(model.AutoCodeHistory),
		new(model.AutoCodeHistoryField),
	)
	g.Execute()
}
