package data

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(file string, server string, database string, username string, password string, port int) {
	var num_datapoints int64

	if len(file) == 0 {
		DB = openMySQL(server, database, username, password, port)
	} else {
		DB, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	}

	DB.AutoMigrate(&Employee{})
	DB.Model(&Employee{}).Count(&num_datapoints)

	if num_datapoints == 0 {
		seedDB(3)
	}
}

func openMySQL(server string, database string, username string, password string, port int) *gorm.DB {
	var url string

	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, server, port, database)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}

func seedDB(num_entries int) {
	for i := 0; i < num_entries; i++ {
		DB.Create(&Employee{Age: GenerateRandomAge(70, 20), Name: GenerateRandomName(10), City: GenerateRandomName(6)})
	}
}
