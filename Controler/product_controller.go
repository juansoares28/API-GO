package controller

import (
	"API/models"
	"API/usercase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usercase.ProductUsercase
}

func NewProductController(usecase usercase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}
func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []models.Product{
		{
			ID:    1,
			Name:  "agua",
			Price: 1.0,
		},
	}

	ctx.JSON(http.StatusOK, gin.H{"products": products})
}
