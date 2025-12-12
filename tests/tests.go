package tests

import (
	"github.com/gofiber/fiber/v3"
	"github.com/market-inventory/server"
)

func Setup() *fiber.App {
	app, _, _ := server.Setup()

	return app
}
