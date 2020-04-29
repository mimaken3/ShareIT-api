package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// UserRepository is interface for infrastructure
type ArticleRepository interface {
	// 記事を投稿
	CreateArticle(createArticle model.Article, lastArticleId uint) (createdArticle model.Article, err error)

	// 記事を取得
	FindArticleByArticleId(articleId uint) (article model.Article, err error)

	// 全記事を取得(ページング)
	FindAllArticles(refPg int) (articles []model.Article, allPagingNum int, err error)

	// 記事を更新
	UpdateArticleByArticleId(willBeUpdatedArticle model.Article) (updatedArticle model.Article, err error)

	// 特定のユーザの全記事を取得(ページング)
	FindArticlesByUserId(userID uint, refPg int) (articles []model.Article, allPagingNum int, err error)

	// 特定のトピックを含む記事を取得
	FindArticlesByTopicId(articleIds []model.ArticleTopic) (articles []model.Article, err error)

	// 指定したトピックを含む記事トピックを取得
	FindArticleIdsByTopicId(topicID uint) (articleIds []model.ArticleTopic, err error)

	// 最後の記事IDを取得
	FindLastArticleId() (lastArticleId uint, err error)

	// 記事のトピックが更新されているか確認
	CheckUpdateArticleTopic(willBeUpdatedArticle model.Article) bool

	// 記事を削除
	DeleteArticleByArticleId(articleId uint) (err error)
}
