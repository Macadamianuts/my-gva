package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gva-lbx/plugin/coder/global"
	"path/filepath"
	"strings"
)

var Viper = new(_viper)

type _viper struct{}

func (c *_viper) Initialization() {
	v := viper.New()
	v.AddConfigPath("configs")
	fileType := "yaml"
	name := strings.Join([]string{"coder", gin.Mode()}, ".")
	v.SetConfigType(fileType)
	v.SetConfigName(name)
	filename := strings.Join([]string{name, fileType}, ".")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("[viper][filename:%s][err:%v]配置文件读取失败!\n", filename, err)
		return
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("[viper][filename:%s]配置文件更新\n", in.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Printf("[viper][filename:%s][err:%v]配置文件更新失败!\n", in.Name, err)
			return
		}
	})
	v.WatchConfig()
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Printf("[viper][filename:%s][err:%v]配置文件更新失败!\n", filename, err)
		return
	}
	global.Viper = v
	global.Config.Server.Root, _ = filepath.Abs("")
	global.Config.Web.Root = filepath.Join(global.Config.Server.Root, "web")
}
