package repositories

import (
	"gorm.io/gorm"
	//"github.com/bersennaidoo/muzeekstore/backend/models/data"
)

type CustomersRepository struct {
	dbClient *gorm.DB
}

func NewCustomersRepository(db *gorm.DB) *CustomersRepository {
	return &CustomersRepository{
		dbClient: db,
	}
}
