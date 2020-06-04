package model

type Destination struct {
	DestinationID         uint `gorm:"primary_key" json:"destination_id"`
	DestinationTypeID     uint `json:"destination_type_id"`
	DestinationTypeNameID uint `json:"destination_type_name_id"`
	BehaviorTypeID        uint `json:"behavior_type_id"`
	BehaviorTypeNameID    uint `json:"behavior_type_name_id"`
}
