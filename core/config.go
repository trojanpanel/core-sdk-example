package core

import (
	"core-sdk-example/module/constant"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var Config = new(AppConfig)

// InitConfig 初始化全局配置文件
func InitConfig() {
	if err := ini.MapTo(Config, constant.ConfigFilePath); err != nil {
		logrus.Errorf("配置文件加载失败 err: %v", err)
		panic(err)
	}
}

type AppConfig struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	Host     string `ini:"host"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type RedisConfig struct {
	Host      string `ini:"host"`
	Port      int    `ini:"port"`
	Password  string `ini:"password"`
	Db        int    `ini:"db"`
	MaxIdle   int    `ini:"max_idle"`
	MaxActive int    `ini:"max_active"`
	Wait      bool   `ini:"wait"`
}
