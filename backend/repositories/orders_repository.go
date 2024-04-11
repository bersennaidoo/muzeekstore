package repositories

import (
	"gorm.io/gorm"
	//"github.com/bersennaidoo/muzeekstore/backend/models/data"
)

type OrdersRepository struct {
	dbClient *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) *OrdersRepository {
	return &OrdersRepository{
		dbClient: db,
	}
}
