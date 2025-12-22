package repositories

import (
	"context"

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

	return err
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
	return PaginateEntity[database.Product](pagination, repository.Database)
}

func (repository ProductRepository) Update(productId int, updatedData dtos.UpdateProductDTO) error {
	result := repository.Database.Model(&database.Product{}).Where("id = ?", productId)

	if result.Error != nil {
		return result.Error
	}

	result = result.Updates(updatedData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repository ProductRepository) Delete(productId int) error {
	result := repository.Database.Where("id = ?", productId).Delete(&database.Product{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
