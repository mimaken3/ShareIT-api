package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// LikeRepository is interface for infrastructure
type LikeRepository interface {
	// 最後のいいねIDを取得
	// GetLastLikeID() (likeID uint, err error)

	// 各記事のいいね情報を取得
	GetLikeInfoByArtiles(userID uint, articles []model.Article) (isLiked []bool, likeNum []int, err error)

	// いいねON/OFF
	// ToggleLikeByArticle(userID uint, isLiked bool, article model.Article) (isLike bool, likeNum int, err error)
}
