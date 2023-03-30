package internal

import (
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gva-lbx/plugin/oss/global"
	"net/http"
	"net/url"
)

type tencent struct{}

var Tencent = new(tencent)

// NewClient 初始化 tencent 得到 *cos.Client
func (t *tencent) NewClient() (*cos.Client, error) {
	_url, err := url.Parse(global.Config.TencentCos.Domain)
	if err != nil {
		return nil, errors.Wrap(err, "[oss][tencent cos]url 拼接失败!")
	}
	baseURL := &cos.BaseURL{BucketURL: _url}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.Config.TencentCos.SecretId,
			SecretKey: global.Config.TencentCos.SecretKey,
		},
	})
	return client, nil
}
