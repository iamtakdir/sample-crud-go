package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/bookstoredb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
