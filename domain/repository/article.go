package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// UserRepository is interface for infrastructure
type ArticleRepository interface {
	// 記事を全取得
	FindAllArticles() (articles []model.Article, err error)

	// 特定のユーザの全記事を取得
	FindArticlesByUserId(userID uint) (articles []model.Article, err error)

	// 特定のトピックを含む記事を取得
	FindArticlesByTopicId(articleIds []model.ArticleTopic) (articles []model.Article, err error)

	// 指定したトピックを含む記事のIDを取得
	FindArticleIdsByTopicId(topicID uint) (articleIds []model.ArticleTopic, err error)
}
