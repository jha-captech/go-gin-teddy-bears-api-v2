package models

// PicnicLocation model
type PicnicLocation struct {
	ID           uint   `gorm:"primaryKey"          json:"id"`
	LocationName string `gorm:"unique;not null"     json:"location_name"`
	Capacity     uint   `gorm:"not null;default:25" json:"capacity"`
	Municipality string `gorm:"not null"            json:"municipality"`
}

type PicnicLocationInput struct {
	LocationName string `json:"location_name"`
	Capacity     int    `json:"capacity"`
	Municipality string `json:"municipality"`
}

func MapInputToPicnicLocation(loc PicnicLocationInput) PicnicLocation {
	return PicnicLocation{
		LocationName: loc.LocationName,
		Capacity:     uint(loc.Capacity),
		Municipality: loc.Municipality,
	}
}
