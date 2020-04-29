package model

type Like struct {
	LikeID    uint `gorm:"primary_key" json:"like_id"`
	UserID    uint `json:"user_id"`
	ArticleID uint `json:"article_Id"`
}
