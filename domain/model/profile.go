package model

type Profile struct {
	ProfileID uint   `gorm:"primary_key" json:"profile_id"`
	UserID    uint   `json:"user_id"`
	Content   string `gorm:"size:1000" json:"content"`
	IsDeleted int8   `json:"-"`
}
