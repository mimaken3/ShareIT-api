package model

import "time"

type Topic struct {
	TopicID        uint      `gorm:"primary_key" json:"user_id"`
	TopicName      string    `gorm:"size:255" json:"topic_name"`
	ProposedUserID uint      `json:"proposed_user_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
	DeletedDate    time.Time `json:"deleted_date"`
	IsDeleted      int8      `json:"-"`
}
