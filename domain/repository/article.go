package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// UserRepository is interface for infrastructure
type ArticleRepository interface {
	// 記事を全取得
	FindAllArticles() (articles []model.Article, err error)
}
