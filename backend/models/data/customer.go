package data

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string  `json:"name"`
	FirstName    string  `json:"firstname"`
	LastName     string  `json:"lastname"`
	Email        string  `json:"email" gorm:"unique"`
	Password     string  `json:"password"`
	CCCustomerid string  `json:"cc_customer_id"`
	LoggedIn     bool    `json:"loggedin"`
	Order        []Order `json:"orders"`
}
