package repositories

import (
	"gorm.io/gorm"
	"github.com/bersennaidoo/muzeekstore/backend/models/data"
)

type ProductsRepository struct {
	dbClient *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		dbClient: db,
	}
}

func (pr *ProductsRepository) List() (products []data.Product, err error) {
	return products, pr.dbClient.Find(&products).Error
}

func (pr *ProductsRepository) Promos() (products []data.Product, err error) {
	return products, pr.dbClient.Where("promotion IS NOT NULL").Find(&products).Error
}

func (pr *ProductsRepository) Create(product data.Product) (data.Product, error) {
	return  product, pr.dbClient.Create(&product).Error
}

