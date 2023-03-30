package oss

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"gva-lbx/plugin/oss/global"
	"gva-lbx/plugin/oss/internal"
	"io"
	"mime/multipart"
	"os"
)

var _ Oss = (*aliyun)(nil)

var Aliyun = new(aliyun)

type aliyun struct{}

func (a *aliyun) Upload(ctx context.Context, name string, reader io.Reader) (filepath string, filename string, err error) {
	var bucket *oss.Bucket
	bucket, err = internal.Aliyun.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	filename = global.Config.AliyunOss.Filename(name)
	key := global.Config.AliyunOss.FileKey(filename)
	filepath = global.Config.AliyunOss.Filepath(key)
	err = bucket.PutObject(key, reader)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][aliyun]上传文件失败!")
	}
	return
}

func (a *aliyun) DeleteFile(ctx context.Context, filename string) error {
	bucket, err := internal.Aliyun.NewClient()
	if err != nil {
		return err
	}
	key := global.Config.AliyunOss.FileKey(filename)
	err = bucket.DeleteObject(key)
	if err != nil {
		return errors.Wrapf(err, "[oss][aliyun][key:%s]删除文件失败!", key)
	}
	return nil
}

func (a *aliyun) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	defer func() { // 文件 defer 关闭
		_ = file.Close()
	}()
	return a.Upload(ctx, file.Name(), file)
}

func (a *aliyun) UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error) {
	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][aliyun]从 header 打开文件流失败!")
	}
	defer func() { // 文件 defer 关闭
		_ = file.Close()
	}()
	return a.Upload(ctx, header.Filename, file)
}
