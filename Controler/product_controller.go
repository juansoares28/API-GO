package controller

import (
	"API/models"
	"API/usercase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usercase.ProductUsecase
}

func NewProductController(usecase usercase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}
func (p *productController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	isertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, isertedProduct)
}
