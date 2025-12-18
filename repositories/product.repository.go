package repositories

import (
	"context"
	"math"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/database"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Database *gorm.DB
}

func (repository ProductRepository) Create(data *dtos.CreateProductDTO) error {
	ctx := context.Background()

	err := gorm.G[database.Product](repository.Database).Create(ctx, &database.Product{
		Title:             data.Title,
		Description:       data.Description,
		Price:             data.Price,
		InventoryQuantity: data.InventoryQuantity,
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository ProductRepository) FindOne(productId int) (*database.Product, error) {
	ctx := context.Background()

	product, err := gorm.G[database.Product](repository.Database).Where("id = ?", productId).First(ctx)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repository ProductRepository) FindAll(pagination dtos.ApiPagination) (dtos.ApiPaginationResponse, error) {

	offset := (pagination.Page - 1) * pagination.PerPage

	var products []database.Product = make([]database.Product, 0)
	var totalItems int64

	result := repository.Database.Model(&database.Product{}).Count(&totalItems)

	if result.Error != nil {
		return dtos.ApiPaginationResponse{
			Records:    products,
			TotalPages: 0,
			ItemsCount: 0,
		}, result.Error
	}

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pagination.PerPage)))

	result = repository.Database.Offset(offset).Limit(pagination.PerPage).Find(&products)

	if result.Error != nil {
		return dtos.ApiPaginationResponse{
			Records:    products,
			TotalPages: totalPages,
			ItemsCount: totalItems,
		}, result.Error
	}

	return dtos.ApiPaginationResponse{
		Records:    products,
		TotalPages: totalPages,
		ItemsCount: totalItems,
	}, nil
}
