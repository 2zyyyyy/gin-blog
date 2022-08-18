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
	Other      string
}

// Jwt 参数结构体
type Jwt struct {
	SigningKey  string
	Issuer      string
	ExpiresTime int64
	BufferTime  int64
}

// QiNiuYun 七牛云配置结构体
type QiNiuYun struct {
	AccessKey    string
	SecretKey    string
	Bucket       string
	QiuNiuServer string
}

// App 定义全局变量
var App AppConfig
var JWT Jwt
var QiNiu QiNiuYun

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
	App.Other = file.Section("database").Key("Other").MustString("")

	// 获取jwt参数
	JWT.SigningKey = file.Section("jwt").Key("SigningKey").MustString("8jhs7H9n")
	JWT.ExpiresTime = file.Section("jwt").Key("ExpiresTime").MustInt64(604800)
	JWT.BufferTime = file.Section("jwt").Key("BufferTime").MustInt64(86400)
	JWT.Issuer = file.Section("jwt").Key("Issuer").MustString("miles")

	// 获取七牛云配置参数
	QiNiu.AccessKey = file.Section("qiniu").Key("AccessKey").MustString("")
	QiNiu.SecretKey = file.Section("qiniu").Key("SecretKey").MustString("")
	QiNiu.Bucket = file.Section("qiniu").Key("Bucket").MustString("go-blog-img")
	QiNiu.QiuNiuServer = file.Section("qiniu").Key("QiuNiuServer").MustString("")
}
