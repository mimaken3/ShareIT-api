package infrastructure

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleTopicInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewArticleTopicDB(db *gorm.DB) repository.ArticleTopicRepository {
	return &articleTopicInfraStruct{db: db}
}

// 記事に紐づく記事トピックを追加
func (articleTopicRepo *articleTopicInfraStruct) CreateArticleTopic(article model.Article, lastArticleTopicId uint) {
	articleTopic := model.ArticleTopic{}
	articleID := article.ArticleID
	topicsStr := article.ArticleTopics
	topicsArr := strings.Split(topicsStr, ",")

	// 記事トピックID
	insertArticleTopicId := lastArticleTopicId

	for _, topicStr := range topicsArr {
		insertArticleTopicId = insertArticleTopicId + 1
		if topicStr != "" {
			topicID, _ := strconv.Atoi(topicStr)
			// INSERT INTO article_topics VALUES(:lastArticleTopicId + 1, :articleID, :topicID);
			articleTopic.ArticleTopicID = insertArticleTopicId
			articleTopic.ArticleID = articleID
			articleTopic.TopicID = uint(topicID)
			articleTopicRepo.db.Create(&articleTopic)
		}
	}
}

// 最後の記事トピックIDを取得
func (articleTopicRepo *articleTopicInfraStruct) FindLastArticleTopicId() (lastArticleTopicId uint, err error) {
	articleTopic := model.ArticleTopic{}
	// SELECT article_topic_id FROM article_topics ORDER BY article_topic_id DESC LIMIT 1;
	articleTopicRepo.db.Select("article_topic_id").Last(&articleTopic)
	lastArticleTopicId = articleTopic.ArticleTopicID
	return
}

// 記事に紐づく記事トピックを削除
func (articleTopicRepo *articleTopicInfraStruct) DeleteArticleTopic(willBeDeletedArticle model.Article) {
	uintArticleID := willBeDeletedArticle.ArticleID

	// 物理削除
	articleTopicRepo.db.Where("article_id = ?", uintArticleID).Delete(&model.ArticleTopic{})
}

// トピックに紐づく記事トピックを削除
func (articleTopicRepo *articleTopicInfraStruct) DeleteArticleTopicByTopicID(topicID uint) (err error) {
	// 物理削除
	articleTopicRepo.db.Where("topic_id = ?", topicID).Delete(&model.ArticleTopic{})

	return
}
