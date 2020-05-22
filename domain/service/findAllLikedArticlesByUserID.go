package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 特定のユーザのいいねした記事を取得(ページング)
func (a *articleServiceStruct) FindAllLikedArticlesByUserID(userID uint, loginUserID uint, refPg int) (articles []model.Article, allPagingNum int, err error) {
	articles, allPagingNum, err = a.articleRepo.FindAllLikedArticlesByUserID(userID, loginUserID, refPg)
	if err != nil {
		log.Println(err)
	}
	return
}
