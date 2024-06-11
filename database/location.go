package database

import (
	"errors"
	"fmt"

	"teddy_bears_api_v2/database/entity"

	"gorm.io/gorm"
)

func (db Database) ListLocations() ([]entity.PicnicLocation, error) {
	var locations []entity.PicnicLocation
	if err := db.Session.Find(&locations).Error; err != nil {
		return nil, fmt.Errorf("ListLocations: %w", err)
	}

	return locations, nil
}

func (db Database) FetchLocationByID(id int) (entity.PicnicLocation, error) {
	var location entity.PicnicLocation
	err := db.Session.
		Where("Id = ?", id).
		First(&location).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.PicnicLocation{}, nil
		}

		return entity.PicnicLocation{}, fmt.Errorf("FetchLocationByID: %w", err)
	}

	return location, nil
}

func (db Database) UpdateLocationByID(id int, loc entity.PicnicLocation) (entity.PicnicLocation, error) {
	err := db.Session.
		Model(&entity.PicnicLocation{}).
		Where("id = ?", id).
		Updates(&loc).
		Error
	if err != nil {
		return entity.PicnicLocation{}, fmt.Errorf("UpdateLocationByID: %w", err)
	}

	// Fetch and return the updated location
	updatedLocation, err := db.FetchLocationByID(id)
	if err != nil {
		return entity.PicnicLocation{}, fmt.Errorf("UpdateLocationByID: %w", err)
	}

	return updatedLocation, nil
}

func (db Database) CreateLocation(location entity.PicnicLocation) (int, error) {
	// add new row
	if err := db.Session.Create(&location).Error; err != nil {
		return 0, fmt.Errorf("CreateLocation: %w", err)
	}

	return int(location.ID), nil
}

func (db Database) DeleteLocationByID(id int) error {
	err := db.Session.
		Where("id = ?", id).
		Delete(&entity.PicnicLocation{}).
		Error
	if err != nil {
		return fmt.Errorf("DeleteLocationByID: %w", err)
	}

	return nil
}
