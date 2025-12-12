package config

type Config struct {
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_DB       string
	DATABASE_HOST     string
	DATABASE_PORT     int
	API_PORT          int
	ENVIRONMENT       string
}
