package database

import (
	"fmt"
	"log/slog"
	"time"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/models"

	"gorm.io/gorm"
)

func Connect(
	config config.Configuration,
	dialector gorm.Dialector,
	formConfig gorm.Config,
	retryTimes int,
) (db *gorm.DB, err error) {
	err = func() error {
		for i := 0; i <= retryTimes; i++ {
			db, err = gorm.Open(dialector, &formConfig)

			if err == nil {
				return nil
			}

			if i == retryTimes {
				return err
			}

			time.Sleep(1 * time.Second)
		}
		return err
	}()
	if err != nil {
		return nil, fmt.Errorf("dataBaseConnect: %w", err)
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
