package internal

import (
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"gva-lbx/plugin/oss/global"
)

var Qiniu = new(qiniu)

type qiniu struct{}

// NewFormUploader 根据配置文件获取 storage.FormUploader
func (q *qiniu) NewFormUploader() (*storage.FormUploader, error) {
	config, err := q.config()
	if err != nil {
		return nil, err
	}
	formUploader := storage.NewFormUploader(config)
	return formUploader, nil
}

// Token 根据配置文件获取 上传token
func (q *qiniu) Token() string {
	credentials := auth.New(global.Config.QiniuKodo.AccessKey, global.Config.QiniuKodo.SecretKey)
	putPolicy := storage.PutPolicy{Scope: global.Config.QiniuKodo.Bucket}
	return putPolicy.UploadToken(credentials)
}

// Config 初始化 qiniu 配置得到 *storage.Config
func (q *qiniu) config() (*storage.Config, error) {
	config := &storage.Config{
		UseHTTPS:      global.Config.QiniuKodo.UseHttps,
		UseCdnDomains: global.Config.QiniuKodo.UseCdnDomains,
	}

	region, err := storage.GetRegion(global.Config.QiniuKodo.AccessKey, global.Config.QiniuKodo.Bucket) // 用来根据ak和bucket来获取空间相关的机房信息
	if err != nil {
		zap.L().Error("根据ak和bucket来获取空间相关的机房信息失败!", zap.String("AccessKey", global.Config.QiniuKodo.AccessKey), zap.String("Bucket", global.Config.QiniuKodo.Bucket))
		return nil, errors.Wrapf(err, "[oss][qiniu]根据ak和bucket来获取空间相关的机房信息失败!")
	}
	config.Region = region
	return config, nil
}

// BucketManager 获取 qiniu 的 *storage.BucketManager
func (q *qiniu) BucketManager() (*storage.BucketManager, error) {
	credentials := auth.New(global.Config.QiniuKodo.AccessKey, global.Config.QiniuKodo.SecretKey)
	config, err := q.config()
	if err != nil {
		return nil, err
	}
	bucketManager := storage.NewBucketManager(credentials, config)
	return bucketManager, nil
}
