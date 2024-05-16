package logic

import (
	"fmt"
	"log/slog"
	"time"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Logic struct {
	Db *gorm.DB
}

type logicSetup struct {
	*config.Config
	gormConfig  *gorm.Config
	dbDialector dialectorFunc
}

type dialectorFunc func(*config.Config) gorm.Dialector

func InitLogic(c *config.Config, d dialectorFunc, g *gorm.Config) (*Logic, error) {
	ls := logicSetup{Config: c, dbDialector: d, gormConfig: g}

	db, err := ls.dataBaseConnect()
	if err != nil {
		return nil, err
	}

	return &Logic{Db: db}, nil
}

func (ls logicSetup) dataBaseConnect() (db *gorm.DB, err error) {
	err = func() error {
		for i := 0; i <= ls.Database.ConnectionRetry; i++ {
			db, err = gorm.Open(ls.dbDialector(ls.Config), ls.gormConfig)

			if err == nil {
				return nil
			}

			if i == ls.Database.ConnectionRetry {
				return err
			}

			time.Sleep(1 * time.Second)
		}
		return err
	}()
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

func SqliteOpen(c *config.Config) gorm.Dialector {
	return sqlite.Open(c.Database.Name)
}
