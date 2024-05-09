package models

import "gorm.io/gorm"

// TeddyBear model
type TeddyBear struct {
	gorm.Model
	Name           string    `gorm:"unique; not null"               json:"name"`
	PrimaryColor   string    `gorm:"not null"                       json:"primary_color"`
	AccentColor    string    `                                      json:"accent_color"`
	IsDressed      bool      `gorm:"not null;default:true"          json:"is_dressed"`
	OwnerName      string    `gorm:"not null"                       json:"owner_name"`
	Characteristic string    `                                      json:"characteristic"`
	Picnics        []*Picnic `gorm:"many2many:picnic_participants;" json:"picnics"`
}

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

func MapTeddyBearToOutput(bear TeddyBear) TeddyBearReturn {
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

func MapInputToTeddyBear(bear TeddyBearInput) TeddyBear {
	var picnics []*Picnic
	for _, ID := range bear.PicnicIDs {
		picnics = append(picnics, &Picnic{Model: gorm.Model{ID: uint(ID)}})
	}

	return TeddyBear{
		Name:           bear.Name,
		PrimaryColor:   bear.PrimaryColor,
		AccentColor:    bear.AccentColor,
		IsDressed:      bear.IsDressed,
		OwnerName:      bear.OwnerName,
		Characteristic: bear.Characteristic,
		Picnics:        picnics,
	}
}
