package main

import (
	"os"
	"path/filepath"

	"gorm.io/gen"
)

func main() {
	pwd, _ := os.Getwd()
	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Join(pwd, "plugin", "{{.Name}}", "model", "dao"),
		OutFile: "query.go",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldCoverable: true,
		WithUnitTest: false,
	})
	g.ApplyBasic()
	g.Execute()
}
