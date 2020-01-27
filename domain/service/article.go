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
}

// DIのための関数
func NewArticleService(a repository.ArticleRepository) ArticleServiceInterface {
	return &articleServiceStruct{articleRepo: a}
}
