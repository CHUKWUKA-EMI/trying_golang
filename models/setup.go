package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
  dsn := "host=localhost user=postgres password=start12345 dbname=book_store port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err !=nil{
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Book{})
	DB = db
}
