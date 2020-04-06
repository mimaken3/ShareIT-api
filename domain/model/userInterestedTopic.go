package model

type UserInterestedTopic struct {
	UserInterestedTopicID uint `gorm:"primary_key" json:"user_interested_topic_id"`
	ArticleID             uint `json:"article_id"`
	TopicID               uint `json:"topic_id"`
}
