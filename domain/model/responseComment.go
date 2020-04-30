package model

import "time"

type ResponseComment struct {
	CommentID   uint      `gorm:"primary_key" json:"comment_id"`
	ArticleID   uint      `json:"article_Id"`
	UserID      uint      `json:"user_id"`
	UserName    string    `gorm:"size:255" json:"user_name"`
	IconName    string    `gorm:"size:255" json:"icon_name"`
	Content     string    `gorm:"size:1000" json:"content"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   int8      `json:"-"`
}
