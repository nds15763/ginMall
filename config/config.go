package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Port string `toml:"port"` // 服务器监听端口

	// 数据库相关参数
	DbUrl    string `toml:"database_url"`    // 数据库URL
	DbPort   string `toml:"database_port"`   // 数据库端口
	DbName   string `toml:"database_name"`   // 数据库名称
	DbUser   string `toml:"database_user"`   // 数据库用户名
	DbPasswd string `toml:"database_passwd"` // 数据库密码

	// Redis数据库相关参数
	RedisUrl    string `toml:"redis_url"`    // 数据库URL
	RedisPort   string `toml:"redis_port"`   // 数据库端口
	RedisPasswd string `toml:"redis_passwd"` // 数据库密码

	// 日志相关参数
	LogLevel string `toml:"log_level"`
	LogDest  string `toml:"log_dest"`
	LogDir   string `toml:"log_dir"`
}

func NewConfiguration() *Configuration {

	return &Configuration{ // 在此提供默认值(为所有参数提供默认值)
		Port: "8081",

		LogDir:   "./log",
		LogLevel: "debug",
		LogDest:  "file",

		DbUrl:    "localhost",
		DbPort:   "3306",
		DbName:   "sa",
		DbUser:   "root",
		DbPasswd: "Your PassWord",

		RedisUrl:  "127.0.0.1",
		RedisPort: "6379",
	}

}
func (this *Configuration) InitFromFile(path string) { // 用于启动时加载配置文件

	if _, err := toml.DecodeFile(path, this); err != nil {
		panic(fmt.Sprintf("can't decode conf file: [%s]", path))
	}

}
