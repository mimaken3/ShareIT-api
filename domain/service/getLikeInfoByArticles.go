package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 各記事のいいね情報を取得
func (l *likeServiceStruct) GetLikeInfoByArtiles(userID uint, articles []model.Article) (updatedArticles []model.Article, err error) {
	isLiked, likeNum, err := l.likeRepo.GetLikeInfoByArtiles(userID, articles)

	for i := 0; i < len(articles); i++ {
		articles[i].IsLiked = isLiked[i]
		articles[i].LikeNum = likeNum[i]
	}
	updatedArticles = articles

	return
}
