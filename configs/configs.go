package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	ServerPort         string
	DSN                string
	Secret             string
	AWSAccessKeyID     string
	AWSAccessKeySecret string
	AWSRegion          string
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

	awsAccessKeyID := os.Getenv("AWS_SNS_ACCESS_KEY_ID")
	if len(appSecret) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	awsAccessKeySecret := os.Getenv("AWS_SNS_SECRET_ACCESS_KEY")
	if len(appSecret) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	awsRegion := os.Getenv("AWS_REGION")
	if len(appSecret) < 1 {
		return ApplicationConfig{}, errors.New("env variables not found")
	}

	return ApplicationConfig{
		ServerPort:         httpPort,
		DSN:                dsn,
		Secret:             appSecret,
		AWSAccessKeyID:     awsAccessKeyID,
		AWSAccessKeySecret: awsAccessKeySecret,
		AWSRegion:          awsRegion,
	}, nil
}
