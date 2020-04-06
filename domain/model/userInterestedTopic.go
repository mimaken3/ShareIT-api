package model

type UserInterestedTopic struct {
	UserInterestedTopicsID uint `gorm:"primary_key" json:"user_interested_topic_id"`
	UserID                 uint `json:"article_id"`
	TopicID                uint `json:"topic_id"`
}
