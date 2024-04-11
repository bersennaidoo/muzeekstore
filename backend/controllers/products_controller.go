package controllers

import (
	"github.com/bersennaidoo/muzeekstore/backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsRepository *repositories.ProductsRepository
}

func NewProductsController(pDB *repositories.ProductsRepository) *ProductsController {
	return &ProductsController{
		productsRepository: pDB,
	}
}

func (pc ProductsController) ListProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From ListProduct")
}

func (pc ProductsController) CreateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From CreateProduct")
}

func (pc ProductsController) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	ctx.JSON(http.StatusOK, "Hello From GetProduct " + productID)
}

func (pc ProductsController) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From UpdateProduct")
}

func (pc ProductsController) DeleteProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From DeleteProduct")
}