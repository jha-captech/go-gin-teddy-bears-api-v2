package logic

import (
	"fmt"

	"teddy_bears_api_v2/database/entity"
)

/* --------------------------------- structs -------------------------------- */

type PicnicLocationInput struct {
	LocationName string `json:"location_name"`
	Capacity     int    `json:"capacity"`
	Municipality string `json:"municipality"`
}

type PicnicLocationReturn struct {
	Id           uint   `json:"id"`
	LocationName string `json:"location_name"`
	Capacity     int    `json:"capacity"`
	Municipality string `json:"municipality"`
}

/* -------------------------------- exported -------------------------------- */

// List all locations from a database session.
func (l Logic) ListLocations() ([]PicnicLocationReturn, error) {
	locations, err := l.DB.ListLocations()
	if err != nil {
		return nil, fmt.Errorf("ListLocations: %w", err)
	}

	var locationsReturn []PicnicLocationReturn
	for _, location := range locations {
		locationReturn := mapPicnicLocationToOutput(location)
		locationsReturn = append(locationsReturn, locationReturn)
	}

	return locationsReturn, nil
}

// Fetch a location by ID from a database session.
func (l Logic) FetchLocationByID(id int) (PicnicLocationReturn, error) {
	location, err := l.DB.FetchLocationByID(id)
	if err != nil {
		return PicnicLocationReturn{}, fmt.Errorf("FetchLocationByID: %w", err)
	}

	locationReturn := mapPicnicLocationToOutput(location)

	return locationReturn, nil
}

// Update a location by id in a database session.
func (l Logic) UpdateLocationByID(id int, locInput PicnicLocationInput) (PicnicLocationReturn, error) {
	loc := mapInputToPicnicLocation(locInput)
	location, err := l.DB.UpdateLocationByID(id, loc)
	if err != nil {
		return PicnicLocationReturn{}, fmt.Errorf("UpdateLocationByID: %w", err)
	}

	locationReturn := mapPicnicLocationToOutput(location)

	return locationReturn, nil
}

// Create a location in a database session.
func (l Logic) CreateLocation(location PicnicLocationInput) (int, error) {
	// set new user
	convertedLoc := mapInputToPicnicLocation(location)

	// add new row
	id, err := l.DB.CreateLocation(convertedLoc)
	if err != nil {
		return 0, fmt.Errorf("CreateLocation: %w", err)
	}

	return id, nil
}

// Delete a location by ID in a database session.
func (l Logic) DeleteLocationByID(id int) error {
	if err := l.DB.DeleteLocationByID(id); err != nil {
		return fmt.Errorf("DeleteLocationByID: %w", err)
	}

	return nil
}

/* -------------------------------- internal -------------------------------- */

// Map struct entity.PicnicLocation to output struct logic.PicnicLocationReturn
func mapPicnicLocationToOutput(loc entity.PicnicLocation) PicnicLocationReturn {
	return PicnicLocationReturn{
		Id:           loc.ID,
		LocationName: loc.LocationName,
		Capacity:     int(loc.Capacity),
		Municipality: loc.Municipality,
	}
}

// Map input struct logic.PicnicLocationReturn to struct entity.PicnicLocation
func mapInputToPicnicLocation(loc PicnicLocationInput) entity.PicnicLocation {
	return entity.PicnicLocation{
		LocationName: loc.LocationName,
		Capacity:     uint(loc.Capacity),
		Municipality: loc.Municipality,
	}
}
