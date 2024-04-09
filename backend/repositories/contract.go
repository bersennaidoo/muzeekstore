package repositories

import (
	"github.com/bersennaidoo/muzeekstore/backend/models"
)

type RepoContract interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
