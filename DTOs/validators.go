package dtos

import (
	"github.com/go-playground/validator/v10"
	"github.com/market-inventory/constants"
)

func ValidateReason(fl validator.FieldLevel) bool {
	reasonID := int(fl.Field().Int())

	_, ok := constants.ReasonsMap[reasonID]

	return ok
}
