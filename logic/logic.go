package logic

import (
	"fmt"
	"log/slog"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Logic struct {
	Db *gorm.DB
}

func InitLogic(config *config.Config) (*Logic, error) {
	db, err := gorm.Open(
		sqlite.Open(config.Database.Name),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("error establishing DB connection: %s", err)
	}

	slog.Info("Database connection established")

	db.AutoMigrate(
		&models.TeddyBear{},
		&models.Picnic{},
		&models.PicnicLocation{},
		&models.PicnicParticipant{},
	)

	return &Logic{Db: db}, nil
}
