package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 記事を取得
func (a *articleServiceStruct) FindArticleByArticleId(articleId uint) (article model.Article, err error) {
	article, err = a.articleRepo.FindArticleByArticleId(articleId)
	if err != nil {
		log.Println(err)
	}

	return
}
