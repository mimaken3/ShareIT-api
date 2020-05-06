package service

import (
	"strconv"
	"strings"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 記事を検索(ページング)
func (a *articleServiceStruct) SearchAllArticles(refPg int, userID uint, topicIDStr string) (searchedArticles []model.Article, allPagingNum int, err error) {
	topicIDStrArr := strings.Split(topicIDStr, " ")

	var topicIDs []uint
	for _, value := range topicIDStrArr {
		intV, _ := strconv.Atoi(value)
		uintV := uint(intV)
		topicIDs = append(topicIDs, uintV)
	}

	if userID == 0 && topicIDs[0] == 0 {
		//「全ユーザ」かつ「トピックの選択なし」の場合
		searchedArticles, allPagingNum, err = a.articleRepo.FindAllArticles(refPg)
		return
	} else if userID != 0 && topicIDs[0] == 0 {
		// 「特定のユーザ」かつ「全トピック」の場合
		searchedArticles, allPagingNum, err = a.articleRepo.FindArticlesByUserId(userID, refPg)
		return
	}
	searchedArticles, allPagingNum, err = a.articleRepo.SearchAllArticles(refPg, userID, topicIDs)
	return
}
