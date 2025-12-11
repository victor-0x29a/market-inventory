package utils

import (
	"log"

	dtos "github.com/market-inventory/DTOs"
)

func ParseCommonError(err error) (*dtos.ApiError, int) {
	log.Fatal(err)
	payload := &dtos.ApiError{
		Errors: []string{"internal error"},
	}
	return payload, 500
}
