package data

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `json:"sell_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}
