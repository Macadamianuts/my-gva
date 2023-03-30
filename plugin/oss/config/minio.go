package config

import (
	"fmt"
	"gva-lbx/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const DataFormat = "2006-01-02"

func (x *Minio) ExpirationTimeDuration() time.Duration {
	parse, err := utils.Duration.Parse(x.ExpirationTime)
	if err != nil {
		return 0
	}
	return parse
}

// Filepath 访问全路径
func (x *Minio) Filepath(key string) string {
	return x.Domain + path.Join(x.Bucket, key)
}

// Filename 格式化文件名
func (x *Minio) Filename(filename string) string {
	names := strings.Split(filename, string(os.PathSeparator))
	filename = names[len(names)-1:][0]
	if x.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), filename)
	}
	return fmt.Sprintf("%s%d_%s", x.Prefix, time.Now().Local().Unix(), filename)
}

// FileKey 文件key
func (x *Minio) FileKey(filename string) string {
	return filepath.Join(x.Path, time.Now().Format(DataFormat), filename)
}
