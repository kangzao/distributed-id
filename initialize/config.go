package initialize

import (
	"distributed-id/global"
	"fmt"
	"github.com/spf13/viper"
)

// GetEnvInfo 区分dev test prod
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv() //自动读取环境变量
	return viper.GetBool(env)
}

// InitConfig 初始化配置信息，从配置文件中读取配置
func InitConfig() {
	debug := GetEnvInfo("DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-debug.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
