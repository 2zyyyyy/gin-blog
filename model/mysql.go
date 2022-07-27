package model

import (
	"fmt"
	"gin-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		utils.App.DbUser, utils.App.DbPassWord, utils.App.DbHost, utils.App.DbPort, utils.App.DbName, utils.App.Other)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置自动迁移
	err = db.AutoMigrate()
	if err != nil {
		fmt.Printf("db.AutoMigrate() failed, err%s\n", err)
		return
	}
	fmt.Println("mysql init success")
}
