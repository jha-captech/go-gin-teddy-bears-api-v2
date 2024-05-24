package entity

import "gorm.io/gorm"

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
