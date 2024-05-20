package config

import (
	"os"
	"strconv"
)

func HydrateConfigFromEnv() *Config {
	return &Config{
		GoEnv: &GoEnv{
			Env: os.Getenv("GO_ENV"),
		},
		Database: &Database{
			User:            os.Getenv("DB_USER"),
			Password:        os.Getenv("DB_PASSWORD"),
			Host:            os.Getenv("DB_HOST"),
			Port:            os.Getenv("DB_PORT"),
			Name:            os.Getenv("DB_NAME"),
			ConnectionRetry: getEnvVarAsInt("DB_CONNECTION_RETRY", 0),
		},
		HTTP: &HTTP{
			Domain: os.Getenv("GO_DOMAIN"),
			Port:   os.Getenv("GO_PORT"),
		},
	}
}

func getEnvVarAsInt(key string, devVal int) int {
	value := os.Getenv(key)
	number, err := strconv.Atoi(value)
	if err != nil {
		return devVal
	}
	return number
}
