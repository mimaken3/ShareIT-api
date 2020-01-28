package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewArticleDB(db *gorm.DB) repository.ArticleRepository {
	return &articleInfraStruct{db: db}
}

// 全記事を取得
func (articleRepo *articleInfraStruct) FindAllArticles() (articles []model.Article, err error) {
	articleRepo.db.Find(&articles)
	return
}

// 特定のトピックを含む記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByTopicId(articleIds []model.ArticleTopic) (articles []model.Article, err error) {
	for i := 0; i < len(articleIds); i++ {
		// TODO: 要修正 毎回articleを作ってる
		var article = model.Article{}
		articleRepo.db.Where("article_id = ?", articleIds[i].ArticleID).Find(&article)
		articles = append(articles, article)
	}
	return
}

// 指定したトピックを含む記事のIDを取得
func (articleRepo *articleInfraStruct) FindArticleIdsByTopicId(topicID uint) (articleIds []model.ArticleTopic, err error) {
	// SELECT * FROM article_topic WHERE topic_id = :topicID;
	articleRepo.db.Where("topic_id = ?", topicID).Find(&articleIds)
	return articleIds, err
}
