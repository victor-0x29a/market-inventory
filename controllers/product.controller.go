package repositories

import (
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
}

func postV1(service *services.ProductService) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload dtos.CreateProductDTO

		structErr, statusCode := utils.ValidateStruct(&payload, c)

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
