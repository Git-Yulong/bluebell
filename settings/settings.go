// Package settings
// @Description:
package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//
//  Init
//  @Description: 初始化viper
//  @return err
//
func Init() (err error) {
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig()

	if err != nil {
		fmt.Printf("viper.ReadInConfig() err=%v\n", err)
		return
	}
	// 监控配置文件的变化
	viper.WatchConfig()
	// 配置文件修改则触发
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return

}
