package oss

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gva-lbx/plugin/oss/global"
	"gva-lbx/plugin/oss/internal"
	"mime/multipart"
	"os"
	"path"
	"time"
)

var _ Oss = (*tencent)(nil)

var Tencent = new(tencent)

type tencent struct{}

func (t *tencent) Filepath(key string) string {
	return path.Join(global.Config.TencentCos.Domain, key)
}

func (t *tencent) Filename(filename string) string {
	if global.Config.TencentCos.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), filename)
	}
	return fmt.Sprintf("%s%d_%s", global.Config.TencentCos.Prefix, time.Now().Local().Unix(), filename)
}

func (t *tencent) FileKey(filename string) string {
	return path.Join(global.Config.TencentCos.Path, filename)
}

func (t *tencent) DeleteFile(ctx context.Context, filename string) error {
	client, err := internal.Tencent.NewClient()
	if err != nil {
		return err
	}
	key := t.FileKey(filename)
	_, err = client.Object.Delete(ctx, key)
	if err != nil {
		return errors.Wrapf(err, "[oss][tencent cos][key:%s]删除失败!", key)
	}
	return nil
}

func (t *tencent) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	var info os.FileInfo
	info, err = file.Stat()
	defer func() { // 文件流 defer 关闭
		_ = file.Close()
	}()

	var client *cos.Client
	client, err = internal.Tencent.NewClient()
	if err != nil {
		return filepath, filename, err
	}

	filename = t.Filename(info.Name())
	key := t.FileKey(filename)
	filepath = t.Filepath(key)
	_, err = client.Object.Put(ctx, key, file, nil)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][tencent cos]上传文件失败!")
	}
	return filepath, filename, nil
}

func (t *tencent) UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error) {
	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][tencent cos]从 header 打开文件流失败!")
	}
	defer func() { // 文件流 defer 关闭
		_ = file.Close()
	}()

	var client *cos.Client
	client, err = internal.Tencent.NewClient()
	if err != nil {
		return filepath, filename, err
	}

	filename = t.Filename(header.Filename)
	key := t.FileKey(filename)
	filepath = t.Filepath(key)
	_, err = client.Object.Put(ctx, key, file, nil)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][tencent cos]上传文件失败!")
	}

	return filepath, filename, nil
}
