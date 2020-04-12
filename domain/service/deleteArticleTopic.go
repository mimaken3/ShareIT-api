package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事に紐づくトピックを削除
func (a *articleTopicServiceStruct) DeleteArticleTopic(willBeDeletedArticle model.Article) {
	a.articleTopicRepo.DeleteArticleTopic(willBeDeletedArticle)
}
