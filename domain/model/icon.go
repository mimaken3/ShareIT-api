package model

type Icon struct {
	IconID   uint   `gorm:"primary_key" json:"icon_id"`
	UserID   uint   `json:"user_id"`
	IconName string `gorm:"size:255" json:"icon_name"`
}
