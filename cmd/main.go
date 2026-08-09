package main

import (
	controller "API/Controler"
	"API/db"
	"API/repository"
	"API/usercase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// repository
	repo := repository.NewProductRepository(dbConnection)
	// user case
	ProductUsecase := usercase.NewProductUsecase(repo)
	// controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.CreateProduct)

	server.GET("/product/:id", ProductController.GetProductById)

	server.Run(":8000")
}
