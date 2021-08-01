package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dataSourceName := "root@tcp(localhost:3306)/vip_system?parseTime=True"
	d, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
