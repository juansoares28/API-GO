package main

import (
	controller "API/Controler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductController := controller.ProductController{}

	server.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
