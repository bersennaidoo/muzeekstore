package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ItemId      string  `json:"item_id"`
	ImageId     string  `json:"image_id"`
	Title       string  `json:"Title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	SalePrice   float64 `json:"sale_price"`
}

type ItemsController struct {
}

func NewItemsController() *ItemsController {
	return &ItemsController{}
}

func (ic ItemsController) ListItems(ctx *gin.Context) {

	items := []item{
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "cookie",
			ImageId:     "cookie",
			Title:       "Cookie",
			Price:       1,
			Description: "May contain nuts",
			SalePrice:   0.50,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
		{
			ItemId:      "coffee",
			ImageId:     "coffee",
			Title:       "Coffee",
			Price:       0.99,
			Description: "",
			SalePrice:   0,
		},
	}
	ctx.JSON(http.StatusOK, items)
}

/*func (pc ProductsController) CreateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From CreateProduct")
}

func (pc ProductsController) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	ctx.JSON(http.StatusOK, "Hello From GetProduct "+productID)
}

func (pc ProductsController) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From UpdateProduct")
}

func (pc ProductsController) DeleteProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From DeleteProduct")
}*/
