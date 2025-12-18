package repositories

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/services"
	"github.com/market-inventory/utils"
)

type ProductController struct {
	App     *fiber.App
	Service *services.ProductService
}

func (controller ProductController) Initialize() {
	v1 := controller.App.Group("/v1/product")

	v1.Post("/", postV1(controller.Service))
	v1.Get("/", getFindAllV1(controller.Service))
	v1.Get("/:productId", getFindOneV1(controller.Service))
}

func postV1(service *services.ProductService) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload dtos.CreateProductDTO

		structErr, statusCode := utils.ValidateStruct(&payload, c.Body())

		if structErr != nil {
			return c.Status(statusCode).JSON(structErr)
		}

		err := service.CreateV1(&payload)

		if err != nil {
			payload, statusCode := utils.ParseCommonError(err)

			return c.Status(statusCode).JSON(payload)
		}

		return c.SendStatus(204)
	}
}

func getFindOneV1(service *services.ProductService) fiber.Handler {
	return func(c fiber.Ctx) error {
		productId, _ := strconv.Atoi(c.Params("productId"))

		params := dtos.FindOneProductDTO{
			ID: productId,
		}

		errorResponse, statusCode := utils.Validator(params)

		if errorResponse != nil {
			return c.Status(statusCode).JSON(errorResponse)
		}

		product, err := service.FindOneV1(params.ID)

		if err != nil {
			errorResponse, statusCode := utils.ParseCommonError(err)

			return c.Status(statusCode).JSON(errorResponse)
		}

		return c.Status(200).JSON(product)
	}
}

func getFindAllV1(service *services.ProductService) fiber.Handler {
	return func(c fiber.Ctx) error {
		pagination := utils.ValidatePagination(c.Query("Page"), c.Query("PerPage"))

		data, err := service.FindAllV1(*pagination)

		if err != nil {
			errorResponse, statusCode := utils.ParseCommonError(err)

			return c.Status(statusCode).JSON(errorResponse)
		}

		return c.Status(200).JSON(data)
	}
}
