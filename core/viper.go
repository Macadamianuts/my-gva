package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gva-lbx/core/internal"
	"gva-lbx/global"
	"path/filepath"
	"strings"
)

var Viper = new(_viper)

type _viper struct{}

// Initialization 初始化
func (c *_viper) Initialization() {
	entries, err := internal.Viper.GetFiles()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	v := viper.New()
	v.AddConfigPath(internal.ConfigPath)
	length := len(entries)
	for i := 0; i < length; i++ {
		filename := entries[i].Name()
		names := strings.Split(filename, ".")
		if len(names) == 3 {
			config := names[0]
			configMode := names[1]
			configType := names[2]
			if configMode != gin.Mode() {
				continue
			}
			v.SetConfigName(strings.Join([]string{config, configMode}, "."))
			v.SetConfigType(configType)
			err = v.ReadInConfig()
			if err != nil {
				fmt.Printf("[viper][filename:%s][err:%v]配置文件读取失败!\n", names, err)
				return
			}
			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Printf("[viper][filename:%s]配置文件更新\n", e.Name)
				if err = v.Unmarshal(&global.Config); err != nil {
					fmt.Printf("[viper][filename:%s][err:%v]配置文件更新失败!\n", e.Name, err)
				}
				filename = filepath.Base(e.Name)
				switch filename {
				case "zap" + gin.Mode() + configType:
					Zap.Initialization()
				case "gorm" + gin.Mode() + configType:
					Gorm.Initialization()
				}
			})
			v.WatchConfig()
			if err = v.Unmarshal(&global.Config); err != nil {
				fmt.Printf("[viper][err:%v]序列化失败!\n", err)
				continue
			}
		}
	}
	global.Viper = v
}

// PluginConfig 插件配置初始化
func (c *_viper) PluginConfig(config any, name string) {
	filename := name + "." + gin.Mode() + ".yaml"
	bytes := internal.Viper.GetFile(filename)
	err := global.Viper.MergeConfig(bytes)
	if err != nil {
		fmt.Printf("[viper][filename:%s][err:%v]配置文件合并失败!\n", filename, err)
		return
	}
	global.Viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("[viper][filename:%s]配置文件更新\n", in.Name)
		if err = global.Viper.Unmarshal(config); err != nil {
			fmt.Printf("[viper][filename:%s][err:%v]配置文件更新失败!\n", in.Name, err)
		}
	})
	global.Viper.WatchConfig()
	if err = global.Viper.Unmarshal(config); err != nil {
		fmt.Printf("[viper][err:%v]序列化失败!\n", err)
	}
}
