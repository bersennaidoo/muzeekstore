package repositories

import (
	"gorm.io/gorm"
	//"github.com/bersennaidoo/muzeekstore/backend/models/data"
)

type ProductsRepository struct {
	dbClient *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		dbClient: db,
	}
}
