package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bersennaidoo/muzeekstore/backend/models/data"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&data.Customer{}, &data.Product{}, &data.Order{})
}