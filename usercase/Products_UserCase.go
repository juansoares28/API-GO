package usercase

import "API/models"

type ProductUsecase struct{}

func NewProductUsecase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return []models.Product{}, nil
}
