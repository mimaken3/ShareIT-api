package model

import "time"

type Article struct {
	ArticleID      uint   `gorm:"primary_key"`
	ArticleTitle   string `gorm:"size:255"`
	CreatedUserID  uint
	ArticleContent string `gorm:"size:1000"`
	ArticleTopics  string `gorm:"size:255"`
	CreatedDate    time.Time
	UpdatedDate    time.Time
	DeletedDate    time.Time
}
