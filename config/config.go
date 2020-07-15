package config

import (
	"fmt"
	"strings"
	"weshierNext/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 全局配置
type Config struct {
	Path string
}

// Init 初始化配置
func Init(confPath string) (err error) {
	c := &Config{
		Path: confPath,
	}
	if err = c.InitConfig(); err != nil {
		return
	}
	logger.Init()
	c.WatchConfig()
	return
}

// InitConfig 初始化配置文件
func (c *Config) InitConfig() (err error) {
	if c.Path != "" {
		// 如果指定了配置文件 则解析配置文件
		viper.SetConfigFile(c.Path)
	} else {
		// 否则使用默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件的类型为 json
	viper.SetConfigType("json")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 环境变量前缀
	viper.SetEnvPrefix("WESHIER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	return
}

// WatchConfig 监听配置文件变化
func (c *Config) WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Logger.Debug(fmt.Sprintf("Config file change: %s\n", e.Name))
	})
}
