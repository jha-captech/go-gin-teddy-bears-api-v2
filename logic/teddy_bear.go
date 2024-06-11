package logic

import (
	"fmt"

	"teddy_bears_api_v2/database/entity"

	"gorm.io/gorm"
)

/* --------------------------------- structs -------------------------------- */

type TeddyBearInput struct {
	Name           string `json:"name"`
	PrimaryColor   string `json:"primary_color"`
	AccentColor    string `json:"accent_color"`
	IsDressed      bool   `json:"is_dressed"`
	OwnerName      string `json:"owner_name"`
	Characteristic string `json:"characteristic"`
	PicnicIDs      []int  `json:"picnic_ids"`
}

type TeddyBearReturn struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	PrimaryColor   string `json:"primary_color"`
	AccentColor    string `json:"accent_color"`
	IsDressed      bool   `json:"is_dressed"`
	OwnerName      string `json:"owner_name"`
	Characteristic string `json:"characteristic"`
	PicnicIDs      []int  `json:"picnic_ids"`
}

/* -------------------------------- exported -------------------------------- */
// List all teddy bears from a database session.
func (l Logic) ListTeddyBears() ([]TeddyBearReturn, error) {
	bears, err := l.DB.ListTeddyBears()
	if err != nil {
		return nil, fmt.Errorf("ListTeddyBears: %w", err)
	}

	var bearsReturn []TeddyBearReturn
	for _, bear := range bears {
		bearReturn := mapTeddyBearToOutput(bear)
		bearsReturn = append(bearsReturn, bearReturn)
	}

	return bearsReturn, nil
}

// List all teddy bears from a database session with pagination.
func (l Logic) ListPaginatedTeddyBears(page, limit int) ([]TeddyBearReturn, error) {
	offset := (page - 1) * limit
	bears, err := l.DB.ListPaginatedTeddyBears(offset, limit)
	if err != nil {
		return nil, fmt.Errorf("ListPaginatedTeddyBears: %w", err)
	}

	var bearsReturn []TeddyBearReturn
	for _, bear := range bears {
		bearReturn := mapTeddyBearToOutput(bear)
		bearsReturn = append(bearsReturn, bearReturn)
	}

	return bearsReturn, nil
}

// Fetch a teddy bear by name from a database session.
func (l Logic) FetchTeddyBearByName(name string) (TeddyBearReturn, error) {
	bear, err := l.DB.FetchTeddyBearByName(name)
	if err != nil {
		return TeddyBearReturn{}, fmt.Errorf("ListPaginatedTeddyBears: %w", err)
	}

	bearReturn := mapTeddyBearToOutput(bear)

	return bearReturn, nil
}

// Update a teddy bear by name in a database session.
func (l Logic) UpdateTeddyBearByName(name string, inputBear TeddyBearInput) (TeddyBearReturn, error) {
	bear := mapInputToTeddyBear(inputBear)

	bear, err := l.DB.UpdateTeddyBearByName(name, bear)
	if err != nil {
		return TeddyBearReturn{}, fmt.Errorf("UpdateTeddyBearByName: %w", err)
	}

	bearReturn := mapTeddyBearToOutput(bear)

	return bearReturn, nil
}

// Create a teddy bear in a database session.
func (l Logic) CreateTeddyBear(inputBear TeddyBearInput) (int, error) {
	// set new user
	bear := mapInputToTeddyBear(inputBear)

	// add new row
	id, err := l.DB.CreateTeddyBear(bear)
	if err != nil {
		return 0, fmt.Errorf("CreateTeddyBear: %w", err)
	}

	return id, nil
}

// Delete a teddy bear by name in a database session.
func (l Logic) DeleteTeddyBearByName(name string) error {
	if err := l.DB.DeleteTeddyBearByName(name); err != nil {
		return fmt.Errorf("DeleteTeddyBearByName: %w", err)
	}

	return nil
}

/* -------------------------------- internal -------------------------------- */

// Map struct entity.TeddyBear to output struct logic.TeddyBearReturn
func mapTeddyBearToOutput(bear entity.TeddyBear) TeddyBearReturn {
	var ids []int
	for _, location := range bear.Picnics {
		ids = append(ids, int(location.ID))
	}

	return TeddyBearReturn{
		Id:             bear.ID,
		Name:           bear.Name,
		PrimaryColor:   bear.PrimaryColor,
		AccentColor:    bear.AccentColor,
		IsDressed:      bear.IsDressed,
		OwnerName:      bear.OwnerName,
		Characteristic: bear.Characteristic,
		PicnicIDs:      ids,
	}
}

// Map input struct logic.TeddyBearReturn to struct entity.TeddyBear
func mapInputToTeddyBear(bear TeddyBearInput) entity.TeddyBear {
	var picnics []*entity.Picnic
	for _, ID := range bear.PicnicIDs {
		picnics = append(picnics, &entity.Picnic{Model: gorm.Model{ID: uint(ID)}})
	}

	return entity.TeddyBear{
		Name:           bear.Name,
		PrimaryColor:   bear.PrimaryColor,
		AccentColor:    bear.AccentColor,
		IsDressed:      bear.IsDressed,
		OwnerName:      bear.OwnerName,
		Characteristic: bear.Characteristic,
		Picnics:        picnics,
	}
}
