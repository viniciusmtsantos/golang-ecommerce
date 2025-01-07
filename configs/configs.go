package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	ServerPort string
	DSN        string
	Secret     string
}

func SetupEnvironment() (cfg ApplicationConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	return ApplicationConfig{ServerPort: httpPort, DSN: dsn, Secret: appSecret}, nil
}
