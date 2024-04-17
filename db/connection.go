package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//postgres

var DSN = "host=localhost user=postgres password=123456 dbname=postgres port=5432 dbname=gorm"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Connection Opened to Database")
	}
}
