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
		likeRepo.db.Table("likes").Where("article_id = ?", articleID).Count(&count)

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

// 最後のいいねIDを取得
func (likeRepo *likeInfraStruct) GetLastLikeID() (lastLikeID uint, err error) {
	like := model.Like{}
	// SELECT like_id FROM likes ORDER BY like_id DESC LIMIT 1;
	if result := likeRepo.db.Select("like_id").Last(&like); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}

	lastLikeID = like.LikeID
	return
}

// いいねを追加
func (likeRepo *likeInfraStruct) AddLike(userID uint, articleID uint, lastLikeID uint) (err error) {
	var like model.Like
	like.LikeID = lastLikeID + 1
	like.UserID = userID
	like.ArticleID = articleID

	likeRepo.db.Create(&like)

	return
}

// いいねを外す
func (likeRepo *likeInfraStruct) DeleteLike(userID uint, articleID uint) (err error) {
	like := model.Like{}
	likeRepo.db.Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&like)
	return
}
