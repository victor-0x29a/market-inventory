package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
)

func Validator(refStruct any) (*dtos.ApiError, int) {
	validate := validator.New()

	err := validate.Struct(refStruct)

	if err != nil {
		parsedErrors := strings.Split(err.Error(), "\n")
		return &dtos.ApiError{
			Errors: parsedErrors,
		}, 422
	}

	return nil, 0
}

func ValidateStruct(initializedVar any, body []byte) (*dtos.ApiError, int) {
	err := json.Unmarshal(body, initializedVar)

	if err != nil {
		log.Println(err)
		return &dtos.ApiError{
			Errors: []string{constants.ErrInternal.Error()},
		}, 500
	}

	response, statusCode := Validator(initializedVar)

	return response, statusCode
}
