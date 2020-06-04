package model

import "time"

type Notification struct {
	NotificationID     uint      `gorm:"primary_key" json:"notification_id"`
	UserID             uint      `json:"user_id"`
	SourceUserID       uint      `json:"source_user_id"`
	SourceUserIconName string    `gorm:"size:500" json:"source_user_icon_name"`
	IsRead             int8      `json:"is_read"`
	DestinationID      uint      `json:"destination_id"`
	CreatedDate        time.Time `json:"created_date"`
}
