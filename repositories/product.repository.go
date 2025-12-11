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
