package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事のトピックが更新されているか確認
func (a *articleServiceStruct) CheckUpdateArticleTopic(willBeUpdatedArticle model.Article) bool {
	isUpdate := a.articleRepo.CheckUpdateArticleTopic(willBeUpdatedArticle)

	return isUpdate
}
