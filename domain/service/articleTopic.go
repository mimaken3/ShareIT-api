package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleTopicServiceStruct struct {
	articleTopicRepo repository.ArticleTopicRepository
}

// Application層はこのInterfaceに依存
type ArticleTopicServiceInterface interface {
	// 記事に紐づくトピックを追加
	CreateArticleTopic(article model.Article)

	// 記事に紐づくトピックを更新
	UpdateArticleTopic(willBeUpdatedArticle model.Article)

	// 記事に紐づくトピックを削除
	DeleteArticleTopic(willBeDeletedArticle model.Article)
}

// DIのための関数
func NewArticleTopicService(a repository.ArticleTopicRepository) ArticleTopicServiceInterface {
	return &articleTopicServiceStruct{articleTopicRepo: a}
}
