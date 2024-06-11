package database

import (
	"fmt"
	"log/slog"
	"time"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database/entity"

	"gorm.io/gorm"
)

type Database struct {
	Session *gorm.DB
}

// Establish database connection and migrate tables before returning database.Database struct
func MustNewDatabase(c config.Configuration, d gorm.Dialector, gc gorm.Config, retryTimes int) Database {
	var (
		db         *gorm.DB
		err        error
		retryCount int
	)

	slog.Info("Attempting to connect to database")

	err = func() error {
		for i := 0; i <= retryTimes; i++ {
			db, err = gorm.Open(d, &gc)

			if err == nil {
				retryCount = i
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
		panic(fmt.Sprintf("dataBaseConnect: %v", err))
	}

	slog.Info("Database connection established", "Retry count", retryCount)

	db.AutoMigrate(
		&entity.TeddyBear{},
		&entity.Picnic{},
		&entity.PicnicLocation{},
		&entity.PicnicParticipant{},
	)

	return Database{Session: db}
}
