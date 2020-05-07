package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 全記事を取得(ページング)
func (a *articleServiceStruct) FindAllArticlesService(refPg int, userID uint) (articles []model.Article, allPagingNum int, err error) {
	articles, allPagingNum, err = a.articleRepo.FindAllArticles(refPg, userID)
	if err != nil {
		log.Println(err)
	}
	return articles, allPagingNum, err
}
