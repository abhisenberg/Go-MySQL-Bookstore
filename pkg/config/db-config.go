package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	temp_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = temp_db
}

func GetDB() *gorm.DB {
	return db
}
