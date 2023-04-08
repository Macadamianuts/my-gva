package service

import (
	"archive/zip"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/plugin/coder/global"
	"gva-lbx/plugin/coder/model/request"
	"gva-lbx/plugin/coder/model/response"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var AutoCode = new(autoCode)

type autoCode struct{}

// Create 创建
func (s *autoCode) Create(ctx context.Context, info request.AutoCodeCreate) error {
	info.Pretreatment()
	if info.AutoMoveFile && AutoCodeHistory.Repeat(ctx, info.AutoCodeHistoryRepeat()) {
		return errors.New("重复创建!")
	} // 查询代码生成器历史 是否重复创建结构体
	templates, files, dirs, err := s.GetNeedList(info)
	if err != nil {
		return err
	}
	for i := 0; i < len(dirs); i++ {
		err = os.MkdirAll(dirs[i], os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "批量创建文件夹失败!")
		}
	} // 批量创建文件夹

	for i := 0; i < len(templates); i++ {
		var file *os.File
		file, err = os.OpenFile(templates[i].AutoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		err = templates[i].Template.Execute(file, info)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return errors.Wrap(err, "关闭文件流失败!")
		}
	} // 生成文件

	defer func() {
		err = os.RemoveAll(AutoCodePath)
		if err != nil {
			zap.L().Error("删除生成的文件失败!", zap.String("path", AutoCodePath), zap.Error(err))
		}
	}() // 移除中间文件

	if info.AutoMoveFile { // 迁移文件
		length := len(templates)
		for i := 0; i < len(templates); i++ {
			var fileInfo os.FileInfo
			fileInfo, err = os.Lstat(templates[i].MoveFilePath)
			if err == nil {
				if !fileInfo.IsDir() {
					return errors.Wrapf(err, "目标文件[%s]已存在", templates[i].MoveFilePath)
				}
			}
			if !os.IsNotExist(err) {
				return errors.Wrapf(err, "目标文件[%s]已存在", templates[i].MoveFilePath)
			}
		} // 判断目标文件是否都可以移动
		for i := 0; i < length; i++ {
			err = os.Rename(templates[i].AutoCodePath, templates[i].MoveFilePath)
			if err != nil {
				return errors.Wrapf(err, "文件从[%s]移动到[%s]", templates[i].AutoCodePath, templates[i].MoveFilePath)
			}
		} // 移动文件
		pwd := global.Config.Server.PluginPath(info.Plugin)
		{
			//{
			//	err = ast.GenApplyBasic(filepath.Join(pwd, "gen", "main.go"), info.Plugin, info.Struct)
			//	if err != nil {
			//		return err
			//	}
			//} // ast 注入plugin/gen/main.go表结构
			//{
			//	err = ast.RouterRegister(filepath.Join(pwd, "main.go"), info.Plugin, info.Struct)
			//	if err != nil {
			//		return err
			//	}
			//} // ast 注入plugin/main.go表结构
			//{
			//	err = ast.GormAutoMigrate(filepath.Join(pwd, "main.go"), info.Plugin, info.Struct)
			//	if err != nil {
			//		return err
			//	}
			//} // ast 注入plugin/main.go路由
		} // ast 注入
		{
			var output []byte
			output, err = exec.Command("go", "run", filepath.Join(pwd, "gen", "main.go")).CombinedOutput()
			if err != nil {
				return errors.Wrap(err, "gen生成文件失败!")
			}
			zap.L().Debug(string(output))
		} // 执行插件的gen/main.go
		err = AutoCodeHistory.Create(ctx, info)
		if err != nil {
			return err
		} // 保存代码生成器的信息

	} else { // 生成压缩包
		err = s.Zip("plugin.zip", files, ".", ".")
		if err != nil {
			return err
		}
	}
	return nil
}

// Preview 预览
func (s *autoCode) Preview(ctx context.Context, info request.AutoCodeCreate) (data map[string]string, err error) {
	return nil, nil
}

