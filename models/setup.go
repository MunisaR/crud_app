package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/////////SetupDB: initialization of mysql database //////////////////////////////////

func SetupDB() *gorm.DB {
	USER := "root"
	PASSWORD := "1@3$5^7mR"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "crud"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
