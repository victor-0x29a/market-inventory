package repositories

import (
	"math"

	dtos "github.com/market-inventory/DTOs"
	"gorm.io/gorm"
)

func PaginateEntity[T any](pagination dtos.ApiPagination, Database *gorm.DB) (dtos.ApiPaginationResponse, error) {
	offset := (pagination.Page - 1) * pagination.PerPage

	var records []T = make([]T, 0)
	var totalItems int64
	var model T

	result := Database.Model(&model).Count(&totalItems)

	if result.Error != nil {
		return dtos.ApiPaginationResponse{
			Records:    records,
			TotalPages: 0,
			ItemsCount: 0,
		}, result.Error
	}

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pagination.PerPage)))

	result = Database.Offset(offset).Limit(pagination.PerPage).Find(&records)

	if result.Error != nil {
		return dtos.ApiPaginationResponse{
			Records:    records,
			TotalPages: totalPages,
			ItemsCount: totalItems,
		}, result.Error
	}

	return dtos.ApiPaginationResponse{
		Records:    records,
		TotalPages: totalPages,
		ItemsCount: totalItems,
	}, nil
}
