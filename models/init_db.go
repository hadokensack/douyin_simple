package models

import (
	"byte_douyin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 通过gorm框架现有数据库连接来初始化
var DB *gorm.DB

func Init_DB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//GORM 的 AutoMigrate() 方法用于初始化多张表
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})

	if err != nil {
		panic(err)
	}
}
