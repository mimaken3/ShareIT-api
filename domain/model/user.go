package model

import "time"

type User struct {
	UserID           uint   `gorm:"primary_key"`
	UserName         string `gorm:"size:255"`
	Email            string `gorm:"size:255"`
	Password         string `gorm:"size:255"`
	InterestedTopics string `gorm:"size:255"`
	CreatedDate      time.Time
	UpdatedDate      time.Time
	DeletedDate      time.Time
}
