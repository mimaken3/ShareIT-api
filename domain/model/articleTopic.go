package model

type ArticleTopic struct {
	ArticleTopicID uint `gorm:"primary_key" json:"article_topic_id"`
	ArticleID      uint `json:"article_id"`
	TopicID        uint `json:"topic_id"`
}
