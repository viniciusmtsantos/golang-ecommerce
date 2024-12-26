package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	ServerPort string
}

func SetupEnvironment() (cfg ApplicationConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	return ApplicationConfig{ServerPort: httpPort}, nil
}
