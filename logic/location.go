package logic

import (
	"errors"
	"fmt"

	"teddy_bears_api_v2/models"

	"gorm.io/gorm"
)

func (l Logic) ListLocations() ([]models.PicnicLocation, error) {
	var locations []models.PicnicLocation
	if err := l.DB.Find(&locations).Error; err != nil {
		return nil, fmt.Errorf("error retrieving picnic locations: %s", err)
	}

	return locations, nil
}

func (l Logic) FetchLocationByID(id int) (*models.PicnicLocation, error) {
	var location models.PicnicLocation
	if err := l.DB.Where("Id = ?", id).First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return &models.PicnicLocation{}, fmt.Errorf(
			"error retrieving picnic location: %s",
			err,
		)
	}

	return &location, nil
}

func (l Logic) UpdateLocationByID(
	id int,
	loc models.PicnicLocationInput,
) (*models.PicnicLocation, error) {
	err := l.DB.
		Model(&models.PicnicLocation{}).
		Where("id = ?", id).
		Updates(&loc).
		Error
	if err != nil {
		return nil, fmt.Errorf("UpdateLocationByID: %w", err)
	}

	// Fetch and return the updated location
	updatedLocation, err := l.FetchLocationByID(id)
	if err != nil {
		return nil, fmt.Errorf("UpdateLocationByID: %w", err)
	}

	return updatedLocation, nil
}

func (l Logic) CreateLocation(location models.PicnicLocationInput) (int, error) {
	// set new user
	convertedLoc := models.MapInputToPicnicLocation(location)

	// add new row
	if err := l.DB.Create(&convertedLoc).Error; err != nil {
		return 0, fmt.Errorf("CreateLocation: %w", err)
	}

	return int(convertedLoc.ID), nil
}

func (l Logic) DeleteLocationByID(id int) error {
	if err := l.DB.Where("id = ?", id).Delete(&models.PicnicLocation{}).Error; err != nil {
		return fmt.Errorf("DeleteLocationByID: %w", err)
	}

	return nil
}
