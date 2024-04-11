package data

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      string  `json:"name"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	LoggedIn  bool    `json:"loggedin"`
	Order     []Order `json:"orders"`
}
