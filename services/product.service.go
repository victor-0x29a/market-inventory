package services

import (
	"errors"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
	"github.com/market-inventory/database"
	"github.com/market-inventory/repositories"
	"gorm.io/gorm"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func (service ProductService) CreateV1(data *dtos.CreateProductDTO) error {
	err := service.Repository.Create(data)

	return err
}

func (service ProductService) FindOneV1(productId int) (*database.Product, error) {
	product, err := service.Repository.FindOne(productId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrProductNotFound
	}

	return product, err
}

func (service ProductService) FindAllV1(pagination dtos.ApiPagination) (dtos.ApiPaginationResponse, error) {
	products, err := service.Repository.FindAll(pagination)

	return products, err
}

func (service ProductService) UpdateV1(productId int, updateData dtos.UpdateProductDTO) error {
	updateErr := service.Repository.Update(productId, updateData)

	if errors.Is(updateErr, gorm.ErrRecordNotFound) {
		return constants.ErrProductNotFound
	}

	return updateErr
}
