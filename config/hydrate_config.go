package config

import "os"

func HydrateConfigFromEnv() *Config {
	return &Config{
		GoEnv: &GoEnv{
			Env: os.Getenv("GO_ENV"),
		},
		Database: &Database{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		HTTP: &HTTP{
			Domain: os.Getenv("GO_DOMAIN"),
			Port:   os.Getenv("GO_PORT"),
		},
	}
}
