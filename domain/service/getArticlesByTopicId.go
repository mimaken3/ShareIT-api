package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 特定のトピックを含む記事を取得
func (a *articleServiceStruct) FindArticlesByTopicIdService(articleIds []model.ArticleTopic) (articles []model.Article, err error) {
	articles, err = a.articleRepo.FindArticlesByTopicId(articleIds)
	if err != nil {
		log.Println(err)
	}
	return
}
