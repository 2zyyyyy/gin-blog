package model

import (
	"fmt"
	"gin-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		utils.App.DbUser, utils.App.DbPassWord, utils.App.DbHost, utils.App.DbPort, utils.App.DbName, utils.App.Other)
	//fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Info),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置自动迁移
	err = db.AutoMigrate(&Article{}, &Category{}, &User{})
	if err != nil {
		fmt.Printf("db.AutoMigrate() failed, err%s\n", err)
		return
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("获取通用数据库对象失败, err%s\n", err)
	}
	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置数据库最大连接数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
