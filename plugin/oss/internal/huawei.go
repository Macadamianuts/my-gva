package internal

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"

	"gva-lbx/plugin/oss/global"
)

var Huawei = new(huawei)

type huawei struct{}

// NewHuaWeiObsClient 初始化 huawei obs 得到 *obs.ObsClient
func (h *huawei) NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	client, err = obs.New(global.Config.HuaWeiObs.AccessKey, global.Config.HuaWeiObs.SecretKey, global.Config.HuaWeiObs.Endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "[oss][huawei obs]连接 huawei obs 失败!")
	}
	return client, err
}
