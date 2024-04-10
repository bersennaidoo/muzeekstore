package data

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      string
	FirstName string
	LastName  string
	Email     string
	Password  string
	LoggedIn  bool
	Order []Order
}
