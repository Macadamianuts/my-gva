package internal

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"gva-lbx/plugin/oss/global"
)

var Minio = new(_minio)

type _minio struct{}

// NewClient 初始化 minio 得到 *minio.Client
func (m *_minio) NewClient() (*minio.Client, error) {
	client, err := minio.New(global.Config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.Config.Minio.AccessKey, global.Config.Minio.SecretKey, global.Config.Minio.Token),
		Secure: global.Config.Minio.UseSsl,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[oss][minio]获取minio.Client对象失败!")
	}
	return client, nil
}
