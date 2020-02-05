package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 特定のユーザの全記事を取得
func (a *articleServiceStruct) FindArticlesByUserIdService(userID uint) (articles []model.Article, err error) {
	articles, err = a.articleRepo.FindArticlesByUserId(userID)
	if err != nil {
		log.Println(err)
	}
	return
}
