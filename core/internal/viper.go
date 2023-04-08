package internal

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

var Viper = new(_viper)

type _viper struct{}

// GetFiles 获取配置文件夹信息
func (i *_viper) GetFiles() ([]os.DirEntry, error) {
	entries, err := os.ReadDir(ConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "[viper][path:%s]获取配置文件夹信息失败!", ConfigPath)
	}
	return entries, nil
}

// GetFile 获取文件信息
func (i *_viper) GetFile(filename string) io.Reader {
	file, err := os.Open(filepath.Join(ConfigPath, filename))
	if err != nil {
		fmt.Printf("[viper][filename:%s]文件不存在\n", filename)
		return nil
	}
	return file
}

// GetFileBytes 根据文件名获取文件内容
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (i *_viper) GetFileBytes(filename string) []byte {
	open, err := os.ReadFile(filepath.Join(ConfigPath, filename))
	if err != nil {
		fmt.Printf("[viper][filename:%s]文件不存在\n", filename)
		return nil
	}
	return open
}

// GetFileBytesByPath 根据文件路径获取文件内容
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (i *_viper) GetFileBytesByPath(path string) []byte {
	open, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[viper][filename:%s]文件不存在\n", path)
		return nil
	}
	return open
}
