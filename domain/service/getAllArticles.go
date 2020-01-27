package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 全記事を取得
func (a *articleServiceStruct) FindAllArticlesService() (articles []model.Article, err error) {
	articles, err = a.articleRepo.FindAllArticles()
	if err != nil {
		log.Println(err)
	}
	return articles, err
}
