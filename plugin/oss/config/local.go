package config

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
)

// Filepath 根据文件路径
func (x *LocalStorage) Filepath(key string) string {
	link, _ := url.JoinPath(x.Domain, key)
	return link
}

// Filename 文件名
func (x *LocalStorage) Filename(filename string) string {
	ext := path.Ext(filename)
	return fmt.Sprintf("%s_%d%s", strings.TrimSuffix(filename, ext), time.Now().Local().Unix(), ext)
}

// FileKey 文件key
func (x *LocalStorage) FileKey(filename string) string {
	return path.Join(x.Path, filename)
}
