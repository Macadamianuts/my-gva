package oss

import (
	"context"
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

func (l *local) Upload(ctx context.Context, name string, reader io.Reader) (filepath string, filename string, err error) {
	var one sync.Once
	one.Do(func() {
		_, err = os.Stat(global.Config.LocalStore.Path)
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(global.Config.LocalStore.Path, os.ModePerm)
			if err != nil {
				zap.L().Error("[oss][local][dir:%s]创建文件夹失败!", zap.String("dir", global.Config.LocalStore.Path))
			}
		}
	})

	filename = global.Config.LocalStore.Filename(name)
	key := global.Config.LocalStore.FileKey(filename)
	filepath = global.Config.LocalStore.Filepath(key)
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

func (l *local) DeleteFile(ctx context.Context, filename string) error {
	key := global.Config.LocalStore.FileKey(filename)
	err := os.Remove(key)
	if err != nil {
		return errors.Wrapf(err, "[oss][local][key:%s]删除文件失败!", key)
	}
	return nil
}

func (l *local) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	defer func() { // 接收文件流 defer 关闭
		_ = file.Close()
	}()
	return l.Upload(ctx, file.Name(), file)
}

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
