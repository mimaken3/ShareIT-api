package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 特定のトピックを含む記事を取得
func (a *articleServiceStruct) FindArticlesByTopicIdService(articleIds []model.ArticleTopic, loginUserID uint, refPg int) (articles []model.Article, allPagingNum int, err error) {
	articles, allPagingNum, err = a.articleRepo.FindArticlesByTopicId(articleIds, loginUserID, refPg)
	if err != nil {
		log.Println(err)
	}
	return
}
