package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleServiceStruct struct {
	articleRepo repository.ArticleRepository
}

// Application層はこのInterfaceに依存
type ArticleServiceInterface interface {
	// 全記事を取得
	FindAllArticlesService() (articles []model.Article, err error)

	// 特定のトピックを含む記事を取得
	FindArticlesByTopicIdService(articleIds []model.ArticleTopic) (articles []model.Article, err error)

	// 指定したトピックを含む記事のIDを取得
	FindArticleIdsByTopicIdService(topicID uint) (articleIds []model.ArticleTopic, err error)
}

// DIのための関数
func NewArticleService(a repository.ArticleRepository) ArticleServiceInterface {
	return &articleServiceStruct{articleRepo: a}
}
