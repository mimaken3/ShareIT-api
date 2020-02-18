package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事に紐づくトピックを更新
func (a *articleTopicServiceStruct) UpdateArticleTopic(willBeUpdatedArticle model.Article) {
	// 記事に紐づくトピックを全削除
	a.articleTopicRepo.DeleteArticleTopic(willBeUpdatedArticle)

	// 最後の記事トピックIDを取得
	lastArticleTopicId, _ := a.articleTopicRepo.FindLastArticleTopicId()

	// 記事に紐づくトピックを追加
	a.articleTopicRepo.CreateArticleTopic(willBeUpdatedArticle, lastArticleTopicId)
}
