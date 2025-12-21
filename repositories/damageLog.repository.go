package repositories

import (
	"context"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/database"
	"gorm.io/gorm"
)

type DamageLogRepository struct {
	Database *gorm.DB
}

func (repository DamageLogRepository) Create(data dtos.CreateDamageLogDTO) error {
	ctx := context.Background()

	err := gorm.G[database.DamageLog](repository.Database).Create(ctx, &database.DamageLog{
		ProductID: uint(data.ProductId),
		Quantity:  data.Quantity,
		Reason:    data.Reason,
	})

	return err
}
