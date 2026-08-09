package controller

import (
	"API/models"
	"API/usercase"
	"net/http"
	"strconv"

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
func (p *productController) GetProductById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := p.productUsecase.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
