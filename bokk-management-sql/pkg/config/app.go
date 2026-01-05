package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Connect() {
	db, err := gorm.Open(
		"mysql",
		"root:root@123@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	database = db
}

func GetDB() *gorm.DB {
	return database
}