package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type likeInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewLikeDB(db *gorm.DB) repository.LikeRepository {
	return &likeInfraStruct{db: db}
}

// 各記事のいいね情報を取得
func (likeRepo *likeInfraStruct) GetLikeInfoByArtiles(userID uint, articles []model.Article) (isLiked []bool, likeNum []int, err error) {

	for _, article := range articles {
		articleID := article.ArticleID
		var count int
		likeRepo.db.Table("likes").Where("user_id = ? AND article_id = ?", userID, articleID).Count(&count)

		// いいね数を格納
		likeNum = append(likeNum, count)

		// 取得するユーザがいいねしてるかどうか
		var like model.Like
		if result := likeRepo.db.Where("user_id = ? AND article_id = ?", userID, articleID).Find(&like); result.Error != nil {
			isLiked = append(isLiked, false)
		} else {
			isLiked = append(isLiked, true)
		}
	}

	return
}
