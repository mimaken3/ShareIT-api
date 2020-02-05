package model

type ArticleTopic struct {
	ArticleTopicID uint `gorm:"primary_key"`
	ArticleID      uint
	TopicID        uint
}
