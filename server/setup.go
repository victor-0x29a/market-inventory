package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/market-inventory/config"
	controllers "github.com/market-inventory/controllers"
	"github.com/market-inventory/database"
	"github.com/market-inventory/repositories"
	"github.com/market-inventory/services"
)

func Setup() (*fiber.App, *config.Config) {
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

	productRepository := repositories.ProductRepository{
		Database: db,
	}
	productService := services.ProductService{
		Repository: &productRepository,
	}

	damageLogRepository := repositories.DamageLogRepository{
		Database: db,
	}
	damageLogService := services.DamageLogService{
		Repository:        &damageLogRepository,
		ProductRepository: &productRepository,
	}

	app := fiber.New()

	productController := controllers.ProductController{
		App:     app,
		Service: &productService,
	}
	damageLogController := controllers.DamageLogController{
		App:     app,
		Service: &damageLogService,
	}

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

	productController.Initialize()
	damageLogController.Initialize()

	return app, conf
}
