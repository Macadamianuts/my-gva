package service

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/plugin/coder/global"
	"gva-lbx/plugin/coder/model/request"
	"gva-lbx/plugin/coder/service/internal"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var AutoCodePlugin = new(autoCodePlugin)

type autoCodePlugin struct{}

// Create .
func (s *autoCodePlugin) Create(ctx context.Context, info request.AutoCodePluginCreate) error {
	root := global.Config.Server.PluginRoot() // 当前工作路径 + 插件存放文件夹
	{
		dirs, err := os.ReadDir(root)
		if err != nil {
			return errors.Wrap(err, "获取插件文件夹失败!")
		}
		for i := 0; i < len(dirs); i++ {
			if dirs[i].IsDir() {
				if info.Name == dirs[i].Name() {
					return errors.Wrap(err, "已存在同名插件")
				}
			}
		}
	} // 检测插件路径含有的插件
	plugin := global.Config.Server.PluginPath(info.Name) // 当前工作路径 + 插件存放文件夹 + 插件名
	{
		err := os.Mkdir(plugin, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "创建插件名[%s]失败!", info.Name)
		}
		newDirs := []string{"api", "model", filepath.Join("model", "request"), "gen", "router", "service"}
		for i := 0; i < len(newDirs); i++ {
			err = os.Mkdir(filepath.Join(plugin, newDirs[i]), os.ModePerm)
			if err != nil {
				return errors.Wrapf(err, "创建文件夹[%s]失败!", newDirs[i])
			}
		}
	} // 创建插件文件夹
	{
		parse, err := template.New("").Parse(string(internal.Gen))
		if err != nil {
			return errors.Wrap(err, "解析模版文件[gen.go]失败!")
		}
		var file *os.File
		path := filepath.Join(global.Config.Server.PluginRoot(), "coder", "service", "internal", "gen.go.tpl")
		{
			path = strings.TrimSuffix(path, TemplateSuffix)
			index := strings.LastIndex(path, string(os.PathSeparator))
			if index != -1 {
				filename := strings.TrimSuffix(path[index+1:], TemplateSuffix)
				path = filepath.Join(plugin, "gen", filename)
			} else {
				return errors.New("模版文件不存在文件路径错误!")
			}
		} // 处理gen.go文件路径
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o666)
		if err != nil {
			return errors.Wrapf(err, "打开文件[filepath:%s]失败!", path)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				zap.L().Error("关闭文件失败!", zap.String("path", path), zap.Error(err))
			}
		}()
		err = parse.Execute(file, info)
		if err != nil {
			return errors.Wrap(err, "模版文件[gen.go]填充失败!")
		}
	} // 创建plugin/gen/main.go
	{
		parse, err := template.New("").Parse(string(internal.Gen))
		if err != nil {
			return errors.Wrap(err, "解析模版文件[main.go]失败!")
		}
		var file *os.File
		path := filepath.Join(global.Config.Server.PluginRoot(), "coder", "service", "internal", "main.go.tpl")
		{
			path = strings.TrimSuffix(path, TemplateSuffix)
			index := strings.LastIndex(path, string(os.PathSeparator))
			if index != -1 {
				filename := strings.TrimSuffix(path[index+1:], TemplateSuffix)
				path = filepath.Join(plugin, filename)
			} else {
				return errors.New("模版文件不存在文件路径错误!")
			}
		} // 处理main.go文件路径
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o666)
		if err != nil {
			return errors.Wrapf(err, "打开文件[filepath:%s]失败!", path)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				zap.L().Error("关闭文件失败!", zap.String("path", path), zap.Error(err))
			}
		}()
		err = parse.Execute(file, info)
		if err != nil {
			return errors.Wrap(err, "模版文件[main.go]填充失败!")
		}
	} // 创建plugin/main.go
	return nil
}
