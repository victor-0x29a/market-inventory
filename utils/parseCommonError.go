package utils

import (
	"errors"
	"log"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
)

func ParseCommonError(err error) (*dtos.ApiError, int) {
	if errors.Is(err, constants.ErrProductNotFound) {
		return &dtos.ApiError{
			Errors: []string{constants.ErrProductNotFound.Error()},
		}, 404
	}

	log.Println(err)

	payload := &dtos.ApiError{
		Errors: []string{constants.ErrInternal.Error()},
	}
	return payload, 500
}
