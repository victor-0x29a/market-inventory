package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	environment := os.Getenv("APP_ENVIRONMENT")

	if environment == "test" {
		return &Config{
			DATABASE_USER:     "",
			DATABASE_PASSWORD: "",
			DATABASE_DB:       "",
			DATABASE_HOST:     "",
			DATABASE_PORT:     3005,
			API_PORT:          3006,
			ENVIRONMENT:       environment,
		}, nil
	}

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbDatabase := os.Getenv("DATABASE_DB")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	unparsedDbPort := os.Getenv("DATABASE_PORT")
	unparsedPort := os.Getenv("API_PORT")
	dbPort, err := strconv.Atoi(unparsedDbPort)

	if err != nil {
		panic("Missing environment var. (1)")
	}

	apiPort, err := strconv.Atoi(unparsedPort)

	if err != nil {
		panic("Missing environment var. (1)")
	}

	config := Config{
		DATABASE_USER:     dbUser,
		DATABASE_PASSWORD: dbPassword,
		DATABASE_DB:       dbDatabase,
		DATABASE_HOST:     dbHost,
		DATABASE_PORT:     dbPort,
		API_PORT:          apiPort,
		ENVIRONMENT:       environment,
	}

	return &config, nil
}
