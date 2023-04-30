package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Setup 启动数据库
func Setup() {
	var err error
	dsn := "root:root1234@(127.0.0.1:13306)/db3?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal("models/modles.go Setup() error=", err.Error())
	}
	DB.AutoMigrate(&User{}, &Record{})
}
