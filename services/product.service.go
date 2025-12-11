package services

import (
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/repositories"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func (service ProductService) CreateV1(data *dtos.CreateProductDTO) error {
	err := service.Repository.Create(data)

	return err
}
