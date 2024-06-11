package logic

import (
	"log/slog"

	"teddy_bears_api_v2/database"
)

type Logic struct {
	DB database.Database
}

// Setup and return logic.Logic struct.
func NewLogic(db database.Database) Logic {
	slog.Info("Setting up new Logic session")
	// panic implemented if other actions added here
	return Logic{DB: db}
}
