package server

import (
	"github.com/bersennaidoo/muzeekstore/backend/controllers"
	"github.com/bersennaidoo/muzeekstore/backend/repositories"
	"gorm.io/gorm"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	productsController *controllers.ProductsController
}

func InitHttpServer(config *viper.Viper, dbClient *gorm.DB) HttpServer {
	productsRepository := repositories.NewProductsRepository(dbClient)
	productsController := controllers.NewProductsController(productsRepository)

	router := gin.Default()

	router.GET("/products", productsController.ListProduct)
	router.POST("/products", productsController.CreateProduct)
	router.GET("/products/:id", productsController.GetProduct)
	router.PUT("/products/:id", productsController.UpdateProduct)
	router.DELETE("/products/:id", productsController.DeleteProduct)

	return HttpServer{
		config:             config,
		router:             router,
		productsController: productsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
