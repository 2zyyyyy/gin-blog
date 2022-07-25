package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// AppConfig 配置文件数据源结构体
type AppConfig struct {
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

// App 定义全局变量
var App AppConfig

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("配置文件读取失败, err%s:\n", err)
		return
	}
	LoadData(file)
}

func LoadData(file *ini.File) {
	// 获取server参数
	App.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	App.HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	// 获取database参数
	App.Db = file.Section("database").Key("Db").MustString("mysql")
	App.DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	App.DbPort = file.Section("database").Key("DbPort").MustString("3306")
	App.DbUser = file.Section("database").Key("DbUser").MustString("root")
	App.DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	App.DbName = file.Section("database").Key("DbName").MustString("gin_blog")
}
