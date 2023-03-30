package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gva-lbx/plugin/oss/global"
	"gva-lbx/plugin/oss/internal"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"sync"
)

var MinIO = new(_minio)

type _minio struct{}

func (m *_minio) Upload(ctx context.Context, name string, reader io.Reader) (filepath string, filename string, err error) {
	var client *minio.Client
	client, err = internal.Minio.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	var one sync.Once
	one.Do(func() {
		var exists bool
		exists, err = client.BucketExists(ctx, global.Config.Minio.Bucket)
		if !exists && err == nil {
			options := minio.MakeBucketOptions{}
			err = client.MakeBucket(ctx, global.Config.Minio.Bucket, options)
			if err != nil {
				zap.L().Error("[oss][minio][bucket:%s]存储桶创建失败!", zap.String("bucket", global.Config.Minio.Bucket))
			}
		}
	})

	return filepath, filename, nil
}

func (m *_minio) DeleteFile(ctx context.Context, filename string) error {
	client, err := internal.Minio.NewClient()
	if err != nil {
		return err
	}
	options := minio.RemoveObjectOptions{GovernanceBypass: true}
	key := global.Config.Minio.FileKey(filename)
	err = client.RemoveObject(ctx, global.Config.Minio.Bucket, key, options)
	if err != nil {
		return errors.Wrapf(err, "[oss][minio][key:%s]删除文件失败!", key)
	}
	return nil
}

func (m *_minio) UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error) {
	var client *minio.Client
	client, err = internal.Minio.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	var exists bool
	exists, err = client.BucketExists(ctx, global.Config.Minio.Bucket)
	if !exists && err == nil {
		options := minio.MakeBucketOptions{}
		err = client.MakeBucket(ctx, global.Config.Minio.Bucket, options)
		if err != nil {
			return filepath, filename, errors.Wrapf(err, "[oss][minio][bucket:%s]存储桶创建失败!", global.Config.Minio.Bucket)
		}
	}
	filename = global.Config.Minio.Filename(file.Name())
	key := global.Config.Minio.FileKey(filename)
	filepath = global.Config.Minio.Filepath(key)
	info, _ := file.Stat()
	options := minio.PutObjectOptions{}

	_, err = client.PutObject(ctx, global.Config.Minio.Bucket, key, file, info.Size(), options)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	return filepath, filename, nil
}

func (m *_minio) UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error) {
	var client *minio.Client
	client, err = internal.Minio.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	var exists bool
	exists, err = client.BucketExists(ctx, global.Config.Minio.Bucket)
	if !exists && err == nil {
		options := minio.MakeBucketOptions{}
		err = client.MakeBucket(ctx, global.Config.Minio.Bucket, options)
		if err != nil {
			return filepath, filename, errors.Wrapf(err, "[oss][minio][bucket:%s]存储桶创建失败!", global.Config.Minio.Bucket)
		}
	} else {
		return filepath, filename, errors.Wrapf(err, "[oss][minio][bucket:%s]存储桶不存在!", global.Config.Minio.Bucket)
	}
	filename = global.Config.Minio.Filename(header.Filename)
	key := global.Config.Minio.FileKey(filename)
	filepath = global.Config.Minio.Filepath(key)
	options := minio.PutObjectOptions{ContentType: header.Header.Get("content-type")}

	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]通过header获取io.Reader对象失败!")
	}
	_, err = client.PutObject(ctx, global.Config.Minio.Bucket, key, file, header.Size, options)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	return filepath, filename, nil
}

func (m *_minio) PresignedPutObject(ctx context.Context, filename string) (*url.URL, error) {
	client, err := internal.Minio.NewClient()
	if err != nil {
		return nil, err
	}
	key := global.Config.Minio.FileKey(filename)
	object, err := client.PresignedPutObject(ctx, global.Config.Minio.Bucket, key, global.Config.Minio.ExpirationTimeDuration())
	if err != nil {
		return nil, errors.Wrap(err, "[oss][minio]获取临时链接失败!")
	}
	return object, nil
}
