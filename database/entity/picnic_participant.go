package entity

type PicnicParticipant struct {
	PicnicID    uint `gorm:"primaryKey" json:"picnic_id"`
	TeddyBearID uint `gorm:"primaryKey" json:"teddy_bear_id"`
}
