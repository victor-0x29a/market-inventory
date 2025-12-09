package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbDatabase := os.Getenv("DATABASE_DB")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	config := Config{
		DATABASE_USER:     dbUser,
		DATABASE_PASSWORD: dbPassword,
		DATABASE_DB:       dbDatabase,
	}

	return &config, nil
}
