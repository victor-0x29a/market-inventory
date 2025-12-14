package utils

import (
	"strconv"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
)

func ValidatePagination(page string, perPage string) *dtos.ApiPagination {
	parsedPage, err := strconv.Atoi(page)

	if err != nil {
		parsedPage = constants.DefaultPage
	}

	parsedPerPage, err := strconv.Atoi(perPage)

	if err != nil {
		parsedPerPage = constants.DefaultPerPage
	}

	return &dtos.ApiPagination{
		Page:    parsedPage,
		PerPage: parsedPerPage,
	}
}