// GetNeedList .
func (s *autoCode) GetNeedList(info request.AutoCodeCreate) (data []response.Template, files []string, dirs []string, err error) {
	templatePath := global.Config.Server.TemplatePath()
	pwd := filepath.Join(templatePath, info.Type)
	templates, names, err := s.GetTemplateFile(pwd, nil, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	length := len(templates)
	data = make([]response.Template, 0, length)
	files = make([]string, 0, length)
	dirs = make([]string, 0, length)
	for i := 0; i < length; i++ {
		var parse *template.Template
		parse, err = template.New(names[i]).Funcs(info.Functions()).ParseFiles(templates[i])
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "解析模版文件失败!")
		}
		entity := response.Template{
			Template:     parse,
			TemplatePath: templates[i],
		}
		{
			prefix := pwd + string(os.PathSeparator)
			base := strings.TrimPrefix(templates[i], prefix)
			if base == "readme.txt.tpl" {
				entity.AutoCodePath = filepath.Join(pwd, AutoCodePath, "readme.txt")
				continue
			}
			index := strings.LastIndex(base, string(os.PathSeparator))
			if index != -1 {
				filename := strings.TrimSuffix(base[index+1:], TemplateSuffix)
				idx := strings.Index(filename, ".")
				if idx != -1 {
					if filename[idx:] != ".go" { // web
						switch filename[:idx] {
						case "api":
							entity.MoveFilePath = filepath.Join(global.Config.Web.ApiPath(info.Plugin), info.Filename+filename[idx:])
						case "form":
							entity.MoveFilePath = filepath.Join(global.Config.Web.FormPath(info.Plugin), info.Filename+filename[idx:])
						case "view":
							entity.MoveFilePath = filepath.Join(global.Config.Web.ViewPath(info.Plugin), info.Filename+filename[idx:])
						}
						filename = filepath.Join(filename[:idx], info.Filename+filename[idx:])
					} else { // server
						switch filename[:idx] {
						case "api":
							entity.MoveFilePath = filepath.Join(global.Config.Server.ApiPath(info.Plugin), info.UnderlineName+filename[idx:])
						case "model":
							entity.MoveFilePath = filepath.Join(global.Config.Server.ModelPath(info.Plugin), info.UnderlineName+filename[idx:])
						case "router":
							entity.MoveFilePath = filepath.Join(global.Config.Server.RouterPath(info.Plugin), info.UnderlineName+filename[idx:])
						case "service":
							entity.MoveFilePath = filepath.Join(global.Config.Server.ServicePath(info.Plugin), info.UnderlineName+filename[idx:])
						case "request":
							entity.MoveFilePath = filepath.Join(global.Config.Server.RequestPath(info.Plugin), info.UnderlineName+filename[idx:])
						}
						filename = filepath.Join(filename[:idx], info.UnderlineName+filename[idx:])
					}
					entity.AutoCodePath = filepath.Join(pwd, AutoCodePath, base[:index], filename)
				}
			}
		} // 处理生成文件的路径
		{
			dir := strings.TrimPrefix(entity.AutoCodePath, pwd+string(os.PathSeparator))
			index := strings.Index(dir, string(os.PathSeparator))
			if index != -1 {
				start := dir[:index+1]
				end := strings.TrimPrefix(dir[index+1:], info.Type+string(os.PathSeparator))
				idx := strings.LastIndex(end, string(os.PathSeparator))
				dirs = append(dirs, start+end[:idx])
			}
		} // dirs 数据组装
		{
			files = append(files, entity.AutoCodePath)
		} // files 数据组装
		data = append(data, entity)
	}
	return data, files, dirs, nil
}

// GetTemplateFile 获取模版文件的全路径
func (s *autoCode) GetTemplateFile(path string, fileList []string, names []string) ([]string, []string, error) {
	if cap(fileList) == 0 {
		fileList = make([]string, 0, 3)
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, errors.Wrap(err, "打开模版文件夹失败!")
	}
	length := len(files)
	for i := 0; i < length; i++ {
		if files[i].IsDir() { // 如果是文件夹的话递归调用获取文件的路径
			fileList, names, err = s.GetTemplateFile(filepath.Join(path, files[i].Name()), fileList, names)
			if err != nil {
				return nil, nil, err
			}
		}
		if strings.HasSuffix(files[i].Name(), TemplateSuffix) {
			names = append(names, files[i].Name())
			fileList = append(fileList, filepath.Join(path, files[i].Name()))
		}
	}
	return fileList, names, nil
}

// Zip 压缩文件
func (s *autoCode) Zip(filename string, files []string, oldForm, newForm string) error {
	newZip, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, "创建zip文件失败!")
	}
	zipWriter := zip.NewWriter(newZip)
	defer func() {
		err = newZip.Close()
		if err != nil {
			zap.L().Error("关闭zip文件流失败!", zap.Error(err))
		}
		err = zipWriter.Close()
		if err != nil {
			zap.L().Error("关闭zip文件写入流失败!", zap.Error(err))
		}
	}()

	for i := 0; i < len(files); i++ {
		err = func(filepath string) error {
			var file *os.File
			file, err = os.Open(filepath)
			if err != nil {
				return errors.Wrap(err, "打开文件失败！")
			}
			defer func() {
				err = file.Close()
				if err != nil {
					zap.L().Error("关闭文件流失败!", zap.Error(err))
				}
			}()
			var info os.FileInfo
			info, err = file.Stat()
			if err != nil {
				return errors.Wrap(err, "获取文件的详细信息失败!")
			}
			var header *zip.FileHeader
			header, err = zip.FileInfoHeader(info)
			if err != nil {
				return errors.Wrap(err, "获取文件的header信息失败!")
			}
			// 使用上面的FileInfoHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
			header.Name = strings.Replace(filepath, oldForm, newForm, -1)
			// 优化压缩 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
			header.Method = zip.Deflate
			var writer io.Writer
			writer, err = zipWriter.CreateHeader(header)
			if err != nil {
				return errors.Wrap(err, "创建文件的新header失败!")
			}
			_, err = io.Copy(writer, file)
			if err != nil {
				return errors.Wrap(err, "复制文件失败!")
			}
			return nil
		}(files[i])
		if err != nil {
			return err
		}
	} // 把files添加到zip中
	return nil
}
