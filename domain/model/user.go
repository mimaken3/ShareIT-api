package model

import "time"

type User struct {
	UserID           uint      `gorm:"primary_key" json:"user_id"`
	UserName         string    `gorm:"size:255" json:"user_name"`
	Email            string    `gorm:"size:255" json:"email"`
	Password         string    `gorm:"size:255" json:"password"`
	InterestedTopics string    `gorm:"size:255" json:"interested_topics"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedDate      time.Time `json:"updated_date"`
	DeletedDate      time.Time `json:"deleted_date"`
}
