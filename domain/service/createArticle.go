package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 記事を投稿
func (a *articleServiceStruct) CreateArticle(createArticle model.Article) (createdArticle model.Article, err error) {
	// 最後の記事IDを取得
	lastArticleId, err := a.articleRepo.FindLastArticleId()

	createdArticle, err = a.articleRepo.CreateArticle(createArticle, lastArticleId)
	if err != nil {
		log.Println(err)
	}
	return
}
