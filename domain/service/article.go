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
	// 記事を投稿
	CreateArticle(createArticle model.Article) (createdArticle model.Article, err error)

	// 記事を取得
	FindArticleByArticleId(articleId uint) (article model.Article, err error)

	// 全記事を取得
	FindAllArticlesService() (articles []model.Article, err error)

	// 記事を更新
	UpdateArticleByArticleId(willBeUpdatedArticle model.Article) (updatedArticle model.Article, err error)

	// 特定のユーザの全記事を取得
	FindArticlesByUserIdService(userID uint) (articles []model.Article, err error)

	// 特定のトピックを含む記事を取得
	FindArticlesByTopicIdService(articleIds []model.ArticleTopic) (articles []model.Article, err error)

	// 指定したトピックを含む記事トピックを取得
	FindArticleIdsByTopicIdService(topicID uint) (articleIds []model.ArticleTopic, err error)

	// 最後の記事IDを取得
	FindLastArticleId() (lastArticleId uint, err error)

	// 記事のトピックが更新されているか確認
	CheckUpdateArticleTopic(willBeUpdatedArticle model.Article) bool

	// 記事を削除
	DeleteArticleByArticleId(articleId uint) (err error)
}

// DIのための関数
func NewArticleService(a repository.ArticleRepository) ArticleServiceInterface {
	return &articleServiceStruct{articleRepo: a}
}
