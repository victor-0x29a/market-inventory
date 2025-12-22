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

func (service DamageLogService) FindAllV1(pagination dtos.ApiPagination) (dtos.ApiPaginationResponse, error) {
	damageLogs, err := service.Repository.FindAll(pagination)

	return damageLogs, err
}
