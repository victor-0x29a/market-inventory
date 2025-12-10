package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/market-inventory/config"
	"github.com/market-inventory/database"
)

func main() {
	conf, err := config.Load()

	if err != nil {
		panic("Error when loading the app config")
	}

	db, err := database.GetConnection(conf)

	if err != nil {
		panic("Error on app starting")
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("Error on app starting (2)")
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			panic("Error on app ending")
		}
	}()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		if err := sqlDB.Ping(); err != nil {
			return c.Status(500).JSON(map[string]string{
				"error": "001",
			})
		}

		return c.Status(200).JSON(map[string]string{
			"hello": "world",
		})
	})

	listenArg := fmt.Sprintf(":%d", conf.API_PORT)

	log.Fatal(app.Listen(listenArg))
}
