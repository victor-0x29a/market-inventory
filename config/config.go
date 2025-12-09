package config

import (
	"os"
	"strconv"

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
	dbHost := os.Getenv("DATABASE_HOST")
	unparsedDbPort := os.Getenv("DATABASE_PORT")
	dbPort, err := strconv.Atoi(unparsedDbPort)

	if err != nil {
		panic("Missing environment var. (1)")
	}

	config := Config{
		DATABASE_USER:     dbUser,
		DATABASE_PASSWORD: dbPassword,
		DATABASE_DB:       dbDatabase,
		DATABASE_HOST:     dbHost,
		DATABASE_PORT:     dbPort,
	}

	return &config, nil
}
