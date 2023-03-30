package utils

import (
	"github.com/pkg/errors"
	"os"
)

var Directory = new(directory)

type directory struct{}

func (d *directory) PathExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		if stat.IsDir() {
			return true, nil
		}
		return false, errors.Errorf("目录下已存在 [%s] 该文件!", path)
	}
	return false, nil
}
