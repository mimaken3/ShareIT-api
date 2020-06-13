package model

import "time"

type ResultNotification struct {
	NotificationID        uint      `gorm:"primary_key" json:"notification_id"`
	UserID                uint      `json:"user_id"`
	SourceUserName        string    `json:"source_user_name"`
	SourceUserID          uint      `json:"source_user_id"`
	SourceUserIconName    string    `gorm:"size:500" json:"source_user_icon_name"`
	IsRead                int8      `json:"is_read"`
	DestinationTypeID     uint      `json:"destination_type_id"`
	DestinationTypeNameID uint      `json:"destination_type_name_id"`
	BehaviorTypeID        uint      `json:"behavior_type_id"`
	BehaviorTypeNameID    uint      `json:"behavior_type_name_id"`
	CreatedDate           time.Time `json:"created_date"`
}
