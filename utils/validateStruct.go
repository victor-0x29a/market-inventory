package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	dtos "github.com/market-inventory/DTOs"
)

func ValidateStruct(initializedVar any, c fiber.Ctx) (*dtos.ApiError, int) {
	err := json.Unmarshal(c.Body(), initializedVar)

	if err != nil {
		log.Fatal(err)
		return &dtos.ApiError{
			Errors: []string{"internal error"},
		}, 500
	}

	validate := validator.New()

	err = validate.Struct(initializedVar)

	if err != nil {
		parsedErrors := strings.Split(err.Error(), "\n")
		return &dtos.ApiError{
			Errors: parsedErrors,
		}, 422
	}

	return nil, 0
}
