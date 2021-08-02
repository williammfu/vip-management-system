package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type DBInfo struct {
	Username string
	Password string
	Host     string
	Port     string
}

func Connect() {
	dbInfo := DBInfo{
		Username: "root",
		Password: "",
		Host:     "localhost",
		Port:     "3306"}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/vip_system?parseTime=True",
		dbInfo.Username, dbInfo.Password, dbInfo.Host, dbInfo.Port)
	d, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
