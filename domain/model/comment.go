package model

import "time"

type Comment struct {
	CommentID   uint      `gorm:"primary_key" json:"comment_id"`
	ArticleID   uint      `json:"article_Id"`
	UserID      uint      `json:"user_id"`
	Content     string    `gorm:"size:1000" json:"content"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   int8      `json:"-"`
}
