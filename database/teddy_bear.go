package database

import (
	"errors"
	"fmt"

	"teddy_bears_api_v2/database/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db Database) ListTeddyBears() ([]entity.TeddyBear, error) {
	var bears []entity.TeddyBear
	err := db.Session.
		Model(&entity.TeddyBear{}).
		Preload("Picnics").
		Find(&bears).
		Error
	if err != nil {
		return nil, fmt.Errorf("ListTeddyBears: %w", err)
	}

	return bears, nil
}

func (db Database) ListPaginatedTeddyBears(offset, limit int) ([]entity.TeddyBear, error) {
	var bears []entity.TeddyBear
	err := db.Session.
		Model(&entity.TeddyBear{}).
		Preload("Picnics").
		Offset(offset).
		Limit(limit).
		Find(&bears).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("ListPaginatedTeddyBears: %w", err)
	}

	return bears, nil
}

func (db Database) FetchTeddyBearByName(name string) (entity.TeddyBear, error) {
	var bear entity.TeddyBear
	err := db.Session.Model(&entity.TeddyBear{}).
		Where("name = ?", name).
		Preload("Picnics").
		First(&bear).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.TeddyBear{}, nil
		}

		return entity.TeddyBear{}, fmt.Errorf("FetchTeddyBearByName: %w", err)
	}

	return bear, nil
}

func (db Database) UpdateTeddyBearByName(name string, bear entity.TeddyBear) (entity.TeddyBear, error) {
	transaction := db.Session.Begin()

	// update item
	err := transaction.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("name = ?", name).
		Omit("Picnics").
		Updates(&bear).
		Error
	if err != nil {
		transaction.Rollback()
		return entity.TeddyBear{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// get updated item
	err = transaction.
		Where("name = ?", name).
		Omit("Picnics").
		First(&bear).
		Error
	if err != nil {
		transaction.Rollback()
		return entity.TeddyBear{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// updated associated tables
	err = transaction.
		Model(&bear).
		Association("Picnics").
		Replace(&bear.Picnics)
	if err != nil {
		transaction.Rollback()
		return entity.TeddyBear{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// commit changes
	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		return entity.TeddyBear{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	bearReturn, err := db.FetchTeddyBearByName(name)
	if err != nil {
		return entity.TeddyBear{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	return bearReturn, nil
}

func (db Database) CreateTeddyBear(bear entity.TeddyBear) (int, error) {
	// add new row
	err := db.Session.
		Create(&bear).
		Save(&bear).
		Error
	if err != nil {
		return 0, fmt.Errorf("CreateTeddyBear: %w", err)
	}

	return int(bear.ID), nil
}

func (db Database) DeleteTeddyBearByName(name string) error {
	// get record
	var bear entity.TeddyBear
	err := db.Session.
		Model(&entity.TeddyBear{}).
		Where("name = ?", name).
		Omit("Picnics").
		First(&bear).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return fmt.Errorf("DeleteTeddyBearByName: %w", err)
	}

	// delete record
	err = db.Session.
		Select(clause.Associations).
		Delete(&bear).
		Error
	if err != nil {
		return fmt.Errorf("DeleteTeddyBearByName: %w", err)
	}

	return nil
}
