package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 記事に紐づくトピックを追加
func (a *articleTopicServiceStruct) CreateArticleTopic(article model.Article) {
	var err error
	// 最後の記事トピックIDを取得
	lastArticleTopicId, err := a.articleTopicRepo.FindLastArticleTopicId()

	if err != nil {
		return
	}

	a.articleTopicRepo.CreateArticleTopic(article, lastArticleTopicId)
}
