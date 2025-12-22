package controllers

import (
	"github.com/gofiber/fiber/v3"
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/utils"
)

func PaginatedRoute(findAll func(pagination dtos.ApiPagination) (dtos.ApiPaginationResponse, error)) fiber.Handler {
	return func(c fiber.Ctx) error {
		pagination := utils.ValidatePagination(c.Query("Page"), c.Query("PerPage"))

		data, err := findAll(*pagination)

		if err != nil {
			errorResponse, statusCode := utils.ParseCommonError(err)

			return c.Status(statusCode).JSON(errorResponse)
		}

		return c.Status(200).JSON(data)
	}
}
