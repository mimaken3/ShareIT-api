package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 記事を更新
func (a *articleServiceStruct) UpdateArticleByArticleId(willBeUpdatedArticle model.Article) (updatedArticle model.Article, err error) {
	updatedArticle, err = a.articleRepo.UpdateArticleByArticleId(willBeUpdatedArticle)

	if err != nil {
		log.Println(err)
	}

	return
}
