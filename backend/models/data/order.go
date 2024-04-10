package data

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID   uint
	ProductID    uint
	Price        float64
	PurchaseDate time.Time
}
