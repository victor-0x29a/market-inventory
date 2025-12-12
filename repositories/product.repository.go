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
