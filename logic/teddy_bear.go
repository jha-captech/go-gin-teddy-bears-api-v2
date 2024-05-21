package logic

import (
	"errors"
	"fmt"

	"teddy_bears_api_v2/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (l Logic) ListTeddyBears() ([]models.TeddyBearReturn, error) {
	var bears []models.TeddyBear
	err := l.DB.
		Model(&models.TeddyBear{}).
		Preload("Picnics").
		Find(&bears).
		Error
	if err != nil {
		return nil, fmt.Errorf("ListTeddyBears: %w", err)
	}

	var bearsReturn []models.TeddyBearReturn
	for _, bear := range bears {
		bearReturn := models.MapTeddyBearToOutput(bear)
		bearsReturn = append(bearsReturn, bearReturn)
	}

	return bearsReturn, nil
}

func (l Logic) ListPaginatedTeddyBears(
	page, limit int,
) ([]models.TeddyBearReturn, error) {
	var bears []models.TeddyBear
	offset := (page - 1) * limit
	err := l.DB.
		Model(&models.TeddyBear{}).
		Preload("Picnics").
		Offset(offset).
		Limit(limit).
		Find(&bears).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("ListPaginatedTeddyBears: %w", err)
	}

	var bearsReturn []models.TeddyBearReturn
	for _, bear := range bears {
		bearReturn := models.MapTeddyBearToOutput(bear)
		bearsReturn = append(bearsReturn, bearReturn)
	}

	return bearsReturn, nil
}

func (l Logic) FetchTeddyBearByName(
	name string,
) (*models.TeddyBearReturn, error) {
	var bear models.TeddyBear
	err := l.DB.Model(&models.TeddyBear{}).
		Where("name = ?", name).
		Preload("Picnics").
		First(&bear).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("FetchTeddyBearByName: %w", err)
	}

	bearReturn := models.MapTeddyBearToOutput(bear)

	return &bearReturn, nil
}

func (l Logic) UpdateTeddyBearByName(
	name string,
	inputBear models.TeddyBearInput,
) (*models.TeddyBearReturn, error) {
	bear := models.MapInputToTeddyBear(inputBear)

	tx := l.DB.Begin()

	// update item
	err := tx.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("name = ?", name).
		Omit("Picnics").
		Updates(&bear).
		Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// get updated item
	err = tx.
		Where("name = ?", name).
		Omit("Picnics").
		First(&bear).
		Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// updated associated tables
	err = tx.
		Model(&bear).
		Association("Picnics").
		Replace(&bear.Picnics)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	bearReturn, err := l.FetchTeddyBearByName(name)
	if err != nil {
		return nil, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	return bearReturn, nil
}

func (l Logic) CreateTeddyBear(inputBear models.TeddyBearInput) (int, error) {
	// set new user
	bear := models.MapInputToTeddyBear(inputBear)

	// add new row
	err := l.DB.
		Create(&bear).
		Save(&bear).
		Error
	if err != nil {
		return 0, fmt.Errorf("CreateTeddyBear: %w", err)
	}

	return int(bear.ID), nil
}

func (l Logic) DeleteTeddyBearByName(name string) error {
	// get record
	var bear models.TeddyBear
	err := l.DB.Model(&models.TeddyBear{}).
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
	err = l.DB.
		Select(clause.Associations).
		Delete(&bear).
		Error
	if err != nil {
		return fmt.Errorf("DeleteTeddyBearByName: %w", err)
	}

	return nil
}
