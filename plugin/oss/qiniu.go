package oss

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/storage"
	"gva-lbx/global"
	"gva-lbx/plugin/oss/internal"
	"mime/multipart"
	"os"
	"path"
	"time"
)

var _ Oss = (*qiniu)(nil)

var Qiniu = new(qiniu)

type qiniu struct{}

func (q *qiniu) Filepath(key string) string {
	return path.Join(global.Config.QiniuKodo.Domain, key)
}

func (q *qiniu) Filename(filename string) string {
	if global.Config.TencentCos.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), filename)
	}
	return fmt.Sprintf("%s%d_%s", global.Config.QiniuKodo.Prefix, time.Now().Local().Unix(), filename)
}

func (q *qiniu) FileKey(filename string) string {
	return path.Join(global.Config.QiniuKodo.Path, filename)
}

func (q *qiniu) DeleteFile(ctx context.Context, filename string) error {
	manager, err := internal.Qiniu.BucketManager()
	if err != nil {
		return err
	}
	key := q.FileKey(filename)
	err = manager.Delete(global.Config.QiniuKodo.Bucket, key)
	if err != nil {
		return errors.Wrapf(err, "[oss][qiniu kodo][key:%s]删除文件失败!", key)
	}
	return nil
}

func (q *qiniu) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	token := internal.Qiniu.Token()
	var uploader *storage.FormUploader
	uploader, err = internal.Qiniu.NewFormUploader()
	if err != nil {
		return filepath, filename, err
	}

	var info os.FileInfo
	info, err = file.Stat()
	filename = q.Filename(file.Name())
	key := q.FileKey(filename)
	filepath = q.Filepath(key)

	ret := storage.PutRet{}
	err = uploader.Put(ctx, &ret, token, key, file, info.Size(), nil)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][qiniu kodo]上传文件失败!")
	}
	return filepath, filename, nil
}

func (q *qiniu) UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error) {
	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][qiniu kodo]从 header 打开文件流失败!")
	}
	defer func() { // 文件流 defer 关闭
		_ = file.Close()
	}()

	token := internal.Qiniu.Token()
	var uploader *storage.FormUploader
	uploader, err = internal.Qiniu.NewFormUploader()
	if err != nil {
		return filepath, filename, err
	}

	filename = q.Filename(header.Filename)
	key := q.FileKey(filename)
	filepath = q.Filepath(key)

	ret := storage.PutRet{}
	err = uploader.Put(ctx, &ret, token, key, file, header.Size, nil)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][qiniu kodo]上传文件失败!")
	}
	return filepath, filename, nil
}
