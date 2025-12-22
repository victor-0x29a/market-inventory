package services

import (
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
	"github.com/market-inventory/repositories"
)

type DamageLogService struct {
	Repository        *repositories.DamageLogRepository
	ProductRepository *repositories.ProductRepository
}

func (service DamageLogService) CreateV1(data *dtos.CreateDamageLogDTO) error {
	_, err := service.ProductRepository.FindOne(data.ProductId)

	if err != nil {
		return constants.ErrProductNotFound
	}

	err = service.Repository.Create(data)

	return err
}
