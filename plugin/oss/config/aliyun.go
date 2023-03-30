package config

import (
	"fmt"
	"net/url"
	"path"
	"time"
)

// Filepath 访问全路径
func (x *AliyunOss) Filepath(key string) string {
	link, _ := url.JoinPath(x.Domain, key)
	return link
}

// Filename 格式化文件名
func (x *AliyunOss) Filename(filename string) string {
	if x.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), filename)
	}
	return fmt.Sprintf("%s%d_%s", x.Prefix, time.Now().Local().Unix(), filename)
}

// FileKey 文件
func (x *AliyunOss) FileKey(filename string) string {
	return path.Join(x.Path, filename)
}
