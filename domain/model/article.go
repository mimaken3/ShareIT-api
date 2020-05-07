package model

import "time"

type Article struct {
	ArticleID      uint      `gorm:"primary_key" json:"article_id"`
	ArticleTitle   string    `gorm:"size:255" json:"article_title"`
	ArticleContent string    `gorm:"size:1000" json:"article_content"`
	ArticleTopics  string    `gorm:"size:255" json:"article_topics"`
	IsLiked        bool      `json:"is_liked"`
	LikeNum        int       `json:"like_num"`
	CreatedUserID  uint      `json:"created_user_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
	DeletedDate    time.Time `json:"deleted_date"`
	IsPrivate      int8      `json:"is_private"`
	IsDeleted      int8      `json:"-"`
}
