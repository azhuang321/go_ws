package initialize

import (
	"chat_api/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

// InitConfig 初始化框架配置
func InitConfig() {
	runMod := GetEnvInfo("pro")
	var configFileName string
	configFileNamePrefix := "config"

	if runMod == "pro" {
		runMod = "pro"
		configFileName = fmt.Sprintf("%s_pro.yaml", configFileNamePrefix)
	} else {
		runMod = "debug"
		configFileName = fmt.Sprintf("%s_debug.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置信息失败：%s\n", err.Error())
		return
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Printf("赋值配置信息失败：%s\n", err.Error())
		return
	}
	global.Config.RunMod = runMod
	fmt.Printf("配置信息：%+v\n", global.Config)

	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Printf("配置信息改变：%s\n", e.Name)
			_ = v.ReadInConfig() // 读取配置数据
			_ = v.Unmarshal(&global.Config)
			global.Config.RunMod = runMod
			fmt.Printf("配置信息：%+v\n", global.Config)
		})
	}()
}
