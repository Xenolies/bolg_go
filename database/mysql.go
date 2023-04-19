package database

/**
* @Author: Xenolies
* @Date: 2023/4/16 12:10
* @Version: 1.0
 */

import (
	"bolg_go/config"
	"bolg_go/model"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" //导入包但不使用，init()
	"gorm.io/gorm"
	"log"
)

var DefaultDb *gorm.DB

func InitMySQL(admin config.Admin) *gorm.DB {
	dsn := admin.DbLoginID + ":" + admin.DbLoginPwd + "@(" + admin.DbAddr + ")/" + admin.DbName + "?" + admin.DbConfig
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("数据库启动异常: %s", err)
	}

	DefaultDb = db
	// 读取配置文件中的Admin 创建表
	fmt.Println("创建表")
	DefaultDb.AutoMigrate(&model.Articles{}, &model.Users{})
	return db
}
