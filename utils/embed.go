package utils

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var Embed = new(_embed)

type _embed struct{}

// Restore 把 embed.FS 持久化文件到二进制运行的当前文件夹中
func Restore(files ...embed.FS) {
	for i := 0; i < len(files); i++ {
		Embed.directory(files[i], ".")
	}
}

// directory
func (e *_embed) directory(files embed.FS, dir string) {
	entries, err := files.ReadDir(dir)
	if err != nil {
		fmt.Printf("[embed] restore folder err[%s]\n", err)
		return
	}
	for i := 0; i < len(entries); i++ {
		if !entries[i].IsDir() {
			e.file(files, filepath.Join(dir, entries[i].Name()), entries[i])
			continue
		}

		nextDir := filepath.Join(dir, entries[i].Name())
		err = os.MkdirAll(nextDir, os.ModePerm) // 创建文件夹, 权限为 os.ModePerm 可自行修改
		if err != nil {
			fmt.Printf("[embed] restore mkdir[%s]\n", nextDir)
			return
		}
		e.directory(files, nextDir)
	}
}

// file
func (e *_embed) file(efs embed.FS, path string, entry fs.DirEntry) {
	if entry.IsDir() { // 文件夹
		e.directory(efs, path)
		return
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		var src fs.File
		src, err = efs.Open(path) // 根据path从embed的到文件数据
		if err != nil {
			fmt.Printf("[embed] restore open file failed, err:[%s]\n", err)
			return
		}
		var dst *os.File
		dst, err = os.Create(path) // 创建本地文件的 writer
		if err != nil {
			fmt.Printf("[embed] restore os create file, err:[%s]\n", err)
			_, _ = fmt.Fprintln(os.Stderr, "[embed restore os create file] write err:", err)
			return
		}
		_, err = io.Copy(dst, src) // 把embed的数据复制到本地
		if err != nil {
			fmt.Printf("[embed] restore io copy file, err:[%s]\n", err)
			return
		}
		defer func() { // 关闭文件流
			_ = src.Close()
			_ = dst.Close()
		}()
		return
	}
	fmt.Printf("[embed] file exist, path[%s]\n", path)
}
