package logic

import (
	"errors"
	"fmt"

	m "teddy_bears_api_v2/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (logic Logic) ListTeddyBears() ([]m.TeddyBearReturn, error) {
	var bears []m.TeddyBear
	err := logic.Db.
		Model(&m.TeddyBear{}).
		Preload("Picnics").
		Find(&bears).
		Error
	if err != nil {
		return nil, fmt.Errorf("error retrieving picnic locations: %s", err)
	}

	var bearsReturn []m.TeddyBearReturn
	for _, bear := range bears {
		bearReturn := m.MapTeddyBearToOutput(bear)
		bearsReturn = append(bearsReturn, bearReturn)
	}

	return bearsReturn, nil
}

func (logic Logic) FetchTeddyBearByName(
	name string,
) (*m.TeddyBearReturn, error) {
	var bear m.TeddyBear
	err := logic.Db.Model(&m.TeddyBear{}).
		Where("name = ?", name).
		Preload("Picnics").
		First(&bear).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf(
			"error retrieving picnic locations: %s",
			err,
		)
	}

	bearReturn := m.MapTeddyBearToOutput(bear)

	return &bearReturn, nil
}

func (logic Logic) UpdateTeddyBearByName(
	name string,
	inputBear m.TeddyBearInput,
) (*m.TeddyBearReturn, error) {
	bear := m.MapInputToTeddyBear(inputBear)

	tx := logic.Db.Begin()

	// update item
	err := tx.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("name = ?", name).
		Omit("Picnics").
		Updates(&bear).
		Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(
			"error retrieving picnic locations: %s",
			err,
		)
	}

	// get updated item
	err = tx.
		Where("name = ?", name).
		Omit("Picnics").
		First(&bear).
		Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(
			"error retrieving updated teddy bear: %s",
			err,
		)
	}

	// updated associated tables
	err = tx.
		Model(&bear).
		Association("Picnics").
		Replace(&bear.Picnics)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(
			"error retrieving picnic locations: %s",
			err,
		)
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(
			"error retrieving picnic locations: %s",
			err,
		)
	}

	bearReturn, err := logic.FetchTeddyBearByName(name)
	if err != nil {
		return nil, err
	}

	return bearReturn, nil
}

func (logic Logic) CreateTeddyBear(inputBear m.TeddyBearInput) (int, error) {
	// set new user
	bear := m.MapInputToTeddyBear(inputBear)

	// add new row
	err := logic.Db.
		Create(&bear).
		Save(&bear).
		Error
	if err != nil {
		return 0, fmt.Errorf("error creating picnic location: %s", err)
	}

	return int(bear.ID), nil
}

func (logic Logic) DeleteTeddyBearByName(name string) error {
	// get record
	var bear m.TeddyBear
	err := logic.Db.Model(&m.TeddyBear{}).
		Where("name = ?", name).
		Omit("Picnics").
		First(&bear).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return fmt.Errorf(
			"error retrieving picnic locations: %s",
			err,
		)
	}

	// delete record
	err = logic.Db.
		Select(clause.Associations).
		Delete(&bear).
		Error
	if err != nil {
		return fmt.Errorf("error deleting picnic location: %s", err)
	}

	return nil
}
