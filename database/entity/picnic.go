package entity

import (
	"time"

	"gorm.io/gorm"
)

type Picnic struct {
	gorm.Model
	PicnicName   string       `gorm:"unique;not null"                json:"picnic_name"`
	LocationID   uint         `                                      json:"location_id"`
	StartTime    time.Time    `gorm:"not null"                       json:"start_time"`
	HasMusic     bool         `gorm:"not null;default:true"          json:"has_music"`
	HasFood      bool         `gorm:"not null;default:true"          json:"has_food"`
	Participants []*TeddyBear `gorm:"many2many:picnic_participants;" json:"participants"`
}
