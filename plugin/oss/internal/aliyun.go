package internal

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"gva-lbx/plugin/oss/global"
)

var Aliyun = new(aliyun)

type aliyun struct{}

// NewClient 初始化 aliyun 得到 *oss.Bucket
func (a *aliyun) NewClient() (*oss.Bucket, error) {
	client, err := oss.New(global.Config.AliyunOss.Endpoint, global.Config.AliyunOss.AccessKeyId, global.Config.AliyunOss.AccessKeySecret) // 创建 oss.Client 实例。
	if err != nil {
		return nil, errors.Wrap(err, "[oss][aliyun oss]创建client实例失败!")
	}
	var bucket *oss.Bucket
	bucket, err = client.Bucket(global.Config.AliyunOss.Bucket) // 获取存储空间
	if err != nil {
		return nil, errors.Wrap(err, "[oss][aliyun oss]获取存储空间失败!")
	}
	return bucket, nil
}
