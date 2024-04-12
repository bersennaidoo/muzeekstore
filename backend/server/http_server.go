package server

import (
	"log"

	"github.com/bersennaidoo/muzeekstore/backend/controllers"
	"github.com/bersennaidoo/muzeekstore/backend/repositories"
	"github.com/gin-contrib/cors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	productsController *controllers.ProductsController
	itemsController    *controllers.ItemsController
}

func InitHttpServer(config *viper.Viper, dbClient *gorm.DB) HttpServer {
	productsRepository := repositories.NewProductsRepository(dbClient)
	productsController := controllers.NewProductsController(productsRepository)
	itemsController := controllers.NewItemsController()

	router := gin.Default()

	/*router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))*/
	router.Use(cors.Default())

	router.GET("/products", productsController.ListProduct)
	router.GET("/promos", productsController.ListPromo)
	router.POST("/products", productsController.CreateProduct)
	router.GET("/products/:id", productsController.GetProduct)
	router.PUT("/products/:id", productsController.UpdateProduct)
	router.DELETE("/products/:id", productsController.DeleteProduct)

	router.GET("/items", itemsController.ListItems)

	return HttpServer{
		config:             config,
		router:             router,
		productsController: productsController,
		itemsController:    itemsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
