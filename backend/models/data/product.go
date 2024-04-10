package data

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Image       string
	ImageAlt    string
	Price       float64
	Promotion   float64
	ProductName string
	Description string
	Orders []Order
}
