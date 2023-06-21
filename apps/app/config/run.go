package config

import (
	"audio/common/globalkey"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func GetConfigData() string {
	//导入配置文件
	wd, _ := os.Getwd()
	viper.SetConfigType("yaml")
	logx.Info(wd)

	viper.SetConfigFile(wd + "/etc/userRpc.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	var _config *globalkey.Config
	//将配置文件读到结构体中
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	host := _config.App.Host
	return host
}

func GetConfigSsl() bool {
	//导入配置文件
	wd, _ := os.Getwd()
	viper.SetConfigType("yaml")
	logx.Info(wd)

	viper.SetConfigFile(wd + "/etc/userRpc.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	var _config *globalkey.Config
	//将配置文件读到结构体中
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	ssl := _config.App.Ssl
	return ssl
}
