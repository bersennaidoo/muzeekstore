package controllers

import (
	"io"
	"encoding/json"
	"github.com/bersennaidoo/muzeekstore/backend/models/data"
	"github.com/bersennaidoo/muzeekstore/backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsRepository *repositories.ProductsRepository
}

func NewProductsController(prp *repositories.ProductsRepository) *ProductsController {
	return &ProductsController{
		productsRepository: prp,
	}
}

func (pc *ProductsController) ListProduct(ctx *gin.Context) {
	products, _ := pc.productsRepository.List()
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductsController) ListPromo(ctx *gin.Context) {
	promos, _ := pc.productsRepository.Promos()
	ctx.JSON(http.StatusOK, promos)
}


func (pc *ProductsController) CreateProduct(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)

	var product data.Product
	_ = json.Unmarshal(body, &product)

	resp, _ := pc.productsRepository.Create(product)
	ctx.JSON(http.StatusOK, resp)
}

func (pc *ProductsController) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	ctx.JSON(http.StatusOK, "Hello From GetProduct " + productID)
}

func (pc *ProductsController) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From UpdateProduct")
}

func (pc *ProductsController) DeleteProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello From DeleteProduct")
}