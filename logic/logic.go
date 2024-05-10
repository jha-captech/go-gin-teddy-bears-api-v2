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

// TODO add db abstraction

func InitLogic(config *config.Config, dbOpen gorm.Dialector) (*Logic, error) {
	db, err := dbSetup(dbOpen)
	if err != nil {
		return nil, fmt.Errorf("error establishing DB connection: %s", err)
	}

	return &Logic{
		Db: db,
	}, nil
}

func dbSetup(dbOpen gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(
		dbOpen,
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

	return db, err
}

func SqliteOpen(config *config.Config) gorm.Dialector {
	return sqlite.Open(config.Database.Name)
}
