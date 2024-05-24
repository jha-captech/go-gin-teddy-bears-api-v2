package entity

import "gorm.io/gorm"

type PicnicLocation struct {
	gorm.Model
	LocationName string `gorm:"unique;not null"     json:"location_name"`
	Capacity     uint   `gorm:"not null;default:25" json:"capacity"`
	Municipality string `gorm:"not null"            json:"municipality"`
}
