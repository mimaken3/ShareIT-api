package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// ArticleTopicRepository is interface for infrastructure
type ArticleTopicRepository interface {
	// 記事に紐づくトピックを追加
	CreateArticleTopic(article model.Article, lastArticleTopicId uint)

	// 最後の記事トピックIDを取得
	FindLastArticleTopicId() (lastArticleTopicId uint, err error)

	// 記事に紐づくトピックを削除
	DeleteArticleTopic(willBeDeletedArticle model.Article)
}
