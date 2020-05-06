package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// LikeRepository is interface for infrastructure
type LikeRepository interface {
	// 最後のいいねIDを取得
	GetLastLikeID() (lastLikeID uint, err error)

	// いいねを追加
	AddLike(userID uint, articleID uint, lastLikeID uint) (err error)

	// いいねを外す
	DeleteLike(userID uint, articleID uint) (err error)

	// 各記事のいいね情報を取得
	GetLikeInfoByArtiles(userID uint, articles []model.Article) (isLiked []bool, likeNum []int, err error)

	// 記事のいいねを削除
	DeleteLikeByArticleID(articleID uint) (err error)

	// ユーザが付けたいいねを全削除
	DeleteLikeByUserID(userID uint) (err error)
}
