package model

import "time"

type User struct {
	UserID           uint      `gorm:"primary_key" json:"user_id"`
	IconName         string    `gorm:"size:500" json:"icon_name"`
	UserName         string    `gorm:"size:255" json:"user_name"`
	Email            string    `gorm:"size:255" json:"email"`
	Password         string    `gorm:"size:255" json:"password"`
	InterestedTopics string    `gorm:"size:255" json:"interested_topics"`
	Profile          string    `gorm:"size:1000" json:"profile"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedDate      time.Time `json:"updated_date"`
	DeletedDate      time.Time `json:"deleted_date"`
	IsDeleted        int8      `json:"-"`
}
