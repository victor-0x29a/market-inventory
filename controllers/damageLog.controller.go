package controllers

import (
	"github.com/gofiber/fiber/v3"
	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
	"github.com/market-inventory/services"
	"github.com/market-inventory/utils"
)

type DamageLogController struct {
	App     *fiber.App
	Service *services.DamageLogService
}

func (controller DamageLogController) Initialize() {
	v1 := controller.App.Group("/v1/damage-log")

	v1.Get("/", damageLogGetFindAllV1(controller.Service))
	v1.Post("/", damageLogPostV1(controller.Service))

	v1.Get("/reasons", damageLogReasonsGetV1)
}

func damageLogPostV1(service *services.DamageLogService) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload dtos.CreateDamageLogDTO

		structErr, statusCode := utils.ValidateStruct(&payload, c.Body())

		if structErr != nil {
			return c.Status(statusCode).JSON(structErr)
		}

		err := service.CreateV1(&payload)

		if err != nil {
			payload, statusCode := utils.ParseCommonError(err)

			return c.Status(statusCode).JSON(payload)
		}

		return c.SendStatus(201)
	}
}

func damageLogGetFindAllV1(service *services.DamageLogService) fiber.Handler {
	return PaginatedRoute(service.FindAllV1)
}

func damageLogReasonsGetV1(c fiber.Ctx) error {
	return c.Status(200).JSON(constants.ReasonsMap)
}
