package logic

import (
	"errors"
	"fmt"

	m "teddy_bears_api_v2/models"

	"gorm.io/gorm"
)

func (logic Logic) ListLocations() ([]m.PicnicLocation, error) {
	var locations []m.PicnicLocation
	if err := logic.Db.Find(&locations).Error; err != nil {
		return nil, fmt.Errorf("error retrieving picnic locations: %s", err)
	}

	return locations, nil
}

func (logic Logic) FetchLocationByID(id int) (*m.PicnicLocation, error) {
	var location m.PicnicLocation
	if err := logic.Db.Where("Id = ?", id).First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return &m.PicnicLocation{}, fmt.Errorf(
			"error retrieving picnic location: %s",
			err,
		)
	}

	return &location, nil
}

func (logic Logic) UpdateLocationByID(
	id int,
	loc m.PicnicLocationInput,
) (*m.PicnicLocation, error) {
	err := logic.Db.Model(&m.PicnicLocation{}).
		Where("id = ?", id).
		Updates(&loc).
		Error
	if err != nil {
		return nil, fmt.Errorf(
			"error updating picnic location: %s",
			err,
		)
	}

	// Fetch and return the updated location
	updatedLocation, err := logic.FetchLocationByID(id)
	if err != nil {
		return nil, err
	}

	return updatedLocation, nil
}

func (logic Logic) CreateLocation(location m.PicnicLocationInput) (int, error) {
	// set new user
	convertedLoc := m.MapInputToPicnicLocation(location)

	// add new row
	if err := logic.Db.Create(&convertedLoc).Error; err != nil {
		return 0, fmt.Errorf("error creating picnic location: %s", err)
	}

	return int(convertedLoc.ID), nil
}

func (logic Logic) DeleteLocationByID(id int) error {
	if err := logic.Db.Where("id = ?", id).Delete(&m.PicnicLocation{}).Error; err != nil {
		return fmt.Errorf("error deleting picnic location: %s", err)
	}

	return nil
}
