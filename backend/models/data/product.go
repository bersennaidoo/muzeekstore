package data

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Image       *string  `json:"img"`
	Imagealt    *string  `json:"imgalt"`
	Price       *float64 `json:"price"`
	Promotion   *float64 `json:"promotion"`
	Productname *string  `json:"productname"`
	Description *string  `json:"desc"`
	Orders      []Order  `json:"orders"`
}
