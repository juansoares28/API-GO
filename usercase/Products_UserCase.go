package usercase

import (
	"API/models"
	"API/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repo}

}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}
func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	productId, erro := pu.repository.CreateProduct(product)
	if erro != nil {
		return models.Product{}, erro
	}
	product.ID = productId
	return product, nil
}
