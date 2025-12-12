package database

import (
	"fmt"
	"log"

	"github.com/market-inventory/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DATABASE_HOST, cfg.DATABASE_USER, cfg.DATABASE_PASSWORD, cfg.DATABASE_DB, cfg.DATABASE_PORT)

	var dialector gorm.Dialector

	if cfg.ENVIRONMENT == "test" {
		dsn = ":memory:"
		dialector = sqlite.Open(dsn)
	} else {
		dialector = postgres.Open(dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal("App external connection is failed.")
	}

	db.AutoMigrate(&Product{})

	return db, err
}
