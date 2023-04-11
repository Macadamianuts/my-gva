package oss

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/plugin/oss/global"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"sync"
)

var Local = new(local)

type local struct{}

// Upload 上传文件
func (l *local) Upload(ctx context.Context, name string, reader io.Reader) (filepath string, filename string, err error) {
	var one sync.Once
	one.Do(func() {
		fmt.Println(global.Config)
		_, err = os.Stat(global.Config.LocalStorage.Path)
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(global.Config.LocalStorage.Path, os.ModePerm)
			if err != nil {
				zap.L().Error("[oss][local][dir:%s]创建文件夹失败!", zap.String("dir", global.Config.LocalStorage.Path))
			}
		}
	})

	filename = global.Config.LocalStorage.Filename(name)
	key := global.Config.LocalStorage.FileKey(filename)
	filepath = global.Config.LocalStorage.Filepath(key)
	var out *os.File
	out, err = os.Create(key)
	if err != nil {
		return filepath, filename, errors.Wrapf(err, "创建文件失败!")
	}
	defer func() { // 创建文件流 defer 关闭
		_ = out.Close()
	}()

	_, err = io.Copy(out, reader) // 传输(拷贝)文件
	if err != nil {
		return filepath, filename, errors.Wrapf(err, "传输(拷贝)文件失败!")
	}
	return filepath, filename, nil
}

// DeleteFile 删除文件
func (l *local) DeleteFile(ctx context.Context, filename string) error {
	key := global.Config.LocalStorage.FileKey(filename)
	err := os.Remove(key)
	if err != nil {
		return errors.Wrapf(err, "[oss][local][key:%s]删除文件失败!", key)
	}
	return nil
}

// UploadByFile .
func (l *local) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	defer func() { // 接收文件流 defer 关闭
		_ = file.Close()
	}()
	return l.Upload(ctx, file.Name(), file)
}

// UploadByHeader .
func (l *local) UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error) {
	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][local]从 header 打开文件流失败!")
	}
	defer func() { // 接收文件流 defer 关闭
		_ = file.Close()
	}()

	return l.Upload(ctx, header.Filename, file)
}
