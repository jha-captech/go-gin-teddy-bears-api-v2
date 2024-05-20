package logic

import (
	"gorm.io/gorm"
)

type Logic struct {
	DB *gorm.DB
}

func InitLogic(dbConnection *gorm.DB) (*Logic, error) {
	return &Logic{DB: dbConnection}, nil
}
