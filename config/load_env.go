package config

import (
	"database/sql"
	"log/slog"

	"github.com/joho/godotenv"
)

type Env struct {
	Db *sql.DB
}

func DotEnvInit() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("DOTENV: No .env file found or error reading file.")
	} else {
		slog.Info("DOTENV: .env found.")
	}
}
